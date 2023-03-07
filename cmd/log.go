package cmd

import (
	"context"
	"fmt"
	"io"

	elog "github.com/labstack/gommon/log"
	"github.com/scylladb/go-log"
)

// echoLogger implements the echo logger interface.
type echoLogger struct {
	base log.Logger
}

// newEchoLogger creates a new echo logger with a base logger.
func newEchoLogger(logger log.Logger) echoLogger {
	return echoLogger{base: logger}
}

// Output returns a discard writer since we don't expose the underlying writer.
func (l echoLogger) Output() io.Writer {
	return io.Discard
}

// Prefix returns an empty string to satisfy interface.
func (l echoLogger) Prefix() string {
	return ""
}

// SetOutput is a stub, don't allow echo to change our logger output.
func (l echoLogger) SetOutput(io.Writer) {}

// SetPrefix is a stub to satisfy interface.
func (l echoLogger) SetPrefix(string) {}

// SetHeader is a stub to satisfy interface.
func (l echoLogger) SetHeader(h string) {}

// Level is a stub to satisfy interface.
func (l echoLogger) Level() elog.Lvl {
	return elog.OFF
}

// SetLevel is a stub to satisfy interface.
func (l echoLogger) SetLevel(level elog.Lvl) {}

func (l echoLogger) Print(i ...interface{}) {
	l.base.Info(context.TODO(), fmt.Sprint(i...))
}

func (l echoLogger) Printf(format string, args ...interface{}) {
	l.base.Info(context.TODO(), fmt.Sprintf(format, args...))
}

func (l echoLogger) Printj(j elog.JSON) {
	l.base.Info(context.TODO(), "", jsonToKeyVals(j)...)
}

func (l echoLogger) Debug(i ...interface{}) {
	l.base.Debug(context.TODO(), fmt.Sprint(i...))
}

func (l echoLogger) Debugf(format string, args ...interface{}) {
	l.base.Debug(context.TODO(), fmt.Sprintf(format, args...))
}

func (l echoLogger) Debugj(j elog.JSON) {
	l.base.Debug(context.TODO(), "", jsonToKeyVals(j)...)
}

func (l echoLogger) Info(i ...interface{}) {
	l.base.Info(context.TODO(), fmt.Sprint(i...))
}

func (l echoLogger) Infof(format string, args ...interface{}) {
	l.base.Info(context.TODO(), fmt.Sprintf(format, args...))
}

func (l echoLogger) Infoj(j elog.JSON) {
	l.base.Info(context.TODO(), "", jsonToKeyVals(j)...)
}

func (l echoLogger) Warn(i ...interface{}) {
	l.Info(i...)
}

func (l echoLogger) Warnf(format string, args ...interface{}) {
	l.Infof(format, args...)
}

func (l echoLogger) Warnj(j elog.JSON) {
	l.Infoj(j)
}

func (l echoLogger) Error(i ...interface{}) {
	l.base.Error(context.TODO(), fmt.Sprint(i...))
}

func (l echoLogger) Errorf(format string, args ...interface{}) {
	l.base.Error(context.TODO(), fmt.Sprintf(format, args...))
}

func (l echoLogger) Errorj(j elog.JSON) {
	l.base.Error(context.TODO(), "", jsonToKeyVals(j)...)
}

func (l echoLogger) Fatal(i ...interface{}) {
	l.base.Fatal(context.TODO(), fmt.Sprint(i...))
}

func (l echoLogger) Fatalj(j elog.JSON) {
	l.base.Fatal(context.TODO(), "", jsonToKeyVals(j)...)
}

func (l echoLogger) Fatalf(format string, args ...interface{}) {
	l.base.Fatal(context.TODO(), fmt.Sprintf(format, args...))
}

func (l echoLogger) Panic(i ...interface{}) {
	l.base.Error(context.TODO(), fmt.Sprint(i...))
	panic("")
}

func (l echoLogger) Panicj(j elog.JSON) {
	l.base.Error(context.TODO(), "", jsonToKeyVals(j)...)
	panic("")
}

func (l echoLogger) Panicf(format string, args ...interface{}) {
	l.base.Error(context.TODO(), fmt.Sprintf(format, args...))
	panic("")
}

func jsonToKeyVals(j elog.JSON) (keyvals []interface{}) {
	for k, v := range j {
		keyvals = append(keyvals, k, v)
	}

	return
}
