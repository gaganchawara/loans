package interceptors

import (
	"context"
	"fmt"

	"github.com/gaganchawara/loans/internal/errorcode"
	"github.com/gaganchawara/loans/pkg/errors"
	"google.golang.org/grpc"
)

// PanicInterceptor is a gRPC unary server interceptor that extracts and sets values from incoming metadata.
// The `headerKeyMap` parameter expects a map in which the key is the header/metadata key, and the value is the key
// in which the value needs to be set in the context.
func PanicInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				// Handle the panic and return an error with Internal gRPC status code
				err = errors.New(ctx, errorcode.InternalServerPanic, fmt.Errorf("error: %v", r)).Report()
			}
		}()

		resp, err = handler(ctx, req)
		return resp, err
	}
}
