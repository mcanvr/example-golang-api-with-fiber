package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// LogLevel represents the severity level of a log message
type LogLevel int

// Log levels
const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

var logLevelNames = []string{
	"DEBUG",
	"INFO",
	"WARN",
	"ERROR",
	"FATAL",
}

// Logger is a structured logger
type Logger struct {
	level  LogLevel
	prefix string
	logger *log.Logger
}

// Configuration for the logger
type Config struct {
	Level  LogLevel
	Prefix string
	Output io.Writer
}

// DefaultConfig returns a default logger configuration
func DefaultConfig() Config {
	return Config{
		Level:  INFO,
		Prefix: "",
		Output: os.Stdout,
	}
}

// NewLogger creates a new logger with the provided configuration
func NewLogger(config Config) *Logger {
	flags := log.Ldate | log.Ltime
	logger := log.New(config.Output, config.Prefix, flags)

	return &Logger{
		level:  config.Level,
		prefix: config.Prefix,
		logger: logger,
	}
}

// formatMessage formats a log message with timestamp and level
func (l *Logger) formatMessage(level LogLevel, format string, args ...interface{}) string {
	timestamp := time.Now().Format("2006/01/02 15:04:05")
	prefix := fmt.Sprintf("[%s] [%s] ", timestamp, logLevelNames[level])
	message := fmt.Sprintf(format, args...)
	return prefix + message
}

// log logs a message at the specified level
func (l *Logger) log(level LogLevel, format string, args ...interface{}) {
	if level >= l.level {
		msg := l.formatMessage(level, format, args...)
		l.logger.Println(msg)
	}
}

// Debug logs a debug message
func (l *Logger) Debug(format string, args ...interface{}) {
	l.log(DEBUG, format, args...)
}

// Info logs an info message
func (l *Logger) Info(format string, args ...interface{}) {
	l.log(INFO, format, args...)
}

// Warn logs a warning message
func (l *Logger) Warn(format string, args ...interface{}) {
	l.log(WARN, format, args...)
}

// Error logs an error message
func (l *Logger) Error(format string, args ...interface{}) {
	l.log(ERROR, format, args...)
}

// Fatal logs a fatal message and exits
func (l *Logger) Fatal(format string, args ...interface{}) {
	l.log(FATAL, format, args...)
	os.Exit(1)
}

// Default global logger
var defaultLogger = NewLogger(DefaultConfig())

// SetDefaultLogger sets the default global logger
func SetDefaultLogger(logger *Logger) {
	defaultLogger = logger
}

// Global logging functions

// Debug logs a debug message using the default logger
func Debug(format string, args ...interface{}) {
	defaultLogger.Debug(format, args...)
}

// Info logs an info message using the default logger
func Info(format string, args ...interface{}) {
	defaultLogger.Info(format, args...)
}

// Warn logs a warning message using the default logger
func Warn(format string, args ...interface{}) {
	defaultLogger.Warn(format, args...)
}

// Error logs an error message using the default logger
func Error(format string, args ...interface{}) {
	defaultLogger.Error(format, args...)
}

// Fatal logs a fatal message using the default logger and exits
func Fatal(format string, args ...interface{}) {
	defaultLogger.Fatal(format, args...)
}
