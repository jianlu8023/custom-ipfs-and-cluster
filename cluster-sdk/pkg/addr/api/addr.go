package api

import (
	"fmt"

	"ipfs-cluster/internal/config"
)

const (
	ip4Tcp  = "/ip4/%v/tcp/%v"
	ip4Quic = "/ip4/%v/udp/%v/quic"
	tcp     = "TCP"
	udp     = "UDP"
)

// GetAPIAddr 获取API地址
// @param config ipfs-cluster配置
// @return string API地址
// @return error 错误信息
func GetAPIAddr(config *config.Config) (string, error) {

	var url string
	switch ptcl := config.Protocol; ptcl {
	case tcp:
		dst := make([]byte, len(ip4Tcp))
		copy(dst, ip4Tcp)
		url = fmt.Sprintf(string(dst), config.Host, config.Port)
	case udp:
		dst := make([]byte, len(ip4Quic))
		copy(dst, ip4Quic)
		url = fmt.Sprintf(string(dst), config.Host, config.Port)
	default:
		return "", fmt.Errorf("unsupported protocol: %s", ptcl)
	}
	return url, nil
}
