package logger

import (
	"github.com/ipfs/go-log/v2"
	logging "github.com/ipfs/go-log/v2"
)

var (
	addLogger    = logging.Logger("add")
	catLogger    = logging.Logger("cat")
	getLogger    = logging.Logger("get")
	cryptoLogger = logging.Logger("crypto")
)

func GetAddLogger() *log.ZapEventLogger {
	return addLogger
}

func GetCatLogger() *log.ZapEventLogger {
	return catLogger
}

func GetGetLogger() *log.ZapEventLogger {
	return getLogger
}

func GetCryptoLogger() *log.ZapEventLogger {
	return cryptoLogger
}
