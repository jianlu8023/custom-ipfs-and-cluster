package verify

import (
	"errors"
	"fmt"
	"net"

	"github.com/ipfs-cluster/ipfs-cluster/api"
	"ipfs-cluster/internal/logger"

	"ipfs-cluster/internal/config"
)

const (
	localAddr = "127.0.0.1:9096"
)

// ValidAddr 验证地址有效性
// 返回第一个有效的地址
// @param address []api.Multiaddr 地址列表
// @return string 第一个有效的地址
// @return error 错误信息
func ValidAddr(address []api.Multiaddr) (string, error) {
	valid := make([]string, 0)
	for _, addr := range address {
		ip, err := addr.ValueForProtocol(4)
		if err != nil {
			logger.GetIPFSLogger().Errorf(">>> 解析MultiAddress失败 %v", err)
			return "", err
		}
		port, err := addr.ValueForProtocol(6)
		if err != nil {
			logger.GetIPFSLogger().Errorf(">>> 解析MultiAddress失败 %v", err)
			return "", err
		}
		if fmt.Sprintf("%s:%s", ip, port) == localAddr {
			logger.GetIPFSLogger().Debugf(">>> 跳过本机地址")
			continue
		}
		if checkIp(ip, port) {
			valid = append(valid, ip)
		}
	}
	if len(valid) == 0 {
		return "", errors.New("No Valid Address ")
	}
	return valid[0], nil
}

// checkIp 检查ip地址是否有效
// @param ip string ip地址
// @param port string 端口号
// @return bool 是否有效
// @example checkIp("127.0.0.1", "53") => true
func checkIp(ip string, port string) bool {
	addr := net.JoinHostPort(ip, port)
	dial, err := net.DialTimeout("tcp", addr, config.Timeout3)
	if err == nil {
		defer func(dial net.Conn) {
			err := dial.Close()
			if err != nil {
				logger.GetIPFSLogger().Errorf(">>> 连接 %v 失败 %v", addr, err)
			}
		}(dial)
		return true
	}
	return false
}
