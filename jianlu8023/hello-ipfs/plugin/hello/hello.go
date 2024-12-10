package hello

import (
	"fmt"

	"github.com/ipfs/kubo/plugin"
)

var (
	Plugins = []plugin.Plugin{
		&Plugin{},
	}
)

type Plugin struct{}

var _ plugin.Plugin = (*Plugin)(nil)

func (p *Plugin) Name() string {
	return "hello"
}

func (p *Plugin) Version() string {
	return "0.0.1"
}

func (p *Plugin) Init(env *plugin.Environment) error {
	fmt.Println(">>> hello plugin init")
	return nil
}

func (p *Plugin) Close() error {
	fmt.Println(">>> hello plugin close")
	return nil
}
