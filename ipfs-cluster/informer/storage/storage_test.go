package storage

import (
	"context"
	"errors"
	"fmt"
	"syscall"
	"testing"

	"github.com/ipfs-cluster/ipfs-cluster/api"
	"github.com/ipfs-cluster/ipfs-cluster/test"

	rpc "github.com/libp2p/go-libp2p-gorpc"
	"github.com/shirou/gopsutil/v4/disk"
)

type badRPCService struct {
}

func badRPCClient(t *testing.T) *rpc.Client {
	s := rpc.NewServer(nil, "mock")
	c := rpc.NewClientWithServer(nil, "mock", s)
	err := s.RegisterName("IPFSConnector", &badRPCService{})
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func (mock *badRPCService) RepoStat(ctx context.Context, in struct{}, out *api.IPFSRepoStat) error {
	return errors.New("fake error")
}

// Returns the first metric
func getMetrics(t *testing.T, inf *Informer) api.Metric {
	t.Helper()
	metrics := inf.GetMetrics(context.Background())
	if len(metrics) != 1 {
		t.Fatal("expected 1 metric")
	}
	return metrics[0]
}

func Test(t *testing.T) {
	ctx := context.Background()
	cfg := &Config{}
	cfg.Default()
	inf, err := NewInformer(cfg)
	if err != nil {
		t.Fatal(err)
	}
	defer inf.Shutdown(ctx)
	m := getMetrics(t, inf)
	if m.Valid {
		t.Error("metric should be invalid")
	}
	inf.SetClient(test.NewMockRPCClient(t))
	m = getMetrics(t, inf)
	if !m.Valid {
		t.Error("metric should be valid")
	}
}

func TestFreeSpace(t *testing.T) {
	ctx := context.Background()
	cfg := &Config{}
	cfg.Default()
	// cfg.MetricType = MetricFreeSpace

	inf, err := NewInformer(cfg)
	if err != nil {
		t.Fatal(err)
	}
	defer inf.Shutdown(ctx)
	m := getMetrics(t, inf)
	if m.Valid {
		t.Error("metric should be invalid")
	}
	inf.SetClient(test.NewMockRPCClient(t))
	m = getMetrics(t, inf)
	if !m.Valid {
		t.Error("metric should be valid")
	}
	// The mock client reports 100KB and 2 pins of 1 KB
	if m.Value != "98000" {
		t.Error("bad metric value")
	}
}

func TestRepoSize(t *testing.T) {
	ctx := context.Background()
	cfg := &Config{}
	cfg.Default()
	// cfg.MetricType = MetricRepoSize

	inf, err := NewInformer(cfg)
	if err != nil {
		t.Fatal(err)
	}
	defer inf.Shutdown(ctx)
	m := getMetrics(t, inf)
	if m.Valid {
		t.Error("metric should be invalid")
	}
	inf.SetClient(test.NewMockRPCClient(t))
	m = getMetrics(t, inf)
	if !m.Valid {
		t.Error("metric should be valid")
	}
	// The mock client reports 100KB and 2 pins of 1 KB
	if m.Value != "2000" {
		t.Error("bad metric value")
	}
}

func TestWithErrors(t *testing.T) {
	ctx := context.Background()
	cfg := &Config{}
	cfg.Default()
	inf, err := NewInformer(cfg)
	if err != nil {
		t.Fatal(err)
	}
	defer inf.Shutdown(ctx)
	inf.SetClient(badRPCClient(t))
	m := getMetrics(t, inf)
	if m.Valid {
		t.Errorf("metric should be invalid")
	}
}

func TestReadStorage(t *testing.T) {
	path := "/data"

	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		t.Fatal(err)
	}
	all := int64(fs.Blocks) * fs.Bsize
	free := int64(fs.Bfree) * fs.Bsize
	used := all - free
	d := diskStatus{All: all, Used: used, Free: free}
	fmt.Println(d)
}

func TestReadStorageWithLib(t *testing.T) {
	partitions, err := disk.Partitions(true)
	if err != nil {
		fmt.Println("Error getting partitions:", err)
		return
	}

	for _, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			fmt.Println("Error getting disk usage for", partition.Mountpoint, ":", err)
			continue
		}
		fmt.Printf("Partition: %s\n", partition.Mountpoint)
		fmt.Printf("Total Size: %.2f GB\n", float64(usage.Total)/(1024*1024*1024))
		fmt.Printf("Used Size: %.2f GB\n", float64(usage.Used)/(1024*1024*1024))
		fmt.Printf("Free Size: %.2f GB\n", float64(usage.Free)/(1024*1024*1024))
		fmt.Println()
	}
}
