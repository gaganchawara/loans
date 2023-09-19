package accounttype

import (
	"context"
	"fmt"

	"github.com/gaganchawara/loans/internal/errorcode"
	"github.com/gaganchawara/loans/pkg/errors"
)

type Type uint8

const (
	User Type = iota + 1
	Admin
)

var typeToString = map[Type]string{
	User:  "USER",
	Admin: "ADMIN",
}

var typeFromString = map[string]Type{
	"USER":  User,
	"ADMIN": Admin,
}

func TypeFromString(ctx context.Context, s string) (Type, errors.Error) {
	val, ok := typeFromString[s]
	if !ok {
		var garbage Type
		return garbage, errors.New(ctx, errorcode.BadRequestError, fmt.Errorf("invalid account type")).Report()
	}

	return val, nil
}

func All() []Type {
	var statusList []Type
	for _, status := range typeFromString {
		statusList = append(statusList, status)
	}
	return statusList
}

func AllString() []string {
	var statusList []string
	for _, status := range typeToString {
		statusList = append(statusList, status)
	}

	return statusList
}

func (s Type) String() string {
	return typeToString[s]
}
