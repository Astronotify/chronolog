package entries

type LambdaBeginLogEntry struct {
	LogEntry
	FunctionName string `json:"function_name"`
	RequestID    string `json:"request_id"`
}

type LambdaEndLogEntry struct {
	LogEntry
	FunctionName string `json:"function_name"`
	RequestID    string `json:"request_id"`
	DurationMs   int64  `json:"duration_ms"`
}
