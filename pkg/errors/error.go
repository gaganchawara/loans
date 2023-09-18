package errors

import (
	"context"
	"errors"
)

// Error is an interface representing application-specific errors.
type Error interface {
	Error() string
	Code() string
	Unwrap() error
	WithField(key string, val string) Error
	WithData(map[string]string) Error
	Context() context.Context
	Data() map[string]string
	Report() Error
}

// Hook is a function signature for error hooks.
type Hook func(Error)

var hooks []Hook

// Initialize initializes error hooks
func Initialize(h ...Hook) {
	hooks = h
}

// e is a concrete implementation of the Error interface.
type e struct {
	ctx  context.Context
	code string
	err  error
	data map[string]string
}

// New creates a new error instance with context, error code, and an optional error.
func New(ctx context.Context, code string, err error) Error {
	if err == nil {
		err = errors.New(code)
	}
	return &e{
		ctx:  ctx,
		code: code,
		err:  err,
		data: map[string]string{},
	}
}

// WithField adds a key-value pair to the error's data.
func (e *e) WithField(key, value string) Error {
	e.data[key] = value

	return e
}

// WithData adds multiple key-value pairs to the error's data.
func (e *e) WithData(data map[string]string) Error {
	for k, v := range data {
		e.data[k] = v
	}

	return e
}

// Report executes registered error hooks.
func (e *e) Report() Error {
	for _, hook := range hooks {
		hook(e)
	}

	return e
}

// Data returns the error's associated data.
func (e *e) Data() map[string]string {
	return e.data
}

// Error returns the error message or code.
func (e *e) Error() string {
	if e.err != nil {
		return e.err.Error()
	}

	return e.code
}

// Unwrap returns the wrapped error, if any.
func (e *e) Unwrap() error {
	return e.err
}

// Code returns the error code.
func (e *e) Code() string {
	return e.code
}

// Context returns the associated context, or a default context if none is set.
func (e *e) Context() context.Context {
	if e.ctx == nil {
		e.ctx = context.Background()
	}
	return e.ctx
}
