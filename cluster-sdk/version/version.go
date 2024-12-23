package version

import (
	"context"

	"ipfs-cluster/internal/logger"
	"ipfs-cluster/sdk"
)

// Version 测试version方法
// @return error 可能错误
func Version() error {
	logger.GetIPFSLogger().Debugf("starting get ipfs cluster version")
	ctx := context.Background()
	_, err := sdk.GetSDK().Version(ctx)
	if err != nil {
		logger.GetIPFSLogger().Errorf("getting ipfs cluster version error %v", err)
		return err
	}
	return nil
}
