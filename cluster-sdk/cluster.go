package ipfs_cluster

import (
	"slices"
	"strconv"

	rclient "github.com/ipfs-cluster/ipfs-cluster/api/rest/client"
	"ipfs-cluster/internal/config"
	"ipfs-cluster/internal/logger"
	"ipfs-cluster/pkg/addr/api"
	"ipfs-cluster/pkg/str"

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

// InitSDKWithMultiClient 使用多配置初始化sdk
// @param strategy 使用的策略
// @param cc 节点配置
// @return bool 是否成功
func InitSDKWithMultiClient(strategy string, cc []struct {
	Host     string
	UserName string
	PassWord string
	Port     int
}) bool {
	logger.GetIPFSLogger().Infof(">>> 开始初始化IPFS CLUSTER SDK ...")
	// 初始化sdk
	var err error
	var s rclient.Client
	clients := make([]*rclient.Config, 0, len(cc))

	for _, v := range cc {
		clients = append(clients, &rclient.Config{
			Host:              v.Host,
			Port:              strconv.Itoa(v.Port),
			Username:          v.UserName,
			Password:          v.PassWord,
			DisableKeepAlives: true,
			LogLevel:          "debug",
			Timeout:           config.Timeout120,
		})
	}

	clients = slices.DeleteFunc(clients, func(r *rclient.Config) bool {
		return r == nil || r.Host == "" || r.Port == ""
	})
	if str.CompareIgnoreCase("Failover", strategy) {
		logger.GetIPFSLogger().Debugf(">>> 使用 故障转移 策略 ...")
		s, err = rclient.NewLBClient(&rclient.Failover{}, clients, config.Retries)
	} else if str.CompareIgnoreCase("RoundRobin", strategy) {
		logger.GetIPFSLogger().Debugf(">>> 使用 节点轮询 策略 ...")
		s, err = rclient.NewLBClient(&rclient.RoundRobin{}, clients, config.Retries)
	} else {
		logger.GetIPFSLogger().Errorf(">>> 未知策略")
		return false
	}

	if err != nil {
		logger.GetIPFSLogger().Errorf(">>> 创建IPFS CLUSTER客户端失败 %v", err)
		return false
	}

	sdk.SetSDK(s)
	if err = version.Version(); err != nil {
		logger.GetIPFSLogger().Errorf(">>> 获取版本信息失败 %v", err)
		return false
	}

	return true
}
