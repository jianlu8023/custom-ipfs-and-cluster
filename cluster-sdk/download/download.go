package download

import (
	"context"

	"ipfs-cluster/internal/logger"
	"ipfs-cluster/sdk"
	myshell "ipfs-cluster/thridparty/shell"
)

// OneFile 下载文件到指定路径
// @param cid 文件cid
// @param path 文件保存路径
// @return error 错误信息
// @example DownloadFile("QmXY2gC5d86QK129MQKF9R9Kprrkv9zm7xJ6N63bEK9uXG", "/path/to/file") => nil
func OneFile(cid string, path string) error {
	logger.GetIPFSLogger().Debugf("starting download file %v from ipfs cluster", cid)
	ctx := context.Background()
	ipfs := sdk.GetSDK().IPFS(ctx)

	err := myshell.Get(ipfs, cid, path)
	if err != nil {
		logger.GetIPFSLogger().Errorf("download file %v error %v", cid, err)
		return err
	}
	return nil
}
