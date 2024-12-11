package hello

import (
	"fmt"

	"github.com/ipfs/kubo/plugin"
)

// Plugins is exported list of plugins that will be loaded.
var Plugins = []plugin.Plugin{
	&helloPlugin{},
}

type helloPlugin struct{}

var _ plugin.Plugin = (*helloPlugin)(nil)

func (*helloPlugin) Name() string {
	return "hello"
}

func (*helloPlugin) Version() string {
	return "0.0.1"
}

func (*helloPlugin) Init(env *plugin.Environment) error {
	fmt.Println("hello plugin init")
	fmt.Println("env ", env)
	return nil
}
