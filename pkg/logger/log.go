package logger

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/trace"
)

var logTraceID = func(ctx context.Context) []interface{} {
	if traceID, ok := TraceIDFromContext(ctx); ok {
		return []interface{}{"traceID", traceID}
	}
	return nil
}

func TraceIDFromContext(ctx context.Context) (string, bool) {
	if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
		return span.TraceID().String(), true
	}
	return "", false
}

func Info(ctx context.Context, msg ...interface{}) {
	kv := []interface{}{log.DefaultMessageKey, fmt.Sprint(msg...)}
	DefaultLogger.Log(log.LevelInfo, append(kv, logTraceID(ctx)...))
}

func Infof(ctx context.Context, format string, msg ...interface{}) {
	kv := []interface{}{log.DefaultMessageKey, fmt.Sprintf(format, msg...)}
	DefaultLogger.Log(log.LevelInfo, append(kv, logTraceID(ctx)...))
}

func Error(ctx context.Context, msg ...interface{}) {
	kv := []interface{}{log.DefaultMessageKey, fmt.Sprint(msg...)}
	DefaultLogger.Log(log.LevelError, append(kv, logTraceID(ctx)...))
}

func Errorf(ctx context.Context, format string, msg ...interface{}) {
	kv := []interface{}{log.DefaultMessageKey, fmt.Sprintf(format, msg...)}
	DefaultLogger.Log(log.LevelError, append(kv, logTraceID(ctx)...))
}

func Fatalf(ctx context.Context, format string, msg ...interface{}) {
	kv := []interface{}{log.DefaultMessageKey, fmt.Sprintf(format, msg...)}
	DefaultLogger.Log(log.LevelFatal, append(kv, logTraceID(ctx)...))
}

func Fatal(ctx context.Context, msg ...interface{}) {
	kv := []interface{}{log.DefaultMessageKey, fmt.Sprint(msg...)}
	DefaultLogger.Log(log.LevelFatal, append(kv, logTraceID(ctx)...))
}

func Log(ctx context.Context, level log.Level, msg string, keyvals ...interface{}) {
	kv := []interface{}{log.DefaultMessageKey, msg}
	kv = append(kv, logTraceID(ctx)...)
	DefaultLogger.Log(level, append(kv, keyvals...))
}

func Debug(ctx context.Context, msg ...interface{}) {
	kv := []interface{}{log.DefaultMessageKey, fmt.Sprint(msg...)}
	DefaultLogger.Log(log.LevelDebug, append(kv, logTraceID(ctx)...))
}

func Debugf(ctx context.Context, format string, msg ...interface{}) {
	kv := []interface{}{log.DefaultMessageKey, fmt.Sprintf(format, msg...)}
	DefaultLogger.Log(log.LevelDebug, append(kv, logTraceID(ctx)...))
}

func Warn(ctx context.Context, msg ...interface{}) {
	kv := []interface{}{log.DefaultMessageKey, fmt.Sprint(msg...)}
	DefaultLogger.Log(log.LevelWarn, append(kv, logTraceID(ctx)...))
}

func Warnf(ctx context.Context, format string, msg ...interface{}) {
	kv := []interface{}{log.DefaultMessageKey, fmt.Sprintf(format, msg...)}
	DefaultLogger.Log(log.LevelWarn, append(kv, logTraceID(ctx)...))
}
