package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mvleandro/chronolog"
	"github.com/mvleandro/chronolog/entries"
	"github.com/mvleandro/chronolog/internal"
)

func main() {

	// Set up the logger configuration
	cfg := chronolog.Config{
		Format: chronolog.FormatJSON,
		Level:  entries.LogLevelWarn, // Set the minimum log level
	}
	chronolog.Setup(cfg)

	ctx := context.Background()

	// Set up context with commit hash, build time and version
	ctx = internal.WithCommitHash(ctx, "abc1234def5678ghijkl9012mnop3456")
	ctx = internal.WithBuildTime(ctx, "2023-10-01T12:00:00Z")
	ctx = internal.WithVersion(ctx, "0.1.0")

	// Set up context with trace and span IDs
	ctx = internal.WithTraceID(ctx, "trace-12345")
	ctx = internal.WithSpanID(ctx, "span-67890")
	ctx = internal.WithParentSpanID(ctx, "parent-span-54321")

	// Simple trace log
	chronolog.Trace(ctx, "Starting profile creation process")

	// Simple debug log with additional data
	chronolog.Debug(ctx, "Validating user input", map[string]interface{}{
		"username": "johndoe"})

	// Simple info log
	chronolog.Info(ctx, "Starting profile creation process")

	// Simple warn log
	chronolog.Warn(ctx, "Profile creation took longer than expected", map[string]interface{}{
		"expected_duration_ms": 100,
		"actual_duration_ms":   150,
	})

	// Log with AdditionalData
	chronolog.Info(ctx, "Profile created successfully", map[string]interface{}{
		"profile_id": "2XjPQ7uFmf9y5nA2fXKf7VcI0dF",
		"user_id":    "8sJpQ7AbcdEf9y5nA2fXKf7Vc999",
		"location": map[string]float64{
			"latitude":  -22.9068,
			"longitude": -43.1729,
		},
	})

	// Error log with additional data
	err := fmt.Errorf("failed to persist profile to database")
	chronolog.Error(ctx, err, map[string]interface{}{
		"retries": 3,
		"region":  "us-east-1",
	})

	// --- Trace log ---
	traceBegin := entries.NewTraceBeginLogEntry(ctx, "ValidateToken")
	chronolog.Entry(ctx, traceBegin)

	time.Sleep(80 * time.Millisecond) // simulate execution

	traceEnd := entries.NewTraceEndLogEntryFromBegin(traceBegin)
	chronolog.Entry(ctx, traceEnd)

	// --- Lambda log ---
	lambdaBegin := entries.NewLambdaBeginLogEntry(ctx, "CreateUserHandler", "req-xyz-001")
	chronolog.Entry(ctx, lambdaBegin)

	time.Sleep(30 * time.Millisecond) // simulate processing

	lambdaEnd := entries.NewLambdaEndLogEntryFromBegin(lambdaBegin)
	chronolog.Entry(ctx, lambdaEnd)

	// --- Operation log ---
	opReq := entries.NewOperationRequestLogEntry(ctx, "CreateUser", "user", "req-abc-123", "/users", "POST")
	chronolog.Entry(ctx, opReq)

	time.Sleep(45 * time.Millisecond) // simulate processing

	opRes := entries.NewOperationResponseLogEntry(opReq, 201)
	chronolog.Entry(ctx, opRes)

	// --- Message lifecycle logs ---
	msgReceived := entries.NewMessageReceivedLogEntry(ctx, "msg-001", "user.created", "email-service")
	chronolog.Entry(ctx, msgReceived)

	time.Sleep(70 * time.Millisecond) // simulate processing

	msgAck := entries.NewMessageAcknowledgedLogEntryFromReceived(msgReceived)
	chronolog.Entry(ctx, msgAck)

	msgReceived2 := entries.NewMessageReceivedLogEntry(ctx, "msg-002", "user.deleted", "email-service")
	chronolog.Entry(ctx, msgReceived2)

	time.Sleep(20 * time.Millisecond) // simulate failed processing

	msgRej := entries.NewMessageRejectedLogEntryFromReceived(msgReceived2, "invalid payload")
	chronolog.Entry(ctx, msgRej)

	// --- Kubernetes context log ---
	k8sLog := entries.NewK8SLogEntry(ctx,
		"prod-cluster",
		"auth-namespace",
		"auth-pod-9f8c",
		"auth-container",
		"node-34a",
		entries.LogLevelInfo,
		"Pod started successfully",
		map[string]interface{}{
			"deployment": "auth-service",
			"version":    "v1.4.2",
		},
	)
	chronolog.Entry(ctx, k8sLog)
}
