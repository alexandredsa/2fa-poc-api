package applog

import (
	"log"
	"os"
)

// Logger is a struct that encapsulates a Go logger.
type Logger struct {
	logger *log.Logger
}

// NewLogger creates a new instance of the Logger struct.
func NewLogger(component string) *Logger {
	return &Logger{
		logger: log.New(os.Stdout, component, log.Ldate|log.Ltime),
	}
}

// Info logs an informational message.
func (l *Logger) Info(message string) {
	l.logger.Println("[INFO]", message)
}

// Error logs an error message.
func (l *Logger) Error(message string) {
	l.logger.Println("[ERROR]", message)
}
