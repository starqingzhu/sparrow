package log

import (
	"sparrow/pkg/log/zaplog"
	"testing"
)

func TestZap(t *testing.T) {
	zaplog.InitWithEncoder("../output/", "testlog", zaplog.NewEncoder())
	zaplog.LoggerSugar.Info("hello", "world")

}
