package interceptors

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// HeaderInterceptor extracts and sets values from incoming metadata headers.
func HeaderInterceptor(headerTags []string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = extractAndSetMetadataValues(ctx, headerTags)
		return handler(ctx, req)
	}
}

// extractAndSetMetadataValues extracts values from metadata and sets them as context values.
func extractAndSetMetadataValues(ctx context.Context, headerTags []string) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx
	}

	for key, values := range md {
		if in(key, headerTags) && len(values) >= 1 {
			ctx = context.WithValue(ctx, key, values[0])
		}
	}

	return ctx
}

// in checks if a string exists in a slice of strings (case-insensitive).
func in(needle string, haystack []string) bool {
	for _, k := range haystack {
		if strings.ToLower(k) == needle {
			return true
		}
	}

	return false
}
