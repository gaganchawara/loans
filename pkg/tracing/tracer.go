package tracing

import (
	"context"

	"github.com/gaganchawara/loans/internal/errorcode"
	"github.com/gaganchawara/loans/pkg/errors"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

type Config struct {
	ServiceName string
	Env         string
	HostName    string
	JaegerPort  string
}

// InitTracer initializes the OpenTracing tracer with a custom configuration.
func InitTracer(ctx context.Context, config Config, entry *logrus.Entry) errors.Error {
	jcfg := &jaegercfg.Configuration{
		ServiceName: config.ServiceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeRateLimiting,
			Param: 100,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: config.HostName + ":" + config.JaegerPort,
		},
	}

	tracer, _, err := jcfg.NewTracer(
		jaegercfg.Logger(&jLogger{entry}),
	)
	if err != nil {
		return errors.New(ctx, errorcode.InternalServerError, err).Report()
	}

	opentracing.SetGlobalTracer(tracer)

	return nil
}
