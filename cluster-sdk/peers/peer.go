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
	logger.GetIPFSLogger().Debugf(">>> 获取集群中节点")
	idSlice := make(chan api.ID)
	ids := make([]api.ID, 0)

	ctx := context.Background()
	go func(ctx context.Context) {
		if err := sdk.GetSDK().Peers(ctx, idSlice); err != nil {
			logger.GetIPFSLogger().Errorf(">>> 获取集群中节点失败 %v", err)
		}
	}(ctx)

	for id := range idSlice {
		ids = append(ids, id)
	}
	return &ids
}
