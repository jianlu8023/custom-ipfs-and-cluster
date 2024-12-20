package logger

import (
	glog "github.com/jianlu8023/go-logger"
	"go.uber.org/zap"
)

var IPFSLogger = glog.NewSugaredLogger(&glog.Config{
	LogLevel:    "debug",
	DevelopMode: false,
	StackLevel:  "error",
	ModuleName:  "[IPFS]",
	Caller:      true,
},
	glog.WithConsoleFormat(),
)

func GetIPFSLogger() *zap.SugaredLogger {
	return IPFSLogger
}
