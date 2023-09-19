package main

import (
	"context"
	"github.com/gaganchawara/loans/internal/iam"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gaganchawara/loans/internal/loan/interfaces"
	"github.com/gaganchawara/loans/internal/loan/repository"
	loanserver "github.com/gaganchawara/loans/internal/loan/server"
	"github.com/gaganchawara/loans/internal/loan/service"
	loansv1 "github.com/gaganchawara/loans/rpc/loans/v1"

	"github.com/gaganchawara/loans/internal/errorcode"

	ctxkeys "github.com/gaganchawara/loans/internal/constants/ctx_keys"

	interceptors "github.com/gaganchawara/loans/pkg/grpcserver/interceptor"

	"github.com/gaganchawara/loans/internal/boot"
	"github.com/gaganchawara/loans/pkg/grpcserver"
	"github.com/gaganchawara/loans/pkg/grpcserver/serverMux"
	"github.com/gaganchawara/loans/pkg/health"
	healthv1 "github.com/gaganchawara/loans/rpc/common/health/v1"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	ierr := boot.Initialize(ctx)
	if ierr != nil {
		log.Fatalf("application boot terminated due to error %v", ierr.Error())
		return
	}

	healthCore := health.NewCore(boot.DB)

	loansRepo := repository.NewRepository(boot.DB)
	loansSvc := service.NewService(loansRepo)

	server, ierr := grpcserver.NewServer(ctx,
		boot.Config.App.ServerAddresses,
		grpcServerRegisterer(healthCore, loansSvc),
		httpHandlerRegisterer(ctx),
		getServerInterceptors(),
		serverMux.DefaultServerMux(),
		getHttpMiddlewares(),
		registerInternalHandler(),
	)
	if ierr != nil {
		log.Fatalf("application boot terminated due to error %v", ierr.Error())
		return
	}

	grpcprometheus.Register(server.GetGrpcServer())

	ierr = server.Start(ctx)
	if ierr != nil {
		log.Fatalf("application boot terminated due to error %v", ierr.Error())
		return
	}

	// Handle SIGINT & SIGTERM - Shutdown gracefully
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	// Block until signal is received.
	<-c
	ierr = server.Stop(ctx, boot.Config.App.ShutdownDelay)
	if ierr != nil {
		log.Fatalf("application boot terminated due to error %v", ierr.Error())
		return
	}
}

func grpcServerRegisterer(healthCore *health.Core, loansSvc interfaces.Service) grpcserver.RegisterGrpcHandlers {
	return func(grpcServer *grpc.Server) error {
		healthv1.RegisterHealthCheckAPIServer(grpcServer, health.NewServer(healthCore))
		loansv1.RegisterLoansAPIServer(grpcServer, loanserver.NewServer(loansSvc))

		return nil
	}
}

func httpHandlerRegisterer(ctx context.Context) grpcserver.RegisterHttpHandlers {
	return func(mux *runtime.ServeMux, address string) error {
		if err := healthv1.RegisterHealthCheckAPIHandlerFromEndpoint(ctx, mux, address,
			[]grpc.DialOption{grpc.WithInsecure()}); err != nil {
			return err
		}
		if err := loansv1.RegisterLoansAPIHandlerFromEndpoint(ctx, mux, address,
			[]grpc.DialOption{grpc.WithInsecure()}); err != nil {
			return err
		}

		return nil
	}
}

func getServerInterceptors() []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		interceptors.UnaryServerBasicAuthInterceptor(
			[]interceptors.BasicAuthCreds{
				boot.Config.Auth,
			}),
		interceptors.HeaderInterceptor(ctxkeys.HeaderKeyMap()),
		interceptors.UnaryServerTraceIdInterceptor(),
		iam.UserAccessInterceptor(),
		interceptors.UnaryServerGitCommitHashInterceptor(boot.Config.App.GitCommitHash),
		interceptors.UnaryServerLoggerInterceptor(ctxkeys.AllKeys()),
		interceptors.UnaryServerGrpcErrorInterceptor(errorcode.ErrorsMap),
		grpcprometheus.UnaryServerInterceptor,
		interceptors.PanicInterceptor(),
	}
}

func getHttpMiddlewares() []grpcserver.HttpMiddlewares {
	return nil
}

func registerInternalHandler() grpcserver.RegisterInternalHandler {
	return func(mux *http.ServeMux) *http.ServeMux {
		mux.Handle("/metrics", promhttp.Handler())
		return mux
	}
}
