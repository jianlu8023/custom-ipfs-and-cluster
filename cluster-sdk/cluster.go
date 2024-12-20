package ipfs_cluster

import (
	rclient "github.com/ipfs-cluster/ipfs-cluster/api/rest/client"
	"github.com/multiformats/go-multiaddr"
	"ipfs-cluster/internal/config"
	"ipfs-cluster/internal/logger"

	"ipfs-cluster/pkg/addr/api"
	"ipfs-cluster/sdk"
	"ipfs-cluster/version"
)

// InitSdk 初始化sdk
// @param host string ipfs-cluster主机
// @param userName string 用户名
// @param passWord string 密码
// @param port int 端口
// @param protocol string 协议 TCP/UDP
// @return bool 是否成功
func InitSdk(host, userName, passWord string, port int, protocol string) bool {
	logger.GetIPFSLogger().Debugf(">>> 开始尝试连接IPFS CLUSTER集群")
	// 初始化sdk
	var err error
	c := &config.Config{
		Host:     host,
		Port:     port,
		Protocol: protocol,
		UserName: userName,
		PassWord: passWord,
	}
	apiAddr, err := api.GetAPIAddr(c)
	if err != nil {
		logger.GetIPFSLogger().Errorf(">>> 获取解析API地址失败 %v", err)
		return false
	}
	newMultiAddr, err := multiaddr.NewMultiaddr(apiAddr)
	if err != nil {
		logger.GetIPFSLogger().Errorf(">>> 解析API地址生成MultiAddress失败 %v", err)
		return false
	}
	s, err := rclient.NewDefaultClient(&rclient.Config{
		DisableKeepAlives: true,
		Password:          c.PassWord,
		Username:          c.UserName,
		APIAddr:           newMultiAddr,
	})
	sdk.SetSDK(s)
	err = version.Version()
	if err != nil {
		logger.GetIPFSLogger().Errorf(">>> 获取版本信息失败 %v", err)
		return false
	}

	return true
}
