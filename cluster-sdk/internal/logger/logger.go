package logger

import (
	glog "github.com/jianlu8023/go-logger"
	"go.uber.org/zap"
)

var ipfsLogger = glog.NewSugaredLogger(&glog.Config{
	LogLevel:    "debug",
	DevelopMode: false,
	StackLevel:  "",
	ModuleName:  "[IPFS]",
	Caller:      false,
},
	glog.WithConsoleFormat(),
)

func GetIPFSLogger() *zap.SugaredLogger {
	return ipfsLogger
}
