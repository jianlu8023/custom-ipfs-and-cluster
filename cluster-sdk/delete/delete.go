package delete

import (
	"context"

	"github.com/ipfs-cluster/ipfs-cluster/api"
	shell "github.com/ipfs/go-ipfs-api"
	"ipfs-cluster/internal/logger"

	"ipfs-cluster/sdk"
)

// Delete 删除文件
// @param cid 文件在ipfs中的hash
// @return bool 是否删除成功
// @return error 错误信息
func Delete(cid string) (bool, error) {

	decodeCid, err := api.DecodeCid(cid)
	if err != nil {
		logger.GetIPFSLogger().Errorf(">>> 解析CID失败 %v", err)
		return false, err
	}

	ctx := context.Background()
	_, err = sdk.GetSDK().Unpin(ctx, decodeCid)
	if err != nil {
		logger.GetIPFSLogger().Errorf(">>> 删除CID %v 失败 %v", cid, err)
		return false, err
	}
	return true, nil

}

// WithGarbageCollection 删除文件并执行垃圾回收
// @param cid 文件在ipfs中的hash
// @return bool 是否删除成功
// @return error 错误信息
// @example DeleteWithGarbageCollection("QmXZKWcqQkv876w9kq891J4vVyYHx16vZm76YVG91Q77") => true, nil
func WithGarbageCollection(cid string) (bool, error) {
	_, err := Delete(cid)
	if err != nil {
		logger.GetIPFSLogger().Errorf(">>> 删除CID %v 失败 %v", cid, err)
		return false, err
	}

	_, err = GarbageCollection()
	if err != nil {
		logger.GetIPFSLogger().Errorf(">>> 删除CID %v 后进行垃圾回收失败 %v", cid, err)
		return false, err
	}
	return true, nil
}

// GarbageCollection 垃圾回收
// @return bool 是否成功
// @return error 错误信息
func GarbageCollection() (bool, error) {

	ctx := context.Background()
	gc, err := sdk.GetSDK().RepoGC(ctx, false)
	if err != nil {
		logger.GetIPFSLogger().Errorf(">>> 集群垃圾回收失败 %v", err)
		return false, err
	}
	logger.GetIPFSLogger().Debugf(">>> 集群垃圾回收结果 %v", gc)
	return true, nil
}

// OnTargetNode 删除文件
// @param cid 文件在ipfs中的hash
// @param targetNode 目标节点 /ip4/127.0.0.1/tcp/5001
// @return bool 是否删除成功
// @return error 错误信息
// @example OnTargetNode("QmXZKWcqQkv876w9kq891J4vVyYHx16vZm76YVG91Q77", "/ip4/127.0.0.1/tcp/5001") => true, nil
func OnTargetNode(cid string, targetNode string) (bool, error) {
	sh := shell.NewShell(targetNode)
	err := sh.Unpin(cid)
	if err != nil {
		logger.GetIPFSLogger().Errorf(">>> 在目标节点 %v 删除CID %v 失败 %v", targetNode, cid, err)
		return false, err
	}
	return true, nil
}

// OnTargetNodeWithGarbageCollection 删除文件并执行垃圾回收
// @param cid 文件在ipfs中的hash
// @param targetNode 目标节点 /ip4/127.0.0.1/tcp/5001
// @return bool 是否删除成功
// @return error 错误信息
func OnTargetNodeWithGarbageCollection(cid string, targetNode string) (bool, error) {
	_, err := OnTargetNode(cid, targetNode)
	if err != nil {
		logger.GetIPFSLogger().Errorf(">>> 在目标节点 %v 删除CID %v 失败 %v", targetNode, cid, err)
		return false, err
	}

	ctx := context.Background()
	sh := shell.NewShell(targetNode)
	err = sh.Request("repo/gc", cid).
		Option("recursive", true).
		Exec(ctx, nil)

	if err != nil {
		logger.GetIPFSLogger().Errorf(">>> 在目标节点 %v 垃圾回收失败 %v", targetNode, err)
		return false, err
	}
	_, err = GarbageCollection()
	if err != nil {
		logger.GetIPFSLogger().Errorf(">>> 集群垃圾回收失败 %v", err)
		return false, err
	}
	return true, nil
}
