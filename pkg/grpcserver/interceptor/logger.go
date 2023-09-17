package interceptors

import (
	"context"

	"github.com/gaganchawara/loans/pkg/logger"
	"google.golang.org/grpc"
)

const contextKey = "context"

func UnaryServerLoggerInterceptor(keys []string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		var fields = make(map[string]interface{}, 0)

		for _, key := range keys {
			fields[key] = ctx.Value(key)
		}

		lgr := logger.Get(ctx)

		ctx = context.WithValue(ctx, logger.CtxKey, lgr.WithFields(map[string]interface{}{
			contextKey: fields,
		}))

		return handler(ctx, req)
	}
}
