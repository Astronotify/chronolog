package main

import (
	"context"
	"time"

	"github.com/Astronotify/chronolog"
	"github.com/Astronotify/chronolog/entries"
)

func main() {
	chronolog.Setup(chronolog.Config{Format: chronolog.FormatPretty})
	ctx := context.Background()

	received := entries.NewMessageReceivedLogEntry(ctx,
		"msg-1", "user.signup", "email-service")
	chronolog.Entry(ctx, received)

	time.Sleep(10 * time.Millisecond)
	ack := entries.NewMessageAcknowledgedLogEntryFromReceived(received)
	chronolog.Entry(ctx, ack)

	received2 := entries.NewMessageReceivedLogEntry(ctx,
		"msg-2", "user.delete", "email-service")
	chronolog.Entry(ctx, received2)

	time.Sleep(5 * time.Millisecond)
	rej := entries.NewMessageRejectedLogEntryFromReceived(received2,
		"invalid data")
	chronolog.Entry(ctx, rej)
}
