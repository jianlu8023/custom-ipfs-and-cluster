package download

import (
	"testing"

	ipfs_cluster "ipfs-cluster"
)

func TestOneFile(t *testing.T) {
	ipfs_cluster.InitSdk("127.0.0.1", "", "", 9094, "UDP")
	t.Log("开始测试下载文件")
	err := OneFile("bafybeiejwdugvwzpw64tbgvkjaonpamvyd6aanbjg5x2cb5wo4hxvb6ehe",
		"/home/user/go/src/gitee.com/jianlu8023/modify-ipfs-with-cluster/cluster-sdk/testdata/demo_download.txt")
	if err != nil {
		t.Errorf("下载文件失败: %v", err)
	}
}
