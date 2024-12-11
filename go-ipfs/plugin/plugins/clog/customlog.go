package clog

import (
	"encoding/json"
	"fmt"

	logging "github.com/ipfs/go-log/v2"
	"github.com/ipfs/kubo/plugin"
)

var Plugins = []plugin.Plugin{
	&clogPlugin{},
}

const defaultLoggerLevel = "error"

var _ plugin.Plugin = &clogPlugin{}

type loggerConfig struct {
	Levels       map[string][]string `json:"Levels"`
	DefaultLevel string              `json:"DefaultLevel"`
}

type clogPlugin struct {
	conf loggerConfig
}

func (l *clogPlugin) Name() string {
	return "datadog-logger"
}

func (l *clogPlugin) Version() string {
	return "0.0.1"
}

// Set log levels for each system (logger)
func (l *clogPlugin) Init(env *plugin.Environment) error {
	err := l.loadConfig(env)
	if err != nil {
		return err
	}

	// set default log level for all loggers
	defaultLevel, err := logging.LevelFromString(l.conf.DefaultLevel)
	if err != nil {
		return err
	}
	logging.SetAllLoggers(defaultLevel)

	for level, subsystems := range l.conf.Levels {
		fmt.Println(">>> set log level: ", level, " for subsystems: ", subsystems)
		for _, subsystem := range subsystems {
			fmt.Println(">>> set log level for subsystem: ", subsystem)
			if err := logging.SetLogLevel(subsystem, level); err != nil {
				return fmt.Errorf("set log level failed for subsystem: %s. Error: %s", subsystem, err.Error())
			}
		}
	}

	return nil
}

func (l *clogPlugin) loadConfig(env *plugin.Environment) error {
	// load config data
	bytes, err := json.Marshal(env.Config)
	if err != nil {
		return err
	}

	fmt.Println(">>> load config data: ", string(bytes))

	if err = json.Unmarshal(bytes, &l.conf); err != nil {
		return err
	}
	if l.conf.DefaultLevel == "" {
		l.conf.DefaultLevel = defaultLoggerLevel
	}
	return nil
}
