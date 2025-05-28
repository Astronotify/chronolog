package entries

type OperationRequestLogEntry struct {
	LogEntry
	OperationName string `json:"operation_name"`
	Resource      string `json:"resource"`
	RequestID     string `json:"request_id"`
	Path          string `json:"path"`
	HTTPMethod    string `json:"http_method"`
}

type OperationResponseLogEntry struct {
	LogEntry
	OperationName string `json:"operation_name"`
	Resource      string `json:"resource"`
	RequestID     string `json:"request_id"`
	HTTPStatus    int    `json:"http_status"`
	DurationMs    int64  `json:"duration_ms"`
}
