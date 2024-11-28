// Package storage implements an ipfs-cluster informer which can provide different
// disk-related metrics from the IPFS daemon as an api.Metric.
package storage

import (
	"context"
	"fmt"
	"strings"
	"sync"

	logging "github.com/ipfs/go-log/v2"
	rpc "github.com/libp2p/go-libp2p-gorpc"

	disklib "github.com/shirou/gopsutil/v4/disk"
	"go.opencensus.io/trace"

	"github.com/ipfs-cluster/ipfs-cluster/api"
)

// MetricMount identifies the mount of metric to fetch from the IPFS daemon.
// type MetricMount string

var MetricName = "storage"

const (
	// MetricFreeSpace provides the available space reported by IPFS
	// MetricFreeSpace MetricType = iota
	// MetricRepoSize provides the used space reported by IPFS
	// MetricRepoSize

	defaultMount = "/"
)

// String returns a string representation for MetricType.
// func (t MetricType) String() string {
// 	switch t {
// 	case MetricFreeSpace:
// 		return "freespace"
// 	case MetricRepoSize:
// 		return "reposize"
// 	}
// 	return ""
// }

var logger = logging.Logger("storage")

// Informer is a simple object to implement the ipfscluster.Informer
// and Component interfaces.
type Informer struct {
	config *Config // set when created, readonly

	mu        sync.Mutex // guards access to following fields
	rpcClient *rpc.Client
}

// NewInformer returns an initialized informer using the given InformerConfig.
func NewInformer(cfg *Config) (*Informer, error) {
	err := cfg.Validate()
	if err != nil {
		return nil, err
	}
	logger.Debugf("setting storage NewInformer")
	return &Informer{
		config: cfg,
	}, nil
}

// Name returns the name of the metric issued by this informer.
func (storage *Informer) Name() string {
	// return storage.config.MetricType.String()
	return MetricName
}

// SetClient provides us with an rpc.Client which allows
// contacting other components in the cluster.
func (storage *Informer) SetClient(c *rpc.Client) {
	storage.mu.Lock()
	defer storage.mu.Unlock()
	storage.rpcClient = c
}

// Shutdown is called on cluster shutdown. We just invalidate
// any metrics from this point.
func (storage *Informer) Shutdown(ctx context.Context) error {
	_, span := trace.StartSpan(ctx, "informer/storage/Shutdown")
	defer span.End()

	storage.mu.Lock()
	defer storage.mu.Unlock()

	storage.rpcClient = nil
	return nil
}

// GetMetrics returns the metric obtained by this Informer. It must always
// return at least one metric.
func (storage *Informer) GetMetrics(ctx context.Context) []api.Metric {
	ctx, span := trace.StartSpan(ctx, "informer/storage/GetMetric")
	defer span.End()

	var err error

	storage.mu.Lock()
	rpcClient := storage.rpcClient
	storage.mu.Unlock()

	if rpcClient == nil {
		return []api.Metric{
			{
				Name:  storage.Name(),
				Valid: false,
			},
		}
	}

	partitions, err := disklib.Partitions(true)
	if err != nil {
		logger.Errorf("getting storage info failed %s", err)
	}
	d := diskStatus{}

	partition := getTargetPartition(partitions, strings.TrimSpace(storage.config.MetricMount))
	if partition == nil {
		logger.Errorf("getting storage partition %v info failed", storage.config.MetricMount)
	} else {
		usage, err := disklib.Usage(partition.Mountpoint)
		if err != nil {
			logger.Errorf("getting storage partition %v info failed %v", partition.Mountpoint, err)
		}
		d.Free = float64(usage.Free) / (1024 * 1024 * 1024)
		d.Used = float64(usage.Used) / (1024 * 1024 * 1024)
		d.All = float64(usage.Total) / (1024 * 1024 * 1024)
		d.Weight = int64(d.Free)
	}

	valid := err == nil
	logger.Debugf("getting storage partition %v info %v valid %v", storage.config.MetricMount, d.String(), valid)

	m := api.Metric{
		Name:          storage.Name(),
		Value:         fmt.Sprintf("%.2f", d.Free),
		Valid:         valid,
		Weight:        d.Weight,
		Partitionable: false,
	}

	m.SetTTL(storage.config.MetricTTL)

	return []api.Metric{m}
}

func getTargetPartition(partitions []disklib.PartitionStat, mount string) *disklib.PartitionStat {
	targetMountLike := make([]string, 0)
	for _, partition := range partitions {
		logger.Debugf("current mount %s %v", partition.Mountpoint, mount == partition.Mountpoint)
		if mount == strings.TrimSpace(partition.Mountpoint) {
			return &partition
		}
		if strings.Contains(partition.Mountpoint, mount) {
			targetMountLike = append(targetMountLike, partition.Mountpoint)
		}

	}
	if len(targetMountLike) != 0 {
		logger.Errorf("config mount %s have not found but found like mounts %v", mount, targetMountLike)
	}
	return nil
}
