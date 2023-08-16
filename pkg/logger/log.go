package logger

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/trace"
)

var logTraceID = func(ctx context.Context) []interface{} {
	if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
		return []interface{}{"traceID", span.TraceID().String()}
	}
	return nil
}

func Info(ctx context.Context, msg ...interface{}) {
	kv := []interface{}{log.DefaultMessageKey, fmt.Sprint(msg...)}
	log.Log(log.LevelInfo, append(kv, logTraceID(ctx)...))
}

func Infof(ctx context.Context, format string, msg ...interface{}) {
	kv := []interface{}{log.DefaultMessageKey, fmt.Sprintf(format, msg...)}
	log.Log(log.LevelInfo, append(kv, logTraceID(ctx)...))
}

func Error(ctx context.Context, msg ...interface{}) {
	kv := []interface{}{log.DefaultMessageKey, fmt.Sprint(msg...)}
	log.Log(log.LevelError, append(kv, logTraceID(ctx)...))
}

func Errorf(ctx context.Context, format string, msg ...interface{}) {
	kv := []interface{}{log.DefaultMessageKey, fmt.Sprintf(format, msg...)}
	log.Log(log.LevelError, append(kv, logTraceID(ctx)...))
}

func Fatalf(ctx context.Context, format string, msg ...interface{}) {
	kv := []interface{}{log.DefaultMessageKey, fmt.Sprintf(format, msg...)}
	log.Log(log.LevelFatal, append(kv, logTraceID(ctx)...))
}
