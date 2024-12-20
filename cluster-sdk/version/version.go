package version

import (
	"context"

	"ipfs-cluster/internal/logger"
	"ipfs-cluster/sdk"
)

// Version 测试version方法
// @return error 可能错误
func Version() error {
	logger.GetIPFSLogger().Debugf(">>> 获取集群版本信息")
	ctx := context.Background()
	_, err := sdk.GetSDK().Version(ctx)
	if err != nil {
		logger.GetIPFSLogger().Errorf(">>> 获取集群版本信息失败: %v", err)
		return err
	}
	return nil
}
