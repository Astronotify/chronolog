package ctx

import (
	"context"

	"github.com/Astronotify/chronolog/internal"
)

// WithTraceID stores the given trace ID in the context.
func WithTraceID(ctx context.Context, traceID string) context.Context {
	return internal.WithTraceID(ctx, traceID)
}

// WithSpanID stores the given span ID in the context.
func WithSpanID(ctx context.Context, spanID string) context.Context {
	return internal.WithSpanID(ctx, spanID)
}

// WithParentSpanID stores the parent span ID in the context.
func WithParentSpanID(ctx context.Context, parentSpanID string) context.Context {
	return internal.WithParentSpanID(ctx, parentSpanID)
}

// WithCommitHash stores the commit hash in the context.
func WithCommitHash(ctx context.Context, commitHash string) context.Context {
	return internal.WithCommitHash(ctx, commitHash)
}

// WithBuildTime stores the build time in the context.
func WithBuildTime(ctx context.Context, buildTime string) context.Context {
	return internal.WithBuildTime(ctx, buildTime)
}

// WithVersion stores the application version in the context.
func WithVersion(ctx context.Context, version string) context.Context {
	return internal.WithVersion(ctx, version)
}
