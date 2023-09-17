package logger

import (
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

// StackTraceHook is a Logrus hook that captures stack traces for specific log levels.
type StackTraceHook struct{}

// Levels returns the log levels for which the hook captures stack traces.
func (hook *StackTraceHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel}
}

// Fire captures and attaches a stack trace to the log entry.
func (hook *StackTraceHook) Fire(entry *logrus.Entry) error {
	stack := debug.Stack()
	entry.Data["stack_trace"] = string(stack)
	return nil
}
