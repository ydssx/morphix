package logger

import (
	"fmt"
	"os"

	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapLogger(callerSkip int) *zap.Logger {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.MessageKey = ""

	core := zapcore.NewCore(zapcore.NewJSONEncoder(cfg), zapcore.Lock(os.Stdout), zap.InfoLevel)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(callerSkip), zap.AddStacktrace(zap.DPanicLevel),zap.Hooks(func(e zapcore.Entry) error {
		return nil
	}))
	return logger
}

var _ log.Logger = (*Logger)(nil)

var DefaultLogger = NewLogger(NewZapLogger(2))

func init() {
	logger := log.With(NewLogger(NewZapLogger(3)))
	log.SetLogger(logger)
}

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
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		l.Zlog.Warn("Keyvalues must appear in pairs", zap.Any("keyvalues", keyvals))
		return nil
	}

	data := make([]zap.Field, 0, len(keyvals)/2)
	for i := 0; i < len(keyvals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), keyvals[i+1]))
	}

	l.Zlog.Log(zapcore.Level(level), "", data...)
	return nil
}

func (l *Logger) Sync() error {
	return l.Zlog.Sync()
}

func (l *Logger) Close() error {
	return l.Sync()
}
