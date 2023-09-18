package tracing

import "github.com/sirupsen/logrus"

type jLogger struct {
	logger *logrus.Entry
}

func (j *jLogger) Infof(msg string, args ...interface{}) {
	return
}

func (j *jLogger) Debugf(msg string, args ...interface{}) {
	j.logger.Debugf(msg, args)
}

func (j *jLogger) Error(msg string) {
	j.logger.Errorf(msg)
}
