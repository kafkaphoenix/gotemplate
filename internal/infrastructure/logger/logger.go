package logger

import (
	"log/slog"
	"os"

	"github.com/rs/zerolog"
	slogzerolog "github.com/samber/slog-zerolog"
)

const (
	TraceLevel = iota - 1 // -1
	DebugLevel            // 0
	InfoLevel             // 1
	WarnLevel             // 2
	ErrorLevel            // 3
	FatalLevel            // 4
	PanicLevel            // 5
)

var zerologLevels = []zerolog.Level{
	zerolog.DebugLevel,
	zerolog.InfoLevel,
	zerolog.WarnLevel,
	zerolog.ErrorLevel,
	zerolog.FatalLevel,
	zerolog.PanicLevel,
}

// Init initializes the logger with zerolog as the backend
// and returns a slog.Logger instance to avoid direct dependency on zerolog.
func Init(logLevel int) *slog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if logLevel < DebugLevel || logLevel > PanicLevel {
		logLevel = InfoLevel // Default
	}

	_logger := zerolog.New(os.Stdout).
		Level(zerologLevels[logLevel]).
		With().Timestamp().Logger()

	return slog.New(
		slogzerolog.Option{Logger: &_logger}.NewZerologHandler(),
	)
}
