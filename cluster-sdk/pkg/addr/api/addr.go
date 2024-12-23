package api

import (
	"fmt"

	"github.com/multiformats/go-multiaddr"
	"ipfs-cluster/internal/config"
	"ipfs-cluster/internal/logger"
)

const (
	ip4Tcp           = "/ip4/%v/tcp/%v"
	ip4Quic          = "/ip4/%v/udp/%v/quic"
	proxyTcp         = "/ip4/%v/tcp/%v"
	proxyDefaultPort = 9095
	tcp              = "TCP"
	udp              = "UDP"
)

// GetAPIAddr 获取API地址
// @param config ipfs-cluster配置
// @return multiaddr.Multiaddr API地址
// @return error 错误信息
func GetAPIAddr(config *config.Config) (multiaddr.Multiaddr, error) {

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
		return nil, fmt.Errorf("unsupported protocol: %s", ptcl)
	}

	return newMultiAddr(url)
}

// GetProxyAddr 获取代理地址
// @param config ipfs-cluster配置
// @return multiaddr.Multiaddr 代理地址
// @return error 错误信息
func GetProxyAddr(config *config.Config) (multiaddr.Multiaddr, error) {

	dst := make([]byte, len(proxyTcp))
	copy(dst, proxyTcp)
	url := fmt.Sprintf(string(dst), config.Host, proxyDefaultPort)

	return newMultiAddr(url)
}

// newMultiAddr 创建MultiAddress
// @param addr 地址
// @return multiaddr.Multiaddr multiaddress
// @return error 错误信息
func newMultiAddr(addr string) (multiaddr.Multiaddr, error) {
	apiMultiAddr, err := multiaddr.NewMultiaddr(addr)
	if err != nil {
		logger.GetIPFSLogger().Errorf("convert to multiAddress error %v", err)
		return nil, err
	}
	return apiMultiAddr, nil
}
