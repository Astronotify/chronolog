package entries

import (
	"context"

	"github.com/mvleandro/chronolog/internal"
)

// ErrorLogEntry represents a structured log entry dedicated to error events.
//
// It extends LogEntry by adding metadata specific to runtime errors, enabling better diagnostics,
// filtering, and categorization of failures during execution.
//
// Fields:
//
//   - ErrorClass: a short identifier that classifies the type of error (e.g., "ValidationError", "IOError").
//     Derived via internal heuristics or conventions.
//   - ErrorMessage: the error message returned by the error object. Also used as the main log message.
//   - StackTrace: an optional string representation of the stack trace where the error occurred.
//     Included only if available and supported by the error type or logger configuration.
type ErrorLogEntry struct {
	LogEntry
	ErrorClass   string `json:"error_class"`
	ErrorMessage string `json:"error_message"`
	StackTrace   string `json:"stack_trace,omitempty"`
}

// NewErrorLogEntry creates a new ErrorLogEntry from a given error and optional metadata.
//
// This constructor extracts trace and build context, classifies the error, and optionally includes
// a stack trace if available. It also merges any additional data provided for enrichment.
//
// Parameters:
//
//   - ctx (context.Context): the execution context from which trace and build metadata are extracted.
//   - err (error): the error instance to log. Must not be nil.
//   - additionalData (...map[string]any): optional variadic key-value maps to include in AdditionalData.
//
// Returns:
//
//   - ErrorLogEntry: a structured and enriched log entry representing the error.
func NewErrorLogEntry(ctx context.Context, err error, additionalData ...map[string]any) ErrorLogEntry {
	entry := ErrorLogEntry{
		LogEntry: NewLogEntry(
			ctx,
			LogLevelError,
			err.Error(),
			internal.MergeAdditionalData(additionalData...),
		),
		ErrorClass:   internal.ClassifyError(err),
		ErrorMessage: err.Error(),
		StackTrace:   internal.ExtractStackTrace(err),
	}
	entry.EventType = "ErrorLogEntry"
	return entry
}
