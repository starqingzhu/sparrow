package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})

	Error(v ...interface{})
	Errorf(format string, v ...interface{})

	Info(v ...interface{})
	Infof(format string, v ...interface{})

	Warning(v ...interface{})
	Warningf(format string, v ...interface{})

	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})

	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
}

//func LibLogger() Logger {
//	return libLogger
//}
//
//func SetLibLogger(l Logger) {
//	libLogger = l
//}

var (
	defaultLogger = &DefaultLogger{Logger: log.New(os.Stderr, "internal ", log.LstdFlags)}
	discardLogger = &DefaultLogger{Logger: log.New(ioutil.Discard, "", 0)}
	libLogger     = Logger(defaultLogger)
)

const (
	callDepth = 2
)

// DefaultLogger is a default implementation of the Logger interface.
type DefaultLogger struct {
	*log.Logger
	debug bool
}

func (l *DefaultLogger) EnableTimestamps() {
	l.SetFlags(l.Flags() | log.Ldate | log.Ltime)
}

func (l *DefaultLogger) EnableDebug() {
	l.debug = true
}

func prefix(lvl, msg string) string {
	return fmt.Sprintf("%s: %s", lvl, msg)
}

func (l *DefaultLogger) Debug(v ...interface{}) {
	if l.debug {
		_ = l.Output(callDepth, prefix("DEBUG", fmt.Sprint(v...)))
	}
}

func (l *DefaultLogger) Debugf(format string, v ...interface{}) {
	if l.debug {
		_ = l.Output(callDepth, prefix("DEBUG", fmt.Sprintf(format, v...)))
	}
}

func (l *DefaultLogger) Info(v ...interface{}) {
	_ = l.Output(callDepth, prefix("INFO", fmt.Sprint(v...)))
}

func (l *DefaultLogger) Infof(format string, v ...interface{}) {
	_ = l.Output(callDepth, prefix("INFO", fmt.Sprintf(format, v...)))
}

func (l *DefaultLogger) Error(v ...interface{}) {
	_ = l.Output(callDepth, prefix("ERROR", fmt.Sprint(v...)))
}

func (l *DefaultLogger) Errorf(format string, v ...interface{}) {
	_ = l.Output(callDepth, prefix("ERROR", fmt.Sprintf(format, v...)))
}

func (l *DefaultLogger) Warning(v ...interface{}) {
	_ = l.Output(callDepth, prefix("WARN", fmt.Sprint(v...)))
}

func (l *DefaultLogger) Warningf(format string, v ...interface{}) {
	_ = l.Output(callDepth, prefix("WARN", fmt.Sprintf(format, v...)))
}

func (l *DefaultLogger) Fatal(v ...interface{}) {
	_ = l.Output(callDepth, prefix("FATAL", fmt.Sprint(v...)))
	os.Exit(1)
}

func (l *DefaultLogger) Fatalf(format string, v ...interface{}) {
	_ = l.Output(callDepth, prefix("FATAL", fmt.Sprintf(format, v...)))
	os.Exit(1)
}

func (l *DefaultLogger) Panic(v ...interface{}) {
	l.Logger.Panic(v...)
}

func (l *DefaultLogger) Panicf(format string, v ...interface{}) {
	l.Logger.Panicf(format, v...)
}
