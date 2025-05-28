package entries

import (
	"time"

	"github.com/mvleandro/chronolog/internal"
)

type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Level     string    `json:"level"`
	EventType string    `json:"event_type"`
	Message   string    `json:"message"`

	// Trace
	TraceID      string `json:"trace_id,omitempty"`
	SpanID       string `json:"span_id,omitempty"`
	ParentSpanID string `json:"parent_span_id,omitempty"`

	// Build Info
	Version    string `json:"version,omitempty"`
	CommitHash string `json:"commit_hash,omitempty"`
	BuildTime  string `json:"build_time,omitempty"`

	// Extensible
	AdditionalData map[string]any `json:"additional_data,omitempty"`
}

func (e *LogEntry) Init() {
	if e.Timestamp.IsZero() {
		e.Timestamp = time.Now().UTC()
	}
	if e.Level == "" {
		e.Level = "info"
	}
	if e.EventType == "" {
		e.EventType = internal.GetStructName(e)
	}
	if e.AdditionalData == nil {
		e.AdditionalData = make(map[string]any)
	}
}
