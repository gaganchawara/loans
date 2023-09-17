package interceptors

import (
	"context"

	"github.com/gaganchawara/loans/pkg/utils"
	"google.golang.org/grpc"
)

const (
	TraceId = "trace_id"
)

// UnaryServerTraceIdInterceptor generates a trace ID and sets it in the gRPC context.
func UnaryServerTraceIdInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		ctx = context.WithValue(ctx, TraceId, utils.GenerateRandomId())
		return handler(ctx, req)
	}
}
