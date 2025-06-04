package chronolog

import (
	"context"
	"log/slog"
	"os"

	"github.com/Astronotify/chronolog/entries"
	"github.com/Astronotify/chronolog/internal"
	Level "github.com/Astronotify/chronolog/level"
)

var logger *slog.Logger
var minimumLogLevel Level.LogLevel = Level.Info

func Setup(cfg Config) {
	cfg.applyDefaults()
	minimumLogLevel = cfg.MinimumLogLevel

	var handler slog.Handler
	switch cfg.Format {
	case FormatPretty:
		handler = internal.NewPrettyConsoleHandler(cfg.Writer)
	case FormatJSON:
		handler = internal.NewJSONOnlyHandler(cfg.Writer)
	default:
		handler = internal.NewJSONOnlyHandler(cfg.Writer)
	}

	logger = slog.New(handler)
}

// Trace logs a detailed message for low-level debugging purposes.
//
// This function is intended for developers to trace execution flow and inspect
// internal state during development or troubleshooting. It should not be used in production
// unless necessary due to its verbose nature.
//
// Parameters:
//   - ctx (context.Context): The context associated with the log entry, used for metadata like trace IDs.
//   - message (string): The log message to be recorded.
//   - additionalData (...map[string]any): Variadic slice of maps containing additional key-value data
//     to include in the structured log entry. These maps are merged internally.
//
// Returns:
//   - None. This function produces side effects by emitting a log entry through the logger pipeline.
func Trace(ctx context.Context, message string, additionalData ...map[string]any) {
	entry := entries.NewLogEntry(
		ctx,
		Level.Trace,
		message,
		internal.MergeAdditionalData(additionalData...),
	)

	write(ctx, entry)
}

// Debug logs a message for detailed debugging information.
//
// This function is intended for developers to log detailed information that is more verbose
// than info logs, useful for diagnosing issues during development or troubleshooting.
//
// Parameters:
//   - ctx (context.Context): The context associated with the log entry, used for metadata like trace IDs.
//   - message (string): The log message to be recorded.
//   - additionalData (...map[string]any): Variadic slice of maps containing additional key-value data
//     to include in the structured log entry. These maps are merged internally.
//
// Returns:
//   - None. This function produces side effects by emitting a log entry through the logger pipeline.
func Debug(ctx context.Context, message string, additionalData ...map[string]any) {
	entry := entries.NewLogEntry(
		ctx,
		Level.Debug,
		message,
		internal.MergeAdditionalData(additionalData...),
	)

	write(ctx, entry)
}

// Info logs an informational message, optionally enriched with additional contextual data.
//
// This function should be used to log general application events that are useful for understanding
// normal system behavior (e.g., startup, configuration loaded, user actions).
//
// Parameters:
//   - ctx (context.Context): The context associated with the log entry. Used to extract metadata
//     such as trace IDs or user/session information.
//   - message (string): The log message to be recorded.
//   - additionalData (...map[string]any): Variadic slice of maps containing additional key-value data
//     to include in the structured log entry. These maps are merged internally.
//
// Returns:
//   - None. This function produces side effects by emitting a log entry through the logger pipeline.
func Info(ctx context.Context, message string, additionalData ...map[string]any) {
	entry := entries.NewLogEntry(
		ctx,
		Level.Info,
		message,
		internal.MergeAdditionalData(additionalData...),
	)

	write(ctx, entry)
}

// Warn logs a warning message indicating a potential issue that is not necessarily an error.
//
// Use this function to highlight situations that may need investigation but are not critical,
// such as configuration overrides, slow responses, or unexpected inputs.
//
// Parameters:
//   - ctx (context.Context): The execution context for the log entry.
//   - message (string): A descriptive warning message.
//   - additionalData (...map[string]any): Optional structured data to attach to the log for debugging.
//
// Returns:
//   - None. The log entry is processed and forwarded to the underlying logging system.
func Warn(ctx context.Context, message string, additionalData ...map[string]any) {
	entry := entries.NewLogEntry(
		ctx,
		Level.Warn,
		message,
		internal.MergeAdditionalData(additionalData...),
	)

	write(ctx, entry)
}

// Error logs a structured error message, along with optional diagnostic data.
//
// This function is intended for actual runtime errors, exceptions, or unexpected failures.
// It transforms the error into a structured log entry using `entries.NewErrorLogEntry`.
//
// Parameters:
//   - ctx (context.Context): The context for propagating metadata like correlation IDs.
//   - err (error): The error instance to log. Expected to be non-nil.
//   - additionalData (...map[string]any): Optional structured metadata to assist in debugging.
//
// Returns:
//   - None. The error is emitted as a structured log entry.
func Error(ctx context.Context, err error, additionalData ...map[string]any) {
	entry := entries.NewErrorLogEntry(
		ctx,
		err,
		internal.MergeAdditionalData(additionalData...),
	)

	write(ctx, entry)
}

// Entry logs a fully preconstructed log entry.
//
// This function is useful when you already have a custom or advanced entry object
// that conforms to your internal logging schema or log serialization format.
//
// Parameters:
//   - ctx (context.Context): The execution context.
//   - entry (any): A prebuilt structured log entry object. May be a map, struct, or custom type,
//     depending on your logger’s capabilities.
//
// Returns:
//   - None. The entry is passed directly to the logger for serialization and dispatch.
func Entry(ctx context.Context, entry any) {
	write(ctx, entry)
}

// write is a low-level utility that emits the final log event to the logger backend.
//
// This internal function abstracts the actual call to the logging system (e.g., slog, zap, zerolog).
// It is not meant to be used directly. Instead, prefer higher-level helpers like Info, Warn, or Error.
//
// Parameters:
//   - ctx (context.Context): The context for metadata propagation.
//   - entry (any): The structured log event to be emitted.
//
// Returns:
//   - None. Side-effect: sends the log entry to the logger.
func write(ctx context.Context, entry any) {
	level := extractLogLevel(entry)
	if !shouldLog(level) {
		return
	}
	if logger == nil {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}
	logger.Log(ctx, mapLogLevel(level), "log", slog.Any("event", entry))
}

func mapLogLevel(level Level.LogLevel) slog.Level {
	switch level {
	case Level.Trace, Level.Debug:
		return slog.LevelDebug
	case Level.Info:
		return slog.LevelInfo
	case Level.Warn:
		return slog.LevelWarn
	case Level.Error:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func (c *Config) applyDefaults() {
	if c.Writer == nil {
		c.Writer = os.Stdout
	}
	if c.Format == "" {
		c.Format = FormatJSON
	}
	if c.MinimumLogLevel == "" {
		c.MinimumLogLevel = Level.Info
	}
}

func extractLogLevel(entry any) Level.LogLevel {
	if e, ok := entry.(interface{ GetLevel() Level.LogLevel }); ok {
		return e.GetLevel()
	}
	// fallback: assume info
	return Level.Info
}

func shouldLog(level Level.LogLevel) bool {
	return Level.LogLevelPriority[level] >= Level.LogLevelPriority[minimumLogLevel]
}
