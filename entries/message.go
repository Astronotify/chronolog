package entries

type MessageReceivedLogEntry struct {
	LogEntry
	MessageID string `json:"message_id"`
	Topic     string `json:"topic"`
	Consumer  string `json:"consumer"`
}

type MessageAcknowledgedLogEntry struct {
	LogEntry
	MessageID string `json:"message_id"`
	Topic     string `json:"topic"`
	Consumer  string `json:"consumer"`
}

type MessageRejectedLogEntry struct {
	LogEntry
	MessageID string `json:"message_id"`
	Topic     string `json:"topic"`
	Consumer  string `json:"consumer"`
	Reason    string `json:"reason"`
}
