package helper

import (
	"testing"

	"github.com/gaganchawara/loans/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func AssertEqualError(t *testing.T, expectedErr errors.Error, actualErr errors.Error) bool {
	if expectedErr == nil {
		return assert.Nil(t, actualErr)
	}

	b := assert.Equal(t, expectedErr.Code(), actualErr.Code())
	b = b && assert.Equal(t, expectedErr.Error(), actualErr.Error())
	b = b && assert.Equal(t, expectedErr.Data(), actualErr.Data())

	return b
}
