package internal

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
)

// JSONOnlyHandler is a slog.Handler that prints only the serialized JSON of the "event" field.
type JSONOnlyHandler struct {
	writer io.Writer
}

// NewJSONOnlyHandler creates a new JSONOnlyHandler.
func NewJSONOnlyHandler(w io.Writer) *JSONOnlyHandler {
	return &JSONOnlyHandler{
		writer: w,
	}
}

func (h *JSONOnlyHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return true
}

func (h *JSONOnlyHandler) Handle(_ context.Context, record slog.Record) error {
	var event any

	record.Attrs(func(attr slog.Attr) bool {
		if attr.Key == "event" {
			event = attr.Value.Any()
			return false // we found it, stop iteration
		}
		return true
	})

	if event == nil {
		return nil // nothing to log
	}

	enc := json.NewEncoder(h.writer)
	return enc.Encode(event)
}

func (h *JSONOnlyHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	// no-op, stateless handler
	return h
}

func (h *JSONOnlyHandler) WithGroup(_ string) slog.Handler {
	// no-op, stateless handler
	return h
}
