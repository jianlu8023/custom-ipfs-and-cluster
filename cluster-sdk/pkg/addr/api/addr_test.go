package api

import (
	"fmt"
	"testing"

	"tools/ipfs-cluster/config"
)

func TestGetAPIAddr(t *testing.T) {
	c := &config.Config{
		Host:     "127.0.0.1",
		Port:     9094,
		Protocol: "UDP",
		UserName: "",
		PassWord: "",
	}
	addr, err := GetAPIAddr(c)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("addr:", addr)
}
