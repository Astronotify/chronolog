package internal

import "context"

const (
	TraceIDKey      = "trace_id"
	SpanIDKey       = "span_id"
	ParentSpanIDKey = "parent_span_id"

	VersionKey    = "version"
	CommitHashKey = "commit_hash"
	BuildTimeKey  = "build_time"
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

func ExtractVersion(ctx context.Context) string {
	v := ctx.Value(VersionKey)
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
