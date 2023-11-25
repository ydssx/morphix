package logger

import (
	"fmt"
	"os"

	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapLogger() *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.MessageKey = ""

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.Lock(os.Stdout), zap.DebugLevel)

	logger := zap.New(core)
	return logger
}

var _ log.Logger = (*Logger)(nil)

var DefaultLogger = NewLogger(NewZapLogger())

func init() {
	logger := log.With(DefaultLogger, "caller", log.Caller(5))
	log.SetLogger(logger)
}

type Logger struct {
	Zlog *zap.Logger
}

func NewLogger(zlog *zap.Logger) *Logger {
	return &Logger{zlog}
}

func (l *Logger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		l.Zlog.Warn("Keyvalues must appear in pairs", zap.Any("keyvalues", keyvals))
		return nil
	}

	var data []zap.Field
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
