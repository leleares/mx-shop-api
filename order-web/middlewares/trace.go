package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func Trace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cfg := jaegercfg.Configuration{
			Sampler: &jaegercfg.SamplerConfig{
				Type:  jaeger.SamplerTypeConst,
				Param: 1,
			},
			Reporter: &jaegercfg.ReporterConfig{
				LogSpans:           true,
				LocalAgentHostPort: "127.0.0.1:6831",
			},
			ServiceName: "mxshop",
		}

		tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
		if err != nil {
			log.Printf("Could not initialize jaeger tracer: %s", err.Error())
			return
		}
		opentracing.SetGlobalTracer(tracer)
		defer closer.Close()

		startSpan := tracer.StartSpan(ctx.Request.URL.Path)
		defer startSpan.Finish()

		ctx.Set("tracer", tracer)
		ctx.Set("parentSpan", startSpan)
		ctx.Next()
	}
}
