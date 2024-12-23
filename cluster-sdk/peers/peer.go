package peers

import (
	"context"

	"github.com/ipfs-cluster/ipfs-cluster/api"
	"ipfs-cluster/internal/logger"

	"ipfs-cluster/sdk"
)

// Peers 获取集群中节点
// @return *[]api.ID 节点列表
func Peers() *[]api.ID {
	logger.GetIPFSLogger().Debugf("starting get ipfs cluster peers")
	idSlice := make(chan api.ID)
	ids := make([]api.ID, 0)

	ctx := context.Background()
	go func(ctx context.Context) {
		if err := sdk.GetSDK().Peers(ctx, idSlice); err != nil {
			logger.GetIPFSLogger().Errorf("getting ipfs cluster peers error %v", err)
		}
	}(ctx)

	for id := range idSlice {
		ids = append(ids, id)
	}
	return &ids
}
