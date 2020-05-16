package tracer2

import (
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jagger "github.com/uber/jaeger-client-go/config"
)

func NewTracer(servicename string, addr string) (opentracing.Tracer, io.Closer, error) {
	cfg := jagger.Configuration{
		ServiceName: servicename,
		Sampler: &jagger.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jagger.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}

	sender, err := jaeger.NewUDPTransport(addr, 0)
	if err != nil {
		return nil, nil, err
	}

	reporter := jaeger.NewRemoteReporter(sender)
	// Initialize tracer with a logger and a metrics factory
	tracer, closer, err := cfg.NewTracer(
		jagger.Reporter(reporter),
	)

	return tracer, closer, err
}
