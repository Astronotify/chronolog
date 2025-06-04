package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Astronotify/chronolog"
	"github.com/Astronotify/chronolog/entries"
)

func main() {
	chronolog.Setup(chronolog.Config{Format: chronolog.FormatPretty})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		reqID := r.Header.Get("X-Request-ID")
		if reqID == "" {
			reqID = fmt.Sprintf("%d", time.Now().UnixNano())
		}

		opReq := entries.NewOperationRequestLogEntry(ctx,
			"http_request", "httpserver", reqID, r.URL.Path, r.Method)
		chronolog.Entry(ctx, opReq)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "hello")

		opRes := entries.NewOperationResponseLogEntry(opReq, http.StatusOK)
		chronolog.Entry(ctx, opRes)
	})

	chronolog.Info(context.Background(), "listening on :8080")
	http.ListenAndServe(":8080", nil)
}
