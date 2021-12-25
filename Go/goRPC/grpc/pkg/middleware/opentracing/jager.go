/*
 * File: jager_client.go
 * Created Date: 2021-12-25 07:12:36
 * Author: ysj
 * Description:
 */

package opentracing

import (
	"grpc-notes/conf"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc/grpclog"
)

var (
	environment        = conf.GetString("mode")
	jaegerCollectorURL = conf.GetString("jaegercollectorurl")
	jaegerAgentURL     = conf.GetString("jaegeragenturl")
)

func JaegerTracer() opentracing.Tracer {
	tracer, _, err := config.Configuration{
		ServiceName: "BROADCAST",
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			CollectorEndpoint:  jaegerCollectorURL,
			LocalAgentHostPort: jaegerAgentURL,
		},
		Tags: []opentracing.Tag{{Key: "environment", Value: environment}},
	}.NewTracer()
	if err != nil {
		grpclog.Errorln(err)
		return nil
	}
	return tracer
}
