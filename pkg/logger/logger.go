package logger

import (
	"context"
	"os"
	"time"

	"github.com/gaganchawara/loans/pkg/errors"
	"github.com/sirupsen/logrus"
)

type contextKey int

const (
	CtxKey contextKey = iota
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})
	logrus.SetOutput(os.Stdout)
	logrus.AddHook(&StackTraceHook{})
}

func Get(ctx context.Context) *logrus.Entry {
	if ctx == nil {
		return newStdLogger()
	}

	k, ok := ctx.Value(CtxKey).(*logrus.Entry)
	if ok {
		return k
	} else {
		return newStdLogger().WithContext(ctx)
	}
}

func newStdLogger() *logrus.Entry {
	return logrus.NewEntry(logrus.StandardLogger())
}

func ErrorLogger() func(error errors.Error) {
	return func(error errors.Error) {
		data := map[string]interface{}{}
		for k, v := range error.Data() {
			data[k] = v
		}

		Get(error.Context()).WithFields(data).WithField("code", error.Code()).Error(error.Error())
	}
}
