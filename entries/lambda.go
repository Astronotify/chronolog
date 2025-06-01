package entries

import (
	"context"
	"time"
)

// LambdaBeginLogEntry represents the initial log entry for an AWS Lambda function invocation.
//
// This entry captures the moment the function begins executing and includes identifying
// information such as the function name and AWS request ID.
//
// Fields:
//
//   - FunctionName: the name of the AWS Lambda function being executed.
//   - RequestID: the unique identifier for the current Lambda invocation (provided by AWS).
type LambdaBeginLogEntry struct {
	LogEntry
	FunctionName string `json:"function_name"`
	RequestID    string `json:"request_id"`
}

// LambdaEndLogEntry represents the final log entry for an AWS Lambda function invocation.
//
// It is used to capture the end of execution along with the duration of the invocation,
// calculated automatically based on the timestamp of the corresponding LambdaBeginLogEntry.
//
// Fields:
//
//   - FunctionName: the name of the AWS Lambda function being executed.
//   - RequestID: the unique identifier for the current Lambda invocation.
//   - DurationMs: the execution time in milliseconds, computed as the difference between
//     the start and end timestamps.
type LambdaEndLogEntry struct {
	LogEntry
	FunctionName string `json:"function_name"`
	RequestID    string `json:"request_id"`
	DurationMs   int64  `json:"duration_ms"`
}

// NewLambdaBeginLogEntry creates a new LambdaBeginLogEntry with execution context and metadata.
//
// This constructor is intended to be used at the start of a Lambda function's execution to
// capture the timestamp, trace information, and identifying metadata.
//
// Parameters:
//
//   - ctx (context.Context): the execution context for metadata extraction.
//   - functionName (string): the name of the Lambda function.
//   - requestID (string): the AWS request ID associated with the invocation.
//   - additionalData (...map[string]any): optional metadata to enrich the log entry.
//
// Returns:
//
//   - LambdaBeginLogEntry: a fully structured log entry marking the start of the function.
func NewLambdaBeginLogEntry(
	ctx context.Context,
	functionName, requestID string,
	additionalData ...map[string]any,
) LambdaBeginLogEntry {
	entry := LambdaBeginLogEntry{
		LogEntry:     NewLogEntry(ctx, Info, "Lambda function started", additionalData...),
		FunctionName: functionName,
		RequestID:    requestID,
	}
	entry.EventType = "LambdaBeginLogEntry"
	return entry
}

// NewLambdaEndLogEntryFromBegin creates a LambdaEndLogEntry based on a corresponding begin entry,
// automatically calculating the duration of the invocation.
//
// This constructor should be called at the end of the Lambda execution. It uses the timestamp
// from the LambdaBeginLogEntry to calculate the duration in milliseconds.
//
// Parameters:
//
//   - begin (LambdaBeginLogEntry): the previously created begin log entry.
//   - additionalData (...map[string]any): optional metadata to include in the final log.
//
// Returns:
//
//   - LambdaEndLogEntry: a structured log entry representing the end of execution with duration.
func NewLambdaEndLogEntryFromBegin(
	begin LambdaBeginLogEntry,
	additionalData ...map[string]any,
) LambdaEndLogEntry {
	duration := time.Since(begin.Timestamp).Milliseconds()

	entry := LambdaEndLogEntry{
		LogEntry:     NewLogEntry(begin.Context, Info, "Lambda function finished", additionalData...),
		FunctionName: begin.FunctionName,
		RequestID:    begin.RequestID,
		DurationMs:   duration,
	}
	entry.EventType = "LambdaEndLogEntry"
	return entry
}
