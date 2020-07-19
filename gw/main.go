package main

import (
	"github.com/davveo/tsquare/lib/token"
	"github.com/davveo/tsquare/lib/wrapper/auth"
	"github.com/micro/cli/v2"
	"github.com/micro/go-plugins/micro/cors/v2"
	"github.com/micro/micro/v2/cmd"
	"github.com/micro/micro/v2/plugin"
)

func init() {
	token := &token.Token{}

	_ = plugin.Register(cors.NewPlugin())

	_ = plugin.Register(plugin.NewPlugin(
		plugin.WithName("auth"),
		plugin.WithHandler(
			auth.JWTAuthWrapper(token),
		),
		plugin.WithInit(func(context *cli.Context) error {
			token.InitConfig(context.String("consul_address"), "micro", "config", "jwt-key", "key")
			return nil
		}),
	))
	//plugin.Register(plugin.NewPlugin(
	//	plugin.WithName("tracer"),
	//	plugin.WithHandler(
	//		stdhttp.TracerWrapper,
	//	),
	//))
	//plugin.Register(plugin.NewPlugin(
	//	plugin.WithName("breaker"),
	//	plugin.WithHandler(
	//		hystrix.BreakerWrapper,
	//	),
	//))
	//plugin.Register(plugin.NewPlugin(
	//	plugin.WithName("metrics"),
	//	plugin.WithHandler(
	//		prometheus.MetricsWrapper,
	//	),
	//))
}

const name = "API gateway"

func main() {
	//stdhttp.SetSamplingFrequency(50)
	//t, io, err := tracer.NewTracer(name, "")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer io.Close()
	//opentracing.SetGlobalTracer(t)
	//
	//hystrixStreamHandler := ph.NewStreamHandler()
	//hystrixStreamHandler.Start()
	//go http.ListenAndServe(net.JoinHostPort("", "81"), hystrixStreamHandler)

	cmd.Init()
}
