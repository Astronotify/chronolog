package chronolog

import (
	"context"
	"log/slog"
	"time"

	"github.com/mvleandro/chronolog/entries"
	"github.com/mvleandro/chronolog/internal"
)

var logger *slog.Logger = slog.Default()

func Configure(handler slog.Handler) {
	logger = slog.New(handler)
}

func Info(ctx context.Context, message string, additionalData ...map[string]any) {
	entry := entries.LogEntry{
		Timestamp:      time.Now().UTC(),
		Level:          "info",
		EventType:      "LogEntry",
		Message:        message,
		TraceID:        internal.ExtractTraceID(ctx),
		SpanID:         internal.ExtractSpanID(ctx),
		ParentSpanID:   internal.ExtractParentSpanID(ctx),
		Version:        internal.ExtractVersion(ctx),
		CommitHash:     internal.ExtractCommitHash(ctx),
		BuildTime:      internal.ExtractBuildTime(ctx),
		AdditionalData: internal.MergeAdditionalData(additionalData...),
	}

	write(entry)
}

func Warn(ctx context.Context, message string, additionalData ...map[string]any) {
	entry := entries.LogEntry{
		Timestamp:      time.Now().UTC(),
		Level:          "warn",
		EventType:      "LogEntry",
		TraceID:        internal.ExtractTraceID(ctx),
		SpanID:         internal.ExtractSpanID(ctx),
		ParentSpanID:   internal.ExtractParentSpanID(ctx),
		Version:        internal.ExtractVersion(ctx),
		CommitHash:     internal.ExtractCommitHash(ctx),
		BuildTime:      internal.ExtractBuildTime(ctx),
		AdditionalData: internal.MergeAdditionalData(additionalData...),
	}

	write(entry)
}

func Error(ctx context.Context, err error, additionalData ...map[string]any) {
	entry := entries.ErrorLogEntry{
		LogEntry: entries.LogEntry{
			Timestamp:      time.Now().UTC(),
			Level:          "error",
			EventType:      "ErrorLogEntry",
			Message:        err.Error(),
			TraceID:        internal.ExtractTraceID(ctx),
			SpanID:         internal.ExtractSpanID(ctx),
			ParentSpanID:   internal.ExtractParentSpanID(ctx),
			Version:        internal.ExtractVersion(ctx),
			CommitHash:     internal.ExtractCommitHash(ctx),
			BuildTime:      internal.ExtractBuildTime(ctx),
			AdditionalData: internal.MergeAdditionalData(additionalData...),
		},
		ErrorClass:   internal.ClassifyError(err),
		ErrorMessage: err.Error(),
		StackTrace:   internal.ExtractStackTrace(err),
	}

	write(entry)
}

func Message(ctx context.Context, entry any) {
	write(entry)
}

func write(entry any) {
	logger.Log(context.Background(), slog.LevelInfo, "log", slog.Any("event", entry))
}
