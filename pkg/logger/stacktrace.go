package logger

import (
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

type StackTraceHook struct{}

func (hook *StackTraceHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel}
}

func (hook *StackTraceHook) Fire(entry *logrus.Entry) error {
	stack := debug.Stack()
	entry.Data["stack_trace"] = string(stack)
	return nil
}
