package zaplog

import (
	"errors"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"time"
)

var Logger *zap.Logger
var LoggerSugar *zap.SugaredLogger

func init() {
	_ = InitWithConfig(NewDevelopmentConfig())
}

func InitWithConfig(cfg zap.Config) (err error) {
	Logger, err = NewLogger(&cfg)
	if err != nil {
		return
	}
	LoggerSugar = Logger.Sugar()

	return
}

func NewDevelopmentConfig() zap.Config { return DefaultDevelopmentZapLoggerConfig }
func NewProductionConfig() zap.Config  { return DefaultZapLoggerConfig }

func NewLogger(cfg *zap.Config, option ...zap.Option) (*zap.Logger, error) {
	if cfg == nil {
		return nil, errors.New("nil zap.Config")
	}
	return cfg.Build(option...)
}

//	func SetLogger(l *zap.Logger) {
//		Logger = l
//		LoggerSugar = Logger.Sugar()
//	}
func NewEncoder() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:       "ts",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		// EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006/01/02 15:04:05.111"))
			//enc.AppendString(strconv.FormatInt(t.UnixNano()/1000, 10))
		},
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

//func NewDevelopmentEncoder() zapcore.EncoderConfig {
//	return zapcore.EncoderConfig{
//		TimeKey:       "ts",
//		LevelKey:      "level",
//		NameKey:       "logger",
//		CallerKey:     "caller",
//		MessageKey:    "msg",
//		StacktraceKey: "stacktrace",
//		LineEnding:    zapcore.DefaultLineEnding,
//		EncodeLevel:   zapcore.CapitalColorLevelEncoder,
//		// EncodeTime:     zapcore.ISO8601TimeEncoder,
//		EncodeTime:     zapcore.EpochTimeEncoder,
//		EncodeDuration: zapcore.StringDurationEncoder,
//		EncodeCaller:   zapcore.ShortCallerEncoder,
//	}
//}
//
//func NewProductionEncoder() zapcore.EncoderConfig {
//	return zapcore.EncoderConfig{
//		TimeKey:       "ts",
//		LevelKey:      "level",
//		NameKey:       "logger",
//		CallerKey:     "caller",
//		MessageKey:    "msg",
//		StacktraceKey: "stacktrace",
//		LineEnding:    zapcore.DefaultLineEnding,
//		EncodeLevel:   zapcore.CapitalColorLevelEncoder,
//		// EncodeTime:     zapcore.ISO8601TimeEncoder,
//		EncodeTime:     zapcore.EpochTimeEncoder,
//		EncodeDuration: zapcore.StringDurationEncoder,
//		EncodeCaller:   zapcore.ShortCallerEncoder,
//	}
//}

func getTrueInfoFileName(path string, fileName string) string {
	var dstStr string
	dstStr += path
	dstStr += fileName
	return dstStr + ".log"
}
func getTrueErrorFileName(path string, fileName string) string {
	var dstStr string
	dstStr += path
	dstStr += fileName
	return dstStr + "_error.log"
}

func getWriter(fileName string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		fileName+".%Y%m%d_%H", // 没有使用go风格反人类的format格式
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(time.Hour*24*30),
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		panic(err)
	}
	return hook
}

func InitWithEncoder(path string, fileName string, cfg zapcore.EncoderConfig) (err error) {
	encoder := zapcore.NewJSONEncoder(cfg) //zapcore.NewConsoleEncoder(cfg)

	//实现两个判断日志等级的interface (其实 zapcore.*Level 自身就是 interface)
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel
	})
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	infoWriter := getWriter(getTrueInfoFileName(path, fileName))
	warnWriter := getWriter(getTrueErrorFileName(path, fileName))

	//创建logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
	)
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	LoggerSugar = Logger.Sugar()

	return nil
}
