package chronolog

import (
	"context"
	"errors"
	"log/slog"
	"reflect"
	"testing"

	Level "github.com/Astronotify/chronolog/level"
)

type capturingHandler struct {
	levels []slog.Level
}

func (h *capturingHandler) Enabled(_ context.Context, _ slog.Level) bool { return true }
func (h *capturingHandler) Handle(_ context.Context, r slog.Record) error {
	h.levels = append(h.levels, r.Level)
	return nil
}
func (h *capturingHandler) WithAttrs(_ []slog.Attr) slog.Handler { return h }
func (h *capturingHandler) WithGroup(string) slog.Handler        { return h }

func TestWriteLevelMapping(t *testing.T) {
	handler := &capturingHandler{}
	logger = slog.New(handler)
	minimumLogLevel = Level.Trace

	ctx := context.Background()
	Trace(ctx, "trace")
	Debug(ctx, "debug")
	Info(ctx, "info")
	Warn(ctx, "warn")
	Error(ctx, errors.New("err"))

	want := []slog.Level{
		slog.LevelDebug,
		slog.LevelDebug,
		slog.LevelInfo,
		slog.LevelWarn,
		slog.LevelError,
	}

	if !reflect.DeepEqual(handler.levels, want) {
		t.Errorf("levels mismatch: got %v want %v", handler.levels, want)
	}
}
