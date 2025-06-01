package internal

import "context"

type contextKey string

const (
	TraceIDKey      contextKey = "trace_id"
	SpanIDKey       contextKey = "span_id"
	ParentSpanIDKey contextKey = "parent_span_id"
	CommitHashKey   contextKey = "commit_hash"
	BuildTimeKey    contextKey = "build_time"
	VersionKey      contextKey = "version"
)

func ExtractTraceID(ctx context.Context) string {
	v := ctx.Value(TraceIDKey)
	if v != nil {
		return v.(string)
	}
	return ""
}

func ExtractSpanID(ctx context.Context) string {
	v := ctx.Value(SpanIDKey)
	if v != nil {
		return v.(string)
	}
	return ""
}

func ExtractParentSpanID(ctx context.Context) string {
	v := ctx.Value(ParentSpanIDKey)
	if v != nil {
		return v.(string)
	}
	return ""
}

func ExtractCommitHash(ctx context.Context) string {
	v := ctx.Value(CommitHashKey)
	if v != nil {
		return v.(string)
	}
	return ""
}

func ExtractBuildTime(ctx context.Context) string {
	v := ctx.Value(BuildTimeKey)
	if v != nil {
		return v.(string)
	}
	return ""
}

func ExtractVersion(ctx context.Context) string {
	v := ctx.Value(VersionKey)
	if v != nil {
		return v.(string)
	}
	return ""
}

func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, TraceIDKey, traceID)
}

func WithSpanID(ctx context.Context, spanID string) context.Context {
	return context.WithValue(ctx, SpanIDKey, spanID)
}

func WithParentSpanID(ctx context.Context, parentSpanID string) context.Context {
	return context.WithValue(ctx, ParentSpanIDKey, parentSpanID)
}

func WithCommitHash(ctx context.Context, commitHash string) context.Context {
	return context.WithValue(ctx, CommitHashKey, commitHash)
}

func WithBuildTime(ctx context.Context, buildTime string) context.Context {
	return context.WithValue(ctx, BuildTimeKey, buildTime)
}

func WithVersion(ctx context.Context, version string) context.Context {
	return context.WithValue(ctx, VersionKey, version)
}
