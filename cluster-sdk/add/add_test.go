package add

import (
	"fmt"
	"testing"

	"ipfs-cluster"
)

func TestAddOneFile(t *testing.T) {
	if ok := ipfs_cluster.InitSdk("127.0.0.1", "", "", 9094, "TCP"); !ok {
		fmt.Println("init sdk error")
		return
	}

	file, err := OneFile(&Info{
		FileName:             "demo.txt",
		FilePath:             "/home/user/go/src/gitee.com/jianlu8023/modify-ipfs-with-cluster/cluster-sdk/testdata/demo.txt",
		ReplicationFactorMax: 1,
		ReplicationFactorMin: 1,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file)

}
