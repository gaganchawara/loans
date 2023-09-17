package grpcserver

import (
	"context"
	"net/http"

	"github.com/gaganchawara/loans/internal/errorcode"
	"github.com/gaganchawara/loans/pkg/errors"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

const (
	RpcMethodKey = "method"
	UriKey       = "uri"
)

// ServerAddresses defines the network addresses for the different server interfaces.
type ServerAddresses struct {
	Grpc     string
	Http     string
	Internal string
}

// Server represents the application's server configuration, maintaining the following network interfaces:
//   - grpcServer: The core gRPC server instance.
//   - httpServer: The HTTP server instance used by grpc-gateway.
//   - internalServer: An HTTP server exposed internally, typically used for sharing data with internal services,
//     such as Prometheus metrics.
type Server struct {
	serverAddresses ServerAddresses
	internalServer  *http.Server
	grpcServer      *grpc.Server
	httpServer      *http.Server
}

// NewServer creates an instance of Server struct. registerGrpcHandlers and registerHttpHandlers should be used
// for registering the grpc handlers and http handlers for grpc-gateway respectively. You can also supply a list of
// interceptors to use for the server.
func NewServer(ctx context.Context,
	serverAddresses ServerAddresses,
	registerGrpcHandlers RegisterGrpcHandlers,
	registerHttpHandlers RegisterHttpHandlers,
	serverInterceptors []grpc.UnaryServerInterceptor,
	muxOptions []runtime.ServeMuxOption,
	httpMiddlewares []HttpMiddlewares,
	registerInternalHandlers RegisterInternalHandler) (*Server, errors.Error) {

	grpcServer, ierr := newGrpcServer(ctx, serverAddresses, registerGrpcHandlers, serverInterceptors)
	if ierr != nil {
		return nil, ierr
	}

	httpServer, ierr := newHttpServer(ctx, serverAddresses, registerHttpHandlers, muxOptions, httpMiddlewares)
	if ierr != nil {
		return nil, ierr
	}

	internalServer, ierr := newInternalServer(serverAddresses, registerInternalHandlers)
	if ierr != nil {
		return nil, ierr
	}

	return &Server{
		serverAddresses: serverAddresses,
		internalServer:  internalServer,
		grpcServer:      grpcServer,
		httpServer:      httpServer,
	}, nil
}

// newGrpcServer creates and configures a new gRPC server instance.
// It uses the provided RegisterGrpcHandlers function to register gRPC handlers and applies
// any specified server interceptors to the server.
func newGrpcServer(ctx context.Context, _ ServerAddresses, r RegisterGrpcHandlers,
	serverInterceptors []grpc.UnaryServerInterceptor) (*grpc.Server, errors.Error) {
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcmiddleware.ChainUnaryServer(serverInterceptors...),
		),
	)

	err := r(grpcServer)
	if err != nil {
		return nil, errors.New(ctx, errorcode.InternalServerError, err).Report()
	}

	return grpcServer, nil
}

// newHttpServer creates and configures a new HTTP server instance for grpc-gateway.
// It uses the provided RegisterHttpHandlers function to register proxy handlers for gRPC-gateway,
// along with optional middleware functions and ServeMux options.
func newHttpServer(ctx context.Context, addresses ServerAddresses, r RegisterHttpHandlers,
	muxOptions []runtime.ServeMuxOption,
	httpMiddlewares []HttpMiddlewares) (*http.Server, errors.Error) {
	// MarshalOptions is added to retain the same response format as grpc-gateway v1 https://grpc-ecosystem.github.io/grpc-gateway/docs/development/grpc-gateway_v2_migration_guide/#we-now-use-the-camelcase-json-names-by-default
	mux := runtime.NewServeMux(
		muxOptions...,
	)

	err := r(mux, addresses.Grpc)
	if err != nil {
		return nil, errors.New(ctx, errorcode.InternalServerError, err).Report()
	}

	handler := http.Handler(mux)
	for _, m := range httpMiddlewares {
		handler = m(handler)
	}

	server := http.Server{Handler: handler}
	return &server, nil
}

// newInternalServer creates and configures a new internal HTTP server instance.
// It uses the provided RegisterInternalHandler function to register internal HTTP handlers
// on a ServeMux, typically used for exposing data to internal services, such as Prometheus metrics.
func newInternalServer(_ ServerAddresses, handler RegisterInternalHandler) (*http.Server, errors.Error) {
	mux := http.NewServeMux()

	mux = handler(mux)

	server := http.Server{Handler: mux}

	return &server, nil
}

// GetHttpServer returns the HTTP server struct.
func (s *Server) GetHttpServer() *http.Server {
	return s.httpServer
}

// GetGrpcServer returns the gRPC server struct.
func (s *Server) GetGrpcServer() *grpc.Server {
	return s.grpcServer
}

// GetInternalServer returns the internal HTTP server struct.
func (s *Server) GetInternalServer() *http.Server {
	return s.internalServer
}
