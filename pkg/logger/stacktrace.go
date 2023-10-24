package logger

import (
	"bytes"
	"runtime/debug"
	"strings"

	"github.com/sirupsen/logrus"
)

// StackTraceHook is a Logrus hook that captures stack traces for specific log levels.
type StackTraceHook struct{}

// Levels returns the log levels for which the hook captures stack traces.
func (hook *StackTraceHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel}
}

// Fire captures and attaches the stack trace with the first 27 lines removed to the log entry.
func (hook *StackTraceHook) Fire(entry *logrus.Entry) error {
	stack := debug.Stack()
	stack = bytes.ReplaceAll(stack, []byte("\t"), []byte("")) // Remove all tabs
	stackArray := strings.Split(string(stack), "\n")
	stackArray = stackArray[19:]
	entry.Data["stack_trace"] = stackArray
	return nil
}
