package loanstatus

import (
	"context"
	"database/sql/driver"
	"fmt"
	"github.com/gaganchawara/loans/internal/errorcode"
	"github.com/gaganchawara/loans/pkg/errors"
)

type Type uint8

const (
	Pending Type = iota + 1
	Approved
	PartiallyPaid
	Paid
)

var typeToString = map[Type]string{
	Pending:       "PENDING",
	Approved:      "APPROVED",
	PartiallyPaid: "PARTIALLY_PAID",
	Paid:          "PAID",
}

var typeFromString = map[string]Type{
	"PENDING":        Pending,
	"APPROVED":       Approved,
	"PARTIALLY_PAID": PartiallyPaid,
	"PAID":           Paid,
}

func TypeFromString(ctx context.Context, s string) (Type, errors.Error) {
	val, ok := typeFromString[s]
	if !ok {
		var garbage Type
		return garbage, errors.New(ctx, errorcode.BadRequestError,fmt.Errorf("invalid Market Segement")).Report()
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

func (s Type) Value() (driver.Value, error) {
	return s.String(), nil
}
