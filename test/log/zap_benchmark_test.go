package log

import (
	"go.uber.org/zap/zapcore"
	"sparrow/pkg/log/zaplog"
	"testing"
)

/*
------------------------压测-------------------------
*/
func BenchmarkZap(b *testing.B) {
	//time.Sleep(time.Second * 5)
	zaplog.InitWithEncoder("../output/", "testlog", zaplog.NewEncoder())
	for n := 0; n < b.N; n++ {
		zaplog.Logger.Log(zapcore.InfoLevel, "hello............")
	}
}

func BenchmarkZapSuper(b *testing.B) {
	//time.Sleep(time.Second * 5)
	zaplog.InitWithEncoder("../output/", "testlog", zaplog.NewEncoder())
	for n := 0; n < b.N; n++ {
		zaplog.LoggerSugar.Info("hello............")
	}
}
