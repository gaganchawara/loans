package health

import (
	"context"

	"github.com/gaganchawara/loans/internal/errorcode"
	"github.com/gaganchawara/loans/pkg/errors"

	healthv1 "github.com/gaganchawara/loans/rpc/common/health/v1"
)

// Server has methods implementing of server rpc.
type Server struct {
	healthv1.UnimplementedHealthCheckAPIServer
	core *Core
}

// NewServer returns a server.
func NewServer(core *Core) *Server {
	return &Server{
		core: core,
	}
}

// AuthFuncOverride interceptor to override default basic auth authentication
func (s *Server) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

// Check returns service's serving status.
func (s *Server) Check(ctx context.Context, _ *healthv1.HealthCheckRequest) (*healthv1.HealthCheckResponse, error) {
	err := s.core.RunHealthCheck(ctx)
	if err != nil {
		return &healthv1.HealthCheckResponse{ServingStatus: healthv1.HealthCheckResponse_SERVING_STATUS_NOT_SERVING},
			errors.New(ctx, errorcode.InternalServerError, err).Report()
	}

	return &healthv1.HealthCheckResponse{ServingStatus: healthv1.HealthCheckResponse_SERVING_STATUS_SERVING}, nil
}

func (s *Server) Ping(ctx context.Context, _ *healthv1.HealthCheckRequest) (*healthv1.HealthCheckResponse, error) {
	err := s.core.Ping(ctx)
	if err != nil {
		return &healthv1.HealthCheckResponse{ServingStatus: healthv1.HealthCheckResponse_SERVING_STATUS_NOT_SERVING},
			errors.New(ctx, errorcode.InternalServerError, err).Report()
	}

	return &healthv1.HealthCheckResponse{ServingStatus: healthv1.HealthCheckResponse_SERVING_STATUS_SERVING}, nil
}
