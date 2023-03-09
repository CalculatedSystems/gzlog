package log

import (
	"io"

	"go.uber.org/zap"
)

func Fatal(args ...any) {
	l.Fatal(args...)
}

// Fatalf wraps the underlying configured loggers Fatalf function
func Fatalf(template string, args ...any) {
	l.Fatalf(template, args...)
}

// Fatalln wraps the underlying configured loggers Fatalln function
func Fatalln(args ...any) {
	l.Fatalln(args...)
}

// Fatalw wraps the underlying configured loggers Fatalw function
func Fatalw(msg string, keysAndValues ...any) {
	l.Fatalw(msg, keysAndValues...)
}

// Panic wraps the underlying configured loggers Panic function
func Panic(args ...any) {
	l.Panic(args...)
}

// Panicf wraps the underlying configured loggers Panicf function
func Panicf(template string, args ...any) {
	l.Panicf(template, args...)
}

// Panicln wraps the underlying configured loggers Panicln function
func Panicln(args ...any) {
	l.Panicln(args...)
}

// Panicw wraps the underlying configured loggers Panicw function
func Panicw(msg string, keysAndValues ...any) {
	l.Panicw(msg, keysAndValues...)
}

// DPanic wraps the underlying configured loggers DPanic function
func DPanic(args ...any) {
	l.DPanic(args...)
}

// DPanicf wraps the underlying configured loggers DPanicf function
func DPanicf(template string, args ...any) {
	l.DPanicf(template, args...)
}

// DPanicln wraps the underlying configured loggers DPanicln function
func DPanicln(args ...any) {
	l.DPanicln(args...)
}

// DPanicw wraps the underlying configured loggers DPanicw function
func DPanicw(msg string, keysAndValues ...any) {
	l.DPanicw(msg, keysAndValues...)
}

// Error wraps the underlying configured loggers Error function
func Error(args ...any) {
	l.Error(args...)
}

// Errorf wraps the underlying configured loggers Errorf function
func Errorf(template string, args ...any) {
	l.Errorf(template, args...)
}

// Errorln wraps the underlying configured loggers Errorln function
func Errorln(args ...any) {
	l.Errorln(args...)
}

// Errorw wraps the underlying configured loggers Errorw function
func Errorw(msg string, keysAndValues ...any) {
	l.Errorw(msg, keysAndValues...)
}

// Warn wraps the underlying configured loggers Warn function
func Warn(args ...any) {
	l.Warn(args...)
}

// Warnf wraps the underlying configured loggers Warnf function
func Warnf(template string, args ...any) {
	l.Warnf(template, args...)
}

// Warnln wraps the underlying configured loggers Warnln function
func Warnln(args ...any) {
	l.Warnln(args...)
}

// Warnw wraps the underlying configured loggers Warnw function
func Warnw(msg string, keysAndValues ...any) {
	l.Warnw(msg, keysAndValues...)
}

// Info wraps the underlying configured loggers Info function
func Info(args ...any) {
	l.Info(args...)
}

// Infof wraps the underlying configured loggers Infof function
func Infof(template string, args ...any) {
	l.Infof(template, args...)
}

// Infoln wraps the underlying configured loggers Infoln function
func Infoln(args ...any) {
	l.Infoln(args...)
}

// Infow wraps the underlying configured loggers Infow function
func Infow(msg string, keysAndValues ...any) {
	l.Infow(msg, keysAndValues...)
}

// Debug wraps the underlying configured loggers Debug function
func Debug(args ...any) {
	l.Debug(args...)
}

// Debugf wraps the underlying configured loggers Debugf function
func Debugf(template string, args ...any) {
	l.Debugf(template, args...)
}

// Debugln wraps the underlying configured loggers Debugln function
func Debugln(args ...any) {
	l.Debugln(args...)
}

// Debugw wraps the underlying configured loggers Debugw function
func Debugw(msg string, keysAndValues ...any) {
	l.Debugw(msg, keysAndValues...)
}

// standard library wrappers

// Print wraps the underlying configured loggers Info function
func Print(args ...any) {
	l.Info(args...)
}

// Printf wraps the underlying configured loggers Infof function
func Printf(template string, args ...any) {
	l.Infof(template, args...)
}

// Println wraps the underlying configured loggers Infoln function
func Println(args ...any) {
	l.Infoln(args...)
}

// Printw wraps the underlying configured loggers Infow function
func Printw(msg string, keysAndValues ...any) {
	l.Infow(msg, keysAndValues...)
}

// Writer returns the internal `zap.Logger` writer
func Writer() io.Writer {
	return zap.NewStdLog(l.Desugar()).Writer()
}
