package entries

import (
	"context"
	"time"

	Level "github.com/Astronotify/chronolog/level"
)

// OperationRequestLogEntry represents the log entry for the start of a logical application operation,
// such as a RESTful request or command execution.
//
// It contains identifying information about the operation being performed, including
// its name, associated resource, and request metadata.
//
// Fields:
//
//   - OperationName: the logical name of the operation being performed (e.g., "CreateUser").
//   - Resource: the type of entity or subsystem being acted upon (e.g., "user", "order").
//   - RequestID: a unique identifier for the operation invocation, used for correlation.
//   - Path: the HTTP path or route associated with the operation.
//   - HTTPMethod: the HTTP method (GET, POST, etc.) used in the request.
type OperationRequestLogEntry struct {
	LogEntry

	OperationName string `json:"operation_name"`
	Resource      string `json:"resource"`
	RequestID     string `json:"request_id"`
	Path          string `json:"path"`
	HTTPMethod    string `json:"http_method"`
}

// OperationResponseLogEntry represents the log entry for the end of an application operation,
// capturing its outcome and the time taken to execute.
//
// Fields:
//
//   - OperationName: the same logical name used in the request log.
//   - Resource: the same resource context used in the request log.
//   - RequestID: a unique identifier matching the original request.
//   - HTTPStatus: the HTTP status code returned as a result of the operation.
//   - DurationMs: the time elapsed from the request to the response, in milliseconds.
type OperationResponseLogEntry struct {
	LogEntry
	OperationName string `json:"operation_name"`
	Resource      string `json:"resource"`
	RequestID     string `json:"request_id"`
	HTTPStatus    int    `json:"http_status"`
	DurationMs    int64  `json:"duration_ms"`
}

// NewOperationRequestLogEntry creates a structured log entry for the start of an operation.
//
// This constructor captures identifying metadata and automatically enriches the log
// with trace, version, and build information from the context.
//
// Parameters:
//
//   - ctx (context.Context): the context used for trace and metadata enrichment.
//   - operationName (string): a descriptive name for the logical operation.
//   - resource (string): the domain or entity the operation applies to.
//   - requestID (string): a unique ID used to correlate request/response pairs.
//   - path (string): the HTTP route or endpoint involved.
//   - httpMethod (string): the HTTP verb used in the request.
//   - additionalData (...map[string]any): optional user-defined metadata to enrich the log.
//
// Returns:
//
//   - OperationRequestLogEntry: a structured entry representing the start of the operation.
func NewOperationRequestLogEntry(
	ctx context.Context,
	operationName, resource, requestID, path, httpMethod string,
	additionalData ...map[string]any,
) OperationRequestLogEntry {
	entry := OperationRequestLogEntry{
		LogEntry:      NewLogEntry(ctx, Level.Info, "Operation request received", additionalData...),
		OperationName: operationName,
		Resource:      resource,
		RequestID:     requestID,
		Path:          path,
		HTTPMethod:    httpMethod,
	}
	entry.EventType = "OperationRequestLogEntry"
	return entry
}

// NewOperationResponseLogEntry creates a structured log entry representing the completion of an operation.
//
// It calculates the duration of the operation by comparing the current time with the timestamp
// in the original request log. Additional metadata can be attached if needed.
//
// Parameters:
//
//   - req (OperationRequestLogEntry): the original operation request entry.
//   - httpStatus (int): the resulting HTTP status code of the operation.
//   - additionalData (...map[string]any): optional user-defined metadata.
//
// Returns:
//
//   - OperationResponseLogEntry: a log entry that includes timing and result details.
func NewOperationResponseLogEntry(
	req OperationRequestLogEntry,
	httpStatus int,
	additionalData ...map[string]any,
) OperationResponseLogEntry {
	duration := time.Since(req.Timestamp).Milliseconds()

	entry := OperationResponseLogEntry{
		LogEntry:      NewLogEntry(req.Context, Level.Info, "Operation response sent", additionalData...),
		OperationName: req.OperationName,
		Resource:      req.Resource,
		RequestID:     req.RequestID,
		HTTPStatus:    httpStatus,
		DurationMs:    duration,
	}
	entry.EventType = "OperationResponseLogEntry"
	return entry
}
