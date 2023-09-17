package serverMux

import (
	"context"
	"net/http"
	"strings"

	"github.com/gaganchawara/loans/pkg/grpcserver"
	"google.golang.org/grpc/metadata"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/encoding/protojson"
)

func DefaultServerMux() []runtime.ServeMuxOption {
	return []runtime.ServeMuxOption{
		// pushing method and uri data to meta so that they can be used for logging
		runtime.WithMetadata(func(ctx context.Context, r *http.Request) metadata.MD {
			md := make(map[string]string)
			if method, ok := runtime.RPCMethod(ctx); ok {
				md[grpcserver.RpcMethodKey] = method
			}
			if pattern, ok := runtime.HTTPPathPattern(ctx); ok {
				md[grpcserver.UriKey] = pattern
			}
			return metadata.New(md)
		}),
		DefaultMarshalOption(),
		DefaultHeaderMatcher(),
	}
}

func DefaultMarshalOption() runtime.ServeMuxOption {
	return runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions:   protojson.MarshalOptions{UseProtoNames: true, EmitUnpopulated: true},
		UnmarshalOptions: protojson.UnmarshalOptions{DiscardUnknown: true},
	})
}

func DefaultHeaderMatcher() runtime.ServeMuxOption {
	return runtime.WithIncomingHeaderMatcher(headerMatcher)
}

func headerMatcher(key string) (string, bool) {
	// Credits: https://github.com/argoproj/argo-rollouts/pull/1862/files
	// Dropping "Connection" header as a workaround for https://github.com/grpc-ecosystem/grpc-gateway/issues/2447
	// The fix is part of grpc-gateway v2.x but not available in v1.x, so workaround should be removed after upgrading to grpc v2.x
	return key, strings.ToLower(key) != "connection"
}
