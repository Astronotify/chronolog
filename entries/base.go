package entries

import (
	"context"
	"time"

	"github.com/Astronotify/chronolog/internal"
	Level "github.com/Astronotify/chronolog/level"
)

// LogEntry represents a structured and enriched log message emitted by the application.
//
// It captures essential metadata such as timestamp, severity level, trace identifiers, and
// additional diagnostic fields that help with observability, debugging, and auditing.
//
// Fields:
//
//   - Context: the execution context. Used internally for extracting trace and build metadata.
//     This field is not serialized to JSON.
//   - Timestamp: UTC timestamp of when the log entry was created.
//   - Level: severity level of the log (trace, info, warn, error).
//   - EventType: a string label that categorizes the type of log entry.
//     Defaults to "LogEntry", but can be overridden by embedding structs.
//   - Message: the human-readable message describing the event or situation.
//
// Trace fields:
//
//   - TraceID: unique identifier for the current trace, if available.
//   - SpanID: identifier for the current span within the trace.
//   - ParentSpanID: identifier of the parent span, if applicable.
//
// Build information:
//
//   - Version: the version of the application at build time.
//   - CommitHash: the git commit hash associated with the build.
//   - BuildTime: the timestamp of the build process.
//
// Library metadata:
//
//   - LibraryName: name of the logging library emitting the log.
//   - LibraryVersion: version of the library in use.
//   - LibraryCommit: commit hash of the library version.
//   - LibraryBuildTime: timestamp when the library was built.
//
// Extensibility:
//
//   - AdditionalData: optional user-defined key-value pairs to enrich the log.
//     Merged from the variadic map arguments passed to the constructor.
type LogEntry struct {
	Context context.Context `json:"-"`

	Timestamp time.Time      `json:"timestamp"`
	Level     Level.LogLevel `json:"level"`
	EventType string         `json:"event_type"`
	Message   string         `json:"message"`

	// Trace metadata
	TraceID      string `json:"trace_id,omitempty"`
	SpanID       string `json:"span_id,omitempty"`
	ParentSpanID string `json:"parent_span_id,omitempty"`

	// Build information
	Version    string `json:"version,omitempty"`
	CommitHash string `json:"commit_hash,omitempty"`
	BuildTime  string `json:"build_time,omitempty"`

	// Library metadata
	LibraryName      string `json:"library_name,omitempty"`
	LibraryVersion   string `json:"library_version,omitempty"`
	LibraryCommit    string `json:"library_commit,omitempty"`
	LibraryBuildTime string `json:"library_build_time,omitempty"`

	// User-defined metadata
	AdditionalData map[string]any `json:"additional_data,omitempty"`
}

// NewLogEntry creates a new LogEntry with automatic enrichment from context and build metadata.
//
// This constructor centralizes the creation of log entries and ensures consistent formatting
// across the application. It extracts trace and build information from the provided context,
// attaches standardized metadata, and optionally merges any additional structured data.
//
// Parameters:
//
//   - ctx (context.Context): the execution context used to extract trace IDs and other metadata.
//   - level (Level.LogLevel): the severity level of the log entry (e.g., LogLevelInfo, LogLevelError).
//   - message (string): a human-readable description of the log event.
//   - additionalData (...map[string]any): optional variadic key-value maps that will be merged
//     into the final `AdditionalData` field.
//
// Returns:
//
//   - LogEntry: a fully enriched structured log entry ready to be serialized or dispatched.
func NewLogEntry(ctx context.Context, level Level.LogLevel, message string, additionalData ...map[string]any) LogEntry {
	return LogEntry{
		Context:          ctx,
		Timestamp:        time.Now().UTC(),
		Level:            level,
		EventType:        "LogEntry",
		Message:          message,
		TraceID:          internal.ExtractTraceID(ctx),
		SpanID:           internal.ExtractSpanID(ctx),
		ParentSpanID:     internal.ExtractParentSpanID(ctx),
		Version:          internal.ExtractVersion(ctx),
		CommitHash:       internal.ExtractCommitHash(ctx),
		BuildTime:        internal.ExtractBuildTime(ctx),
		LibraryName:      internal.LibraryName,
		LibraryVersion:   internal.LibraryVersion,
		LibraryCommit:    internal.LibraryCommit,
		LibraryBuildTime: internal.LibraryBuildTime,
		AdditionalData:   internal.MergeAdditionalData(additionalData...),
	}
}

func (l LogEntry) GetLevel() Level.LogLevel {
	return l.Level
}
