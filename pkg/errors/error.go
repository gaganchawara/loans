package errors

import (
	"context"
	"errors"
)

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

type Hook func(Error)

var hooks []Hook

func Initialize(h ...Hook) {
	hooks = h
}

type Err struct {
	ctx   context.Context
	code  string
	err   error
	data  map[string]string
}

func New(ctx context.Context, code string, err error) Error {
	if err == nil {
		err = errors.New(code)
	}
	return &Err{
		ctx:   ctx,
		code:  code,
		err:   err,
		data:  map[string]string{},
	}
}

func (e *Err) WithField(key, value string) Error {
	e.data[key] = value

	return e
}

func (e *Err) WithData(data map[string]string) Error {
	for k, v := range data {
		e.data[k] = v
	}

	return e
}

func (e *Err) Report() Error {
	for _, hook := range hooks {
		hook(e)
	}

	return e
}

func (e *Err) Data() map[string]string {
	return e.data
}

func (e *Err) Error() string {
	if e.err != nil {
		return e.err.Error()
	}

	return e.code
}

func (e *Err) Unwrap() error {
	return e.err
}

func (e *Err) Code() string {
	return e.code
}

func (e *Err) Context() context.Context {
	if e.ctx == nil {
		e.ctx = context.Background()
	}
	return e.ctx
}
