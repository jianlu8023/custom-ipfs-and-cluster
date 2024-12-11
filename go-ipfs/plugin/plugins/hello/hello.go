package hello

import (
	"fmt"

	coreiface "github.com/ipfs/kubo/core/coreiface"
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

func (*helloPlugin) Init(*plugin.Environment) error {
	fmt.Println("hello plugin init")
	return nil
}

func (*helloPlugin) Start(coreiface.CoreAPI) error {
	fmt.Println("hello plugin start")
	return nil
}

func (*helloPlugin) Close() error {
	fmt.Println("hello plugin close")
	return nil
}
