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

	begin := entries.NewLambdaBeginLogEntry(ctx,
		"HelloLambda", "req-123")
	chronolog.Entry(ctx, begin)

	time.Sleep(20 * time.Millisecond)

	end := entries.NewLambdaEndLogEntryFromBegin(begin)
	chronolog.Entry(ctx, end)
}
