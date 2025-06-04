package ctx_test

import (
	"context"
	"testing"

	chronologctx "github.com/Astronotify/chronolog/ctx"
	"github.com/Astronotify/chronolog/internal"
)

func TestContextSetters(t *testing.T) {
	tests := []struct {
		name      string
		setter    func(context.Context) context.Context
		extractor func(context.Context) string
		want      string
	}{
		{
			name:      "traceID",
			setter:    func(c context.Context) context.Context { return chronologctx.WithTraceID(c, "trace") },
			extractor: internal.ExtractTraceID,
			want:      "trace",
		},
		{
			name:      "spanID",
			setter:    func(c context.Context) context.Context { return chronologctx.WithSpanID(c, "span") },
			extractor: internal.ExtractSpanID,
			want:      "span",
		},
		{
			name:      "parentSpanID",
			setter:    func(c context.Context) context.Context { return chronologctx.WithParentSpanID(c, "parent") },
			extractor: internal.ExtractParentSpanID,
			want:      "parent",
		},
		{
			name:      "commitHash",
			setter:    func(c context.Context) context.Context { return chronologctx.WithCommitHash(c, "commit") },
			extractor: internal.ExtractCommitHash,
			want:      "commit",
		},
		{
			name:      "buildTime",
			setter:    func(c context.Context) context.Context { return chronologctx.WithBuildTime(c, "time") },
			extractor: internal.ExtractBuildTime,
			want:      "time",
		},
		{
			name:      "version",
			setter:    func(c context.Context) context.Context { return chronologctx.WithVersion(c, "v1") },
			extractor: internal.ExtractVersion,
			want:      "v1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := tt.setter(context.Background())
			got := tt.extractor(ctx)
			if got != tt.want {
				t.Errorf("%s: expected %q, got %q", tt.name, tt.want, got)
			}
		})
	}
}
