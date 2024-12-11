package clog

import (
	"encoding/json"
	"fmt"

	logging "github.com/ipfs/go-log/v2"
	"github.com/ipfs/kubo/plugin"
)

var log = logging.Logger("plugin/customlog")

var Plugins = []plugin.Plugin{
	&customLogLevelPlugin{},
}

const defaultLoggerLevel = "error"

var _ plugin.Plugin = &customLogLevelPlugin{}

type loggerConfig struct {
	Levels       map[string][]string `json:"Levels"`
	DefaultLevel string              `json:"DefaultLevel"`
}

type customLogLevelPlugin struct {
	conf loggerConfig
}

func (l *customLogLevelPlugin) Name() string {
	return "datadog-logger"
}

func (l *customLogLevelPlugin) Version() string {
	return "0.0.1"
}

// Init Set log levels for each system (logger)
func (l *customLogLevelPlugin) Init(env *plugin.Environment) error {
	log.Debugf("starting init custom log plugin")
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
		// log.Debugf("setting level %v for subsystems %v", level, subsystems)
		for _, subsystem := range subsystems {
			log.Debugf("setting level %v for subsystem %v", level, subsystem)
			if err := logging.SetLogLevel(subsystem, level); err != nil {
				return fmt.Errorf("set log level failed for subsystem: %s. Error: %s", subsystem, err.Error())
			}
		}
	}

	return nil
}

func (l *customLogLevelPlugin) loadConfig(env *plugin.Environment) error {
	// load config data
	bytes, err := json.Marshal(env.Config)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(bytes, &l.conf); err != nil {
		return err
	}
	log.Debugf("loaded plugin config data %v", l.conf)
	if l.conf.DefaultLevel == "" {
		log.Debugf("default log level not set, setting to %v", defaultLoggerLevel)
		l.conf.DefaultLevel = defaultLoggerLevel
	}
	return nil
}
