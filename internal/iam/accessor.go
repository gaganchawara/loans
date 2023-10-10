package iam

import (
	"context"
	"fmt"

	ctxkeys "github.com/gaganchawara/loans/internal/constants/ctx_keys"
	"github.com/gaganchawara/loans/internal/enums/accounttype"
	"github.com/gaganchawara/loans/internal/errorcode"
	"github.com/gaganchawara/loans/pkg/errors"
)

// GetAccountType retrieves the account type from the context and returns it.
// If the account type is not available or cannot be parsed, it returns an error.
func GetAccountType(ctx context.Context) (accounttype.Type, errors.Error) {
	accountTypeStr, ok := ctx.Value(ctxkeys.AccountType).(string)
	if !ok {
		return accounttype.Type(0), errors.New(ctx, errorcode.AuthenticationError,
			fmt.Errorf("account type not available")).Report()
	}

	accountType, ierr := accounttype.TypeFromString(ctx, accountTypeStr)
	if ierr != nil {
		return accounttype.Type(0), ierr
	}

	return accountType, nil
}

// GetMethod retrieves the gRPC method name from the context and returns it.
// If the method name is not available, it returns an error.
func GetMethod(ctx context.Context) (string, errors.Error) {
	method, ok := ctx.Value(ctxkeys.RpcMethodKey).(string)
	if !ok {
		return "", errors.New(ctx, errorcode.InternalServerError,
			fmt.Errorf("route not available")).Report()
	}

	return method, nil
}

// GetAdminId retrieves the Admin ID name from the context and returns it.
func GetAdminId(ctx context.Context) (string, errors.Error) {
	adminId, ok := ctx.Value(ctxkeys.AdminID).(string)
	if !ok {
		return "", errors.New(ctx, errorcode.InternalServerError,
			fmt.Errorf("admin id not available")).Report()
	}

	return adminId, nil
}

// GetUserId retrieves the Admin ID name from the context and returns it.
func GetUserId(ctx context.Context) (string, errors.Error) {
	userID, ok := ctx.Value(ctxkeys.UserID).(string)
	if !ok {
		return "", errors.New(ctx, errorcode.InternalServerError,
			fmt.Errorf("user id not available")).Report()
	}

	return userID, nil
}
