package download

import (
	"context"

	"tools/internal/logger"
	"tools/ipfs-cluster/sdk"
	myshell "tools/ipfs-cluster/thridparty/shell"
)

// OneFile 下载文件到指定路径
// @param cid 文件cid
// @param path 文件保存路径
// @return error 错误信息
// @example DownloadFile("QmXY2gC5d86QK129MQKF9R9Kprrkv9zm7xJ6N63bEK9uXG", "/path/to/file") => nil
func OneFile(cid string, path string) error {
	logger.GetIpfsLogger().Debugf(">>> 开始从IPFS下载CID %v 对应文件 ", cid)
	ctx := context.Background()
	ipfs := sdk.GetSDK().IPFS(ctx)
	// err := ipfs.Get(cid, path)
	err := myshell.Get(ipfs, cid, path)
	if err != nil {
		logger.GetIpfsLogger().Errorf(">>> 下载CID %v 对应文件失败: %v", cid, err)
		return err
	}
	return nil
}
