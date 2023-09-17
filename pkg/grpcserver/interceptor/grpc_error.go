package interceptors

import (
	"context"

	"github.com/gaganchawara/loans/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errorMap map[string]codes.Code

// UnaryServerGrpcErrorInterceptor returns a grpc.UnaryServerInterceptor suitable for converting
// errors.Error to grpc status error
func UnaryServerGrpcErrorInterceptor(errMap map[string]codes.Code) grpc.UnaryServerInterceptor {
	errorMap = errMap
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			ierr, b := err.(errors.Error)
			if b == false {
				return resp, status.Error(codes.Unknown, err.Error())
			}

			code, b := errorMap[ierr.Code()]
			if b == false {
				code = codes.Unknown
			}

			return resp, status.Error(code, ierr.Error())
		}

		return resp, err
	}
}
