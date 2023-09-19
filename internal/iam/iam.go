package iam

import (
	"context"
	"fmt"

	"github.com/gaganchawara/loans/internal/enums/accounttype"
	"github.com/gaganchawara/loans/internal/errorcode"
	"github.com/gaganchawara/loans/pkg/errors"
	"google.golang.org/grpc"
)

// This is Identity and Access Management (IAM) layer.
// it ensures that users can only access data and perform actions that they are authorized to.

// UserRouteAccess is a map that defines which gRPC methods are accessible to users.
var UserRouteAccess = map[string]bool{
	"/loans.v1.LoansAPI/GetLoans":    true,
	"/loans.v1.LoansAPI/ApplyLoan":   true,
	"/loans.v1.LoansAPI/ApproveLoan": false,
	"/loans.v1.LoansAPI/RejectLoan":  false,
}

// UserAccessInterceptor is a gRPC interceptor that checks whether a user has access
// to a specific gRPC method based on their account type and the method's availability
// in the UserRouteAccess map. If the user does not have access, it returns an
// authorization error.
func UserAccessInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {
		accountType, ierr := GetAccountType(ctx)
		if ierr != nil {
			return nil, ierr
		}

		if accountType == accounttype.User {
			_, ierr := GetUserId(ctx)
			if ierr != nil {
				return nil, errors.New(ctx, errorcode.AuthenticationError, ierr).Report()
			}

			method, ierr := GetMethod(ctx)
			if ierr != nil {
				return nil, ierr
			}

			userAllowed, ok := UserRouteAccess[method]
			if !ok {
				return "", errors.New(ctx, errorcode.InternalServerError,
					fmt.Errorf("route access not available")).Report()
			}

			if !userAllowed {
				return nil, errors.New(ctx, errorcode.AuthorizationError,
					fmt.Errorf("the access is not allowed for the user")).Report()
			}
		} else {
			_, ierr := GetAdminId(ctx)
			if ierr != nil {
				return nil, errors.New(ctx, errorcode.AuthenticationError, ierr).Report()
			}
		}

		return handler(ctx, req)
	}
}

func ValidateAccessToUser(ctx context.Context, userID string) errors.Error {
	accountType, ierr := GetAccountType(ctx)
	if ierr != nil {
		return ierr
	}

	if accountType == accounttype.Admin {
		return nil
	}

	accessorUserID, ierr := GetUserId(ctx)
	if ierr != nil {
		return ierr
	}

	if accessorUserID != userID {
		return errors.New(ctx, errorcode.AuthenticationError, ierr).Report()
	}

	return nil
}
