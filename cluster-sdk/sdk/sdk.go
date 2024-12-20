package sdk

import (
	rclient "github.com/ipfs-cluster/ipfs-cluster/api/rest/client"
)

var (
	sdk rclient.Client
)

// GetSDK 获取sdk
// @return rclient.Client sdk
func GetSDK() rclient.Client {
	return sdk
}

// SetSDK 设置sdk
// @param s rclient.Client
func SetSDK(s rclient.Client) {
	sdk = s
}
