package logging

import (
	"log"
	"os"
)

// Logger wraps the standard library logger.
type Logger struct {
	*log.Logger
}

// New creates a new logger instance.
func New() *Logger {
	return &Logger{Logger: log.New(os.Stdout, "", log.LstdFlags)}
}
