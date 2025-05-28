package entries

type ErrorLogEntry struct {
	LogEntry
	ErrorClass   string `json:"error_class"`
	ErrorMessage string `json:"error_message"`
	StackTrace   string `json:"stack_trace,omitempty"`
}
