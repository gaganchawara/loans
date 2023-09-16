package errors

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var counter = 0

func instrumentErrors() func(iError Error) {
	return func(iError Error) {
		// sentry.GetHubFromContext(iError.Context()).CaptureException(iError)
		counter++
	}
}

func TestError(t *testing.T) {
	Initialize(instrumentErrors())

	err := New(context.Background(), "INTERNAL_SERVER_ERROR", errors.New("oho")).WithField(
		"key", "value").Report()

	assert.Equal(t, "oho", err.Error())
	assert.Equal(t, "INTERNAL_SERVER_ERROR", err.Code())
	assert.Equal(t, 1, counter)
}

func TestErrorUnwrap(t *testing.T) {
	Initialize(instrumentErrors())

	ErrOldPassword := errors.New("this user requires old password authentication")

	err := New(context.Background(), "INTERNAL_SERVER_ERROR", ErrOldPassword).WithField(
		"key", "value").Report()

	assert.True(t, errors.Is(err, ErrOldPassword))
}
