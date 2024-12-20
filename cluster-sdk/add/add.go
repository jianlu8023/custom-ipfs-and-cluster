package add

import (
	"context"
	"time"

	"github.com/ipfs-cluster/ipfs-cluster/api"
	"ipfs-cluster/internal/logger"

	"ipfs-cluster/sdk"
)

// OneFile 添加单个文件
// @param conf *Info 上传文件新息
// @return *api.AddedOutput 返回上传文件新息
// @return error 错误信息
func OneFile(conf *Info) (*api.AddedOutput, error) {
	logger.GetIPFSLogger().Debugf(">>> 开始上传 %v 文件到IPFS", conf.FileName)

	params := api.DefaultAddParams()
	// 循环添加文件
	paths := make([]string, 0)
	paths = append(paths, conf.FilePath)

	// 设置集群中副本情况，最小和最大
	if conf.ReplicationFactorMin != 0 {
		params.ReplicationFactorMin = conf.ReplicationFactorMin
	}
	if conf.ReplicationFactorMax != 0 {
		params.ReplicationFactorMax = conf.ReplicationFactorMax
	}

	// 设置文件过期时间
	if len(conf.ExpireAt) > 0 {
		duration, err := time.ParseDuration(conf.ExpireAt)
		if err != nil {
			logger.GetIPFSLogger().Errorf(">>> 解析过期时间失败: %v", err)
		} else {
			params.ExpireAt = time.Now().Add(duration)
		}
	}

	params.Recursive = false

	params.Name = conf.FileName

	// p.Metadata = parseMetadata(c.StringSlice("metadata"))
	// p.Name = name
	// if c.String("allocations") != "" {
	// 	p.UserAllocations = api.StringsToPeers(strings.Split(c.String("allocations"), ","))
	// }
	// p.NoPin = c.Bool("no-pin")
	// p.Format = c.String("format")
	// p.Shard = shard
	// p.ShardSize = c.Uint64("shard-size")
	// p.Shard = false
	// p.Recursive = c.Bool("recursive")
	// p.Local = c.Bool("local")
	// p.Layout = c.String("layout")
	// p.Chunker = c.String("chunker")
	// p.RawLeaves = c.Bool("raw-leaves")
	// p.Hidden = c.Bool("hidden")
	// p.Wrap = c.Bool("wrap-with-directory") || len(paths) > 1
	// p.CidVersion = c.Int("cid-version")
	// p.HashFun = c.String("hash")

	params.CidVersion = 1

	out := make(chan api.AddedOutput)
	// ctx, cancelFunc := context.WithTimeout(context.Background(), config.Timeout10)
	// ctx, cancelFunc := context.WithCancel(context.Background())
	// defer cancelFunc()
	ctx := context.Background()
	go func(c context.Context) {
		if err := sdk.GetSDK().Add(c, paths, params, out); err != nil {
			logger.GetIPFSLogger().Errorf(">>> 添加文件 %v 失败 %v", conf.FileName, err)
		}
	}(ctx)

	select {
	case outed := <-out:
		logger.GetIPFSLogger().Debugf(">>> 添加文件 %v 成功", conf.FileName)
		return &outed, nil
	case <-ctx.Done():
		logger.GetIPFSLogger().Errorf(">>> 添加文件 %v 超时", conf.FileName)
		return nil, ctx.Err()
	}
}
