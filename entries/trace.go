package entries

type TraceBeginLogEntry struct {
	LogEntry
	Name string `json:"name"`
}

type TraceEndLogEntry struct {
	LogEntry
	Name       string `json:"name"`
	DurationMs int64  `json:"duration_ms"`
}
