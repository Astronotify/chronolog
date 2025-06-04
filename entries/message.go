package entries

import (
	"context"
	"time"

	Level "github.com/Astronotify/chronolog/level"
)

// MessageReceivedLogEntry represents a log entry indicating that a message has been received
// by a consumer from a given topic.
//
// This entry captures key metadata that identifies the message, the topic it was published to,
// and the consumer that received it. It serves as the starting point for tracking message lifecycle.
//
// Fields:
//
//   - MessageID: the unique identifier of the message.
//   - Topic: the topic or queue where the message was published.
//   - Consumer: the name or identifier of the consumer that received the message.
type MessageReceivedLogEntry struct {
	LogEntry
	MessageID string `json:"message_id"`
	Topic     string `json:"topic"`
	Consumer  string `json:"consumer"`
}

// MessageAcknowledgedLogEntry represents a log entry indicating successful processing of a message.
//
// This entry includes the same identifying fields as the received log, along with the duration
// between reception and acknowledgment.
//
// Fields:
//
//   - MessageID: the unique identifier of the message.
//   - Topic: the topic or queue where the message was published.
//   - Consumer: the name or identifier of the consumer that processed the message.
//   - DurationMs: the time taken to process the message, in milliseconds.
type MessageAcknowledgedLogEntry struct {
	LogEntry
	MessageID  string `json:"message_id"`
	Topic      string `json:"topic"`
	Consumer   string `json:"consumer"`
	DurationMs int64  `json:"duration_ms"`
}

// MessageRejectedLogEntry represents a log entry indicating that a message was rejected during processing.
//
// This entry includes metadata about the message and consumer, along with the rejection reason
// and duration of processing.
//
// Fields:
//
//   - MessageID: the unique identifier of the message.
//   - Topic: the topic or queue where the message was published.
//   - Consumer: the name or identifier of the consumer that rejected the message.
//   - Reason: the reason for rejecting the message.
//   - DurationMs: the time spent processing the message before rejection, in milliseconds.
type MessageRejectedLogEntry struct {
	LogEntry
	MessageID  string `json:"message_id"`
	Topic      string `json:"topic"`
	Consumer   string `json:"consumer"`
	Reason     string `json:"reason"`
	DurationMs int64  `json:"duration_ms"`
}

// NewMessageReceivedLogEntry creates a log entry marking the moment a message is received by a consumer.
//
// Parameters:
//
//   - ctx (context.Context): context used to enrich the log with trace/build info.
//   - messageID (string): the unique ID of the received message.
//   - topic (string): the topic or queue from which the message was received.
//   - consumer (string): the name of the receiving consumer.
//   - additionalData (...map[string]any): optional metadata to enrich the log.
//
// Returns:
//
//   - MessageReceivedLogEntry: a structured log entry for message receipt.
func NewMessageReceivedLogEntry(
	ctx context.Context,
	messageID, topic, consumer string,
	additionalData ...map[string]any,
) MessageReceivedLogEntry {
	entry := MessageReceivedLogEntry{
		LogEntry:  NewLogEntry(ctx, Level.Info, "Message received", additionalData...),
		MessageID: messageID,
		Topic:     topic,
		Consumer:  consumer,
	}
	entry.EventType = "MessageReceivedLogEntry"
	return entry
}

// NewMessageAcknowledgedLogEntryFromReceived creates a log entry indicating successful processing
// of a previously received message, automatically calculating the duration.
//
// Parameters:
//
//   - received (MessageReceivedLogEntry): the original message receipt log entry.
//   - additionalData (...map[string]any): optional metadata to enrich the log.
//
// Returns:
//
//   - MessageAcknowledgedLogEntry: a structured log with timing information.
func NewMessageAcknowledgedLogEntryFromReceived(
	received MessageReceivedLogEntry,
	additionalData ...map[string]any,
) MessageAcknowledgedLogEntry {
	duration := time.Since(received.Timestamp).Milliseconds()

	entry := MessageAcknowledgedLogEntry{
		LogEntry:   NewLogEntry(received.Context, Level.Info, "Message acknowledged", additionalData...),
		MessageID:  received.MessageID,
		Topic:      received.Topic,
		Consumer:   received.Consumer,
		DurationMs: duration,
	}
	entry.EventType = "MessageAcknowledgedLogEntry"
	return entry
}

// NewMessageRejectedLogEntryFromReceived creates a log entry indicating failed processing of a message,
// capturing the rejection reason and processing duration.
//
// Parameters:
//
//   - received (MessageReceivedLogEntry): the original message receipt log entry.
//   - reason (string): the reason for rejecting the message.
//   - additionalData (...map[string]any): optional metadata to enrich the log.
//
// Returns:
//
//   - MessageRejectedLogEntry: a structured log with rejection details and timing.
func NewMessageRejectedLogEntryFromReceived(
	received MessageReceivedLogEntry,
	reason string,
	additionalData ...map[string]any,
) MessageRejectedLogEntry {
	duration := time.Since(received.Timestamp).Milliseconds()

	entry := MessageRejectedLogEntry{
		LogEntry:   NewLogEntry(received.Context, Level.Error, "Message rejected", additionalData...),
		MessageID:  received.MessageID,
		Topic:      received.Topic,
		Consumer:   received.Consumer,
		Reason:     reason,
		DurationMs: duration,
	}
	entry.EventType = "MessageRejectedLogEntry"
	return entry
}
