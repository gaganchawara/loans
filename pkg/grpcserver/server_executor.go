package grpcserver

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/gaganchawara/loans/internal/errorcode"
	"github.com/gaganchawara/loans/pkg/errors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

// Start listening on all 3 net interfaces.
//
// Please note that Start is a blocking call and will not return until all servers stop or there is an error on one.
// To stop gracefully, you need to call Stop from another goroutine.
func (s *Server) Start(ctx context.Context) errors.Error {
	g, ctx := errgroup.WithContext(ctx)

	// Start the gRPC server.
	grpcListener, err := net.Listen("tcp", s.serverAddresses.Grpc)
	if err != nil {
		return errors.New(ctx, errorcode.InternalServerError, err).Report()
	}

	g.Go(func() error {
		return s.grpcServer.Serve(grpcListener)
	})

	// Start internal HTTP server. Used for exposing prometheus metrics.
	internalListener, err := net.Listen("tcp", s.serverAddresses.Internal)
	if err != nil {
		return errors.New(ctx, errorcode.InternalServerError, err).Report()
	}

	g.Go(func() error {
		err = s.internalServer.Serve(internalListener)
		if err != nil && err != http.ErrServerClosed {
			return err
		}

		return nil
	})

	// Start HTTP server for gRPC gateway.
	httpListener, err := net.Listen("tcp", s.serverAddresses.Http)
	if err != nil {
		return errors.New(ctx, errorcode.InternalServerError, err).Report()
	}

	g.Go(func() error {
		err = s.httpServer.Serve(httpListener)
		if err != nil && err != http.ErrServerClosed {
			return err
		}

		return nil
	})

	if err = g.Wait(); err != nil {
		return errors.New(ctx, errorcode.InternalServerError, err).Report()
	}

	return nil
}

// Stop will stop listening on the net interface. This will block until all open connections are drained.
func (s *Server) Stop(ctx context.Context, shutdownTimeoutSeconds int) errors.Error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(shutdownTimeoutSeconds)*time.Second)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)
	// gracefully shutdown grpc-server
	g.Go(func() error {
		return gracefullyShutdownGrpcServerWithTimeout(ctx, s.grpcServer)
	})

	// gracefully shutdown internal server
	g.Go(func() error {
		return s.internalServer.Shutdown(ctx)
	})

	// gracefully shutdown http server for gateway
	g.Go(func() error {
		return s.httpServer.Shutdown(ctx)
	})

	if err := g.Wait(); err != nil {
		return errors.New(ctx, errorcode.InternalServerError, err).Report()
	}

	return nil
}

// gracefullyShutdownGrpcServerWithTimeout manages graceful shutdown of a gRPC server with a timeout.
// Since the grpc.Server package does not provide a GracefulStop with a timeout option,
// this function initiates a GracefulStop in a background goroutine and simultaneously listens for
// context cancellation. If the context deadline is reached, it forcefully terminates connections using Stop.
func gracefullyShutdownGrpcServerWithTimeout(ctx context.Context, server *grpc.Server) error {
	stopped := make(chan bool)
	go func() {
		server.GracefulStop()
		close(stopped)
	}()

	select {
	case <-ctx.Done():
		server.Stop()
		return ctx.Err()
	case <-stopped:
		return nil
	}
}
