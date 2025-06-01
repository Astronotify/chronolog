package chronolog

import (
	"io"

	"github.com/mvleandro/chronolog/entries"
)

type Format string

const (
	FormatJSON   Format = "json"
	FormatPretty Format = "pretty"
)

type Config struct {
	Writer          io.Writer
	Format          Format
	MinimumLogLevel entries.LogLevel
}
