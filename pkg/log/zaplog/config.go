package zaplog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*
zap配置说明
Level：日志级别,只不过它需要的类型是AtomicLevel。所以需要使用zap.NewAtomicLevelAt。
	func NewAtomicLevelAt(l zapcore.Level) AtomicLevel {
		a := NewAtomicLevel()
		a.SetLevel(l)
		return a
	}

	DebugLevel Level = iota - 1
	// InfoLevel 默认级别
	InfoLevel
	// WarnLevel
	WarnLevel
	// 项目运行中，错误的日志
	ErrorLevel
	// 记录一条消息后，如果是开发环境，则调用panic
	DPanicLevel
	// 记录一条消息后，则调用panic
	PanicLevel
	// 记录一条消息后调用 os.Exit(1).
	FatalLevel
	Development：是否是开发环境。如果是开发模式，对DPanicLevel进行堆栈跟踪
	DisableCaller： 不显示调用函数的文件名称和行号。默认情况下，所有日志都显示。
	DisableStacktrace：是否禁用堆栈跟踪捕获。默认对Warn级别以上和生产error级别以上的进行堆栈跟踪。
	Sampling：抽样策略。设置为nil禁用采样。
	Encoding：编码方式，支持json, console, 也支持自定义（需要通过RegisterEncoder注册）
	EncoderConfig： 生成格式的一些配置，后面详细介绍
	OutputPaths： 是url或文件路径列表，用于写入日志输出。
	ErrorOutputPaths：内部日志程序错误路径列表。默认为标准错误。
	InitialFields：加入一些初始的字段数据，比如项目名
	EncoderConfig配置说明:
		MessageKey：输入信息的key名
		LevelKey：输出日志级别的key名
		TimeKey：输出时间的key名
		NameKey CallerKey StacktraceKey跟以上类似，看名字就知道
		LineEnding：每行的分隔符。基本zapcore.DefaultLineEnding 即"\\n"
		EncodeLevel：基本zapcore.LowercaseLevelEncoder。将日志级别字符串转化为小写
		EncodeTime：输出的时间格式
		EncodeDuration：一般zapcore.SecondsDurationEncoder,执行消耗的时间转化成浮点型的秒
		EncodeCaller：一般zapcore.ShortCallerEncoder，以包/文件:行号 格式化调用堆栈
		EncodeName：可选值。
*/

var DefaultDevelopmentZapLoggerConfig = zap.Config{
	Level:       zap.NewAtomicLevelAt(ConvertToZapLevel(DefaultLogLevel)),
	Development: true,
	// Encoding:    "console",
	Encoding: "json",
	EncoderConfig: zapcore.EncoderConfig{
		// copied from "zap.NewDevelopmentEncoderConfig" with some updates
		TimeKey:       "ts",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		// EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		// EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeTime:     zapcore.EpochMillisTimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	},
	OutputPaths:      []string{"stdout"},
	ErrorOutputPaths: []string{"stderr"},
}

// DefaultZapLoggerConfig defines default zap logger configuration.
var DefaultZapLoggerConfig = zap.Config{
	Level:       zap.NewAtomicLevelAt(ConvertToZapLevel(DefaultLogLevel)),
	Development: false,
	Encoding:    "json",

	// copied from "zap.NewProductionEncoderConfig" with some updates
	EncoderConfig: zapcore.EncoderConfig{
		TimeKey:       "ts",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		// EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeTime:     zapcore.EpochMillisTimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	},
	// Use "/dev/null" to discard all
	OutputPaths:      []string{"stderr"},
	ErrorOutputPaths: []string{"stderr"},
}

type WitchLogOptFunc func(option *zap.Config)

// 根据传入配置初始化日志
func InitWithConfigFunc(opts ...WitchLogOptFunc) (err error) {
	defConf := &DefaultDevelopmentZapLoggerConfig
	for _, f := range opts {
		f(defConf)
	}
	Logger, err = NewLogger(defConf)
	if err != nil {
		return
	}
	LoggerSugar = Logger.Sugar()
	return
}

// 设置日志级别
// debug/info/warn/error/dpanic/panic/fatal
func WithLogLevel(logLevel string) WitchLogOptFunc {
	return func(option *zap.Config) {
		option.Level = zap.NewAtomicLevelAt(ConvertToZapLevel(logLevel))
	}
}

// 是否为开发模式
func WithIsDevelop(dev bool) WitchLogOptFunc {
	return func(option *zap.Config) {
		option.Development = dev
	}
}
