package zaplog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sparrow/pkg/log"
	"strings"
)

var DefaultLogLevel = "debug"

// ConvertToZapLevel converts log level string to zapcore.Level.
func ConvertToZapLevel(lvl string) zapcore.Level {
	lvl = strings.ToLower(lvl)
	switch lvl {
	case log.Debug:
		return zap.DebugLevel
	case log.Info:
		return zap.InfoLevel
	case log.Warn:
		return zap.WarnLevel
	case log.Lerror:
		return zap.ErrorLevel
	case log.Dpanic:
		return zap.DPanicLevel
	case log.Panic:
		return zap.PanicLevel
	case log.Fatal:
		return zap.FatalLevel
	default:
		panic(fmt.Sprintf("unknown level %q", lvl))
	}
}
