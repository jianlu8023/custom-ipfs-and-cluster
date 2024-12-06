package logger

import (
	"github.com/ipfs/go-log/v2"
	logging "github.com/ipfs/go-log/v2"
)

var (
	cryptoLogger = logging.Logger("crypto")
)

func GetCryptoLogger() *log.ZapEventLogger {
	return cryptoLogger
}
