package entries

import (
	"context"
	"time"

	Level "github.com/Astronotify/chronolog/level"
)

// TraceBeginLogEntry represents the beginning of a trace section in the application flow.
//
// This log entry is typically used to mark the start of a measurable unit of work,
// allowing correlation with a matching TraceEndLogEntry to calculate total duration.
//
// Fields:
//
//   - Name: a human-readable label identifying the trace section (e.g., "LoadUserProfile").
type TraceBeginLogEntry struct {
	LogEntry
	Name string `json:"name"`
}

// TraceEndLogEntry represents the end of a trace section previously marked by a TraceBeginLogEntry.
//
// It includes the total time spent executing the traced block, measured in milliseconds.
//
// Fields:
//
//   - Name: the same name provided in the TraceBeginLogEntry.
//   - DurationMs: the elapsed time between the TraceBeginLogEntry and this entry.
type TraceEndLogEntry struct {
	LogEntry
	Name       string `json:"name"`
	DurationMs int64  `json:"duration_ms"`
}

// NewTraceBeginLogEntry creates a new trace entry to mark the start of a measurable block of execution.
//
// This entry records the timestamp and contextual metadata, allowing future correlation
// with a corresponding TraceEndLogEntry to compute total execution time.
//
// Parameters:
//
//   - ctx (context.Context): the execution context used for trace and metadata enrichment.
//   - name (string): a label identifying the trace section.
//   - additionalData (...map[string]any): optional structured metadata for enrichment.
//
// Returns:
//
//   - TraceBeginLogEntry: a structured entry marking the start of a traceable operation.
func NewTraceBeginLogEntry(
	ctx context.Context,
	name string,
	additionalData ...map[string]any,
) TraceBeginLogEntry {
	entry := TraceBeginLogEntry{
		LogEntry: NewLogEntry(
			ctx,
			Level.Trace,
			"Trace started",
			additionalData...,
		),
		Name: name,
	}
	entry.EventType = "TraceBeginLogEntry"
	return entry
}

// NewTraceEndLogEntryFromBegin creates a matching end log entry for a given TraceBeginLogEntry,
// calculating the total duration of the traced operation in milliseconds.
//
// This function is used to close out a traced block and record how long it took to execute.
//
// Parameters:
//
//   - begin (TraceBeginLogEntry): the previously generated trace start entry.
//   - additionalData (...map[string]any): optional additional metadata.
//
// Returns:
//
//   - TraceEndLogEntry: a structured entry marking the end of a traceable operation.
func NewTraceEndLogEntryFromBegin(
	begin TraceBeginLogEntry,
	additionalData ...map[string]any,
) TraceEndLogEntry {
	duration := time.Since(begin.Timestamp).Milliseconds()

	entry := TraceEndLogEntry{
		LogEntry: NewLogEntry(
			begin.Context,
			Level.Trace,
			"Trace ended",
			additionalData...,
		),
		Name:       begin.Name,
		DurationMs: duration,
	}
	entry.EventType = "TraceEndLogEntry"
	return entry
}
