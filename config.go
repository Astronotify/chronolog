package chronolog

import (
	"io"

	Level "github.com/Astronotify/chronolog/level"
)

type Format string

const (
	FormatJSON   Format = "json"
	FormatPretty Format = "pretty"
)

type Config struct {
	Writer          io.Writer
	Format          Format
	MinimumLogLevel Level.LogLevel
}
