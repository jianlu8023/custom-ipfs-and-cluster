package delete

import (
	"context"

	"github.com/ipfs-cluster/ipfs-cluster/api"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/jianlu8023/go-tools/pkg/format/json"
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
		logger.GetIPFSLogger().Errorf("parse cid error %v", err)
		return false, err
	}

	ctx := context.Background()
	_, err = sdk.GetSDK().Unpin(ctx, decodeCid)
	if err != nil {
		logger.GetIPFSLogger().Errorf("delete cid %v error %v", cid, err)
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
		logger.GetIPFSLogger().Errorf("delete cid %v error %v", cid, err)
		return false, err
	}

	_, err = GarbageCollection()
	if err != nil {
		logger.GetIPFSLogger().Errorf("gc error %v", err)
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
		logger.GetIPFSLogger().Errorf("ipfs cluster gc error %v", err)
		return false, err
	}
	toJSON, _ := json.ToJSON(gc)
	logger.GetIPFSLogger().Debugf("ipfs cluster gc result %v", toJSON)
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
		logger.GetIPFSLogger().Errorf("delete cid %v on %v error %v", cid, targetNode, err)
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
		logger.GetIPFSLogger().Errorf("delete cid %v on %v error %v", cid, targetNode, err)
		return false, err
	}

	ctx := context.Background()
	sh := shell.NewShell(targetNode)

	if err = sh.Request("repo/gc", cid).
		Option("recursive", true).
		Exec(ctx, nil); err != nil {
		logger.GetIPFSLogger().Errorf("ipfs gc on %v error %v", targetNode, err)
		return false, err
	}
	_, err = GarbageCollection()
	if err != nil {
		logger.GetIPFSLogger().Errorf("ipfs cluster gc error %v", err)
		return false, err
	}
	return true, nil
}
