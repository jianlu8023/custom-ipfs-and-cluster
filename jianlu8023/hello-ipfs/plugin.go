package hello_ipfs

import (
	"gitee.com/jianlu8023/hello-ipfs/plugin/hello"
	"github.com/ipfs/kubo/plugin"
)

var Plugins = []plugin.Plugin{
	&hello.Plugin{},
}
