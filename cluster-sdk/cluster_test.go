package ipfs_cluster

import (
	"fmt"
	"testing"
)

func TestInitSdk(t *testing.T) {
	ok := InitSdk("172.25.138.45", "", "", 9094, "TCP")

	fmt.Println(ok)
}
