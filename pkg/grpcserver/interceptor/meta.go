package interceptors

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// HeaderInterceptor is a gRPC unary server interceptor that extracts and sets values from incoming metadata.
// The `headerKeyMap` parameter expects a map in which the key is the header/metadata key, and the value is the key
// in which the value needs to be set in the context.
func HeaderInterceptor(headerKeyMap map[string]string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = extractAndSetMetadataValues(ctx, headerKeyMap)
		return handler(ctx, req)
	}
}

// extractAndSetMetadataValues extracts values from metadata and sets them as context values.
func extractAndSetMetadataValues(ctx context.Context, m map[string]string) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx
	}

	haystack := GetMapKeys(m)

	for key, values := range md {
		if in(key, haystack) && len(values) >= 1 {
			ctx = context.WithValue(ctx, m[key], values[0])
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

// GetMapKeys returns the keys of a map as a slice of strings.
func GetMapKeys(inputMap map[string]string) []string {
	keys := make([]string, len(inputMap))
	i := 0
	for key := range inputMap {
		keys[i] = key
		i++
	}

	return keys
}
