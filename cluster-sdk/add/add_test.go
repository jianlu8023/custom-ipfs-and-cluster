package add

import (
	"fmt"
	"testing"

	"ipfs-cluster"
)

func TestAddOneFile(t *testing.T) {
	if ok := ipfs_cluster.InitSdk("172.25.138.46", "", "", 9094, "TCP"); !ok {
		fmt.Println("init sdk error")
		return
	}

	file, err := OneFile(&Info{
		FileName:             "demo.txt",
		FilePath:             "./testdata/demo.txt",
		ReplicationFactorMin: 1,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file)

}
