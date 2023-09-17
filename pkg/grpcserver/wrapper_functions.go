package grpcserver

import (
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// HttpMiddlewares is a function type that can be used to apply middleware to an HTTP handler.
type HttpMiddlewares func(handler http.Handler) http.Handler

// RegisterInternalHandler is a function type that registers internal HTTP handlers on a ServeMux.
type RegisterInternalHandler func(mux *http.ServeMux) *http.ServeMux

// RegisterGrpcHandlers is a function type used to register gRPC server handlers.
type RegisterGrpcHandlers func(server *grpc.Server) error

// RegisterHttpHandlers is a function type used to register gRPC-gateway proxy handlers.
type RegisterHttpHandlers func(mux *runtime.ServeMux, address string) error
