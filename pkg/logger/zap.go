package logger

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
)

func NewZapLogger() *zap.Logger {
	l := zap.NewExample()
	return l
}

var _ log.Logger = (*Logger)(nil)

var DefaultLogger = NewLogger(NewZapLogger())

type Logger struct {
	Zlog *zap.Logger
}

func NewLogger(zlog *zap.Logger) *Logger {
	return &Logger{zlog}
}

func (l *Logger) Log(level log.Level, keyvals ...interface{}) error {
	for i, v := range keyvals {
		if r, ok := v.([]interface{}); ok {
			keyvals = append(keyvals[:i], r...)
		}
	}
	keylen := len(keyvals)
	if keylen == 0 || keylen%2 != 0 {
		l.Zlog.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}
	var msg string
	data := make([]zap.Field, 0, (keylen/2)+1)
	for i := 0; i < keylen; i += 2 {
		if keyvals[i] == log.DefaultMessageKey {
			msg = keyvals[i+1].(string)
		} else {
			data = append(data, zap.Any(fmt.Sprint(keyvals[i]), keyvals[i+1]))
		}
	}
	switch level {
	case log.LevelDebug:
		l.Zlog.Debug(msg, data...)
	case log.LevelInfo:
		l.Zlog.Info(msg, data...)
	case log.LevelWarn:
		l.Zlog.Warn(msg, data...)
	case log.LevelError:
		l.Zlog.Error(msg, data...)
	case log.LevelFatal:
		l.Zlog.Fatal(msg, data...)
	}
	return nil
}

func (l *Logger) Sync() error {
	return l.Zlog.Sync()
}

func (l *Logger) Close() error {
	return l.Sync()
}
