package ipfs_cluster

import (
	"strconv"

	rclient "github.com/ipfs-cluster/ipfs-cluster/api/rest/client"
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
	logger.GetIPFSLogger().Debugf("starting connect ipfs cluster")
	// 初始化sdk
	var err error
	c := &config.Config{
		Host:     host,
		Port:     port,
		Protocol: protocol,
		UserName: userName,
		PassWord: passWord,
	}
	apiMultiAddr, err := api.GetAPIAddr(c)
	if err != nil {
		logger.GetIPFSLogger().Errorf(">>> 解析API地址生成API MultiAddress失败 %v", err)
		return false
	}

	proxyMultiAddr, err := api.GetProxyAddr(c)
	if err != nil {
		logger.GetIPFSLogger().Errorf(">>> 解析API地址生成PROXY MultiAddress失败 %v", err)
		return false
	}
	s, err := rclient.NewLBClient(
		&rclient.Failover{},
		[]*rclient.Config{
			{
				Host:              host,
				Port:              strconv.Itoa(port),
				DisableKeepAlives: true,
				Password:          c.PassWord,
				Username:          c.UserName,
				APIAddr:           apiMultiAddr,
				ProxyAddr:         proxyMultiAddr,
				// LogLevel:          "debug",
				LogLevel: "info",
				Timeout:  config.Timeout120,
			},
		},
		config.Retries)
	sdk.SetSDK(s)
	err = version.Version()
	if err != nil {
		logger.GetIPFSLogger().Errorf(">>> 获取版本信息失败 %v", err)
		return false
	}
	logger.GetIPFSLogger().Debugf("connect ipfs cluster success")
	return true
}
