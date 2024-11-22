package gogger

import (
	"io"
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

// LogLevel is the log level.
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
)

// Logger is the logger. It wraps the zerolog logger.
type Logger struct {
	zerolog.Logger
	file *os.File
}

// Close closes the logger.
func (l *Logger) Close() error {
	if l.file != nil {
		return l.file.Close()
	}
	return nil
}

// SetLevel sets the log level.
func (l *Logger) SetLevel(level LogLevel) {
	zerolog.SetGlobalLevel(getZerologLevel(level))
	l.Logger = l.Logger.Level(getZerologLevel(level))
}

// LogStructuredData logs structured data.
func (l *Logger) LogStructuredData(data map[string]interface{}) {
	event := l.Info()
	for key, value := range data {
		event = event.Interface(key, value)
	}
	event.Msg("Structured data")
}

// LogErrorWithStack logs an error with stack trace.
func (l *Logger) LogErrorWithStack(err error) {
	l.Error().Stack().Err(err).Msg("An error occurred")
}

// NewLogger creates a new logger with the given output path and log level.
func NewLogger(outputPath string, level LogLevel) (*Logger, error) {
	// Create log file directory if not exists
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return nil, err
	}

	file, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		return nil, err
	}

	var writer io.Writer = file

	// Multi-level writer for console and file
	multi := zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stdout}, writer)

	// Set up zerolog logger with output to both console and file
	logger := log.Output(multi).Level(getZerologLevel(level)).With().Timestamp().Logger()

	// Add stack trace to errors
	logger = logger.With().Stack().Logger()

	// Set zerolog global logger
	zerolog.SetGlobalLevel(getZerologLevel(level))
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	return &Logger{logger, file}, nil
}

// Config is the configuration for the logger.
type Config struct {
	OutputPath string   // Path to the log file (stdout for console)
	Level      LogLevel // Log level (DEBUG, INFO, WARNING, ERROR)
	Format     string   // Log format (json, text)
}

// NewLoggerWithConfig creates a new logger with the given configuration.
func NewLoggerWithConfig(config Config) (*Logger, error) {
	if config.OutputPath == "stdout" {
		config.OutputPath = os.Stdout.Name()
	}

	logger, err := NewLogger(config.OutputPath, config.Level)
	if err != nil {
		return nil, err
	}

	if config.Format == "json" {
		logger.Logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: true})
	}

	return logger, nil
}

// getZerologLevel converts the log level to zerolog level.
func getZerologLevel(level LogLevel) zerolog.Level {
	switch level {
	case DEBUG:
		return zerolog.DebugLevel
	case INFO:
		return zerolog.InfoLevel
	case WARNING:
		return zerolog.WarnLevel
	case ERROR:
		return zerolog.ErrorLevel
	default:
		return zerolog.InfoLevel
	}
}
