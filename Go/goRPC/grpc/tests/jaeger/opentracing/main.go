/*
 * File: opentracing.go
 * Created Date: 2021-12-25 04:38:23
 * Author: ysj
 * Description:  分布式链路追踪
 */

package opentracing

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/grpclog"

	"grpc-notes/conf"
)

const (
	service = "broadcast"
	id      = 1
)

var (
	environment        = conf.GetString("mode")
	jaegerCollectorURL = conf.GetString("jaegercollectorurl")
	jaegerAgentURL     = conf.GetString("jaegeragenturl")
)

func InitTracerProvider() (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	// The HTTP endpoint for sending spans directly to a collector
	// i.e. http://localhost:14268/api/traces
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(jaegerCollectorURL)))
	if err != nil {
		return nil, err
	}

	//tracer provider
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exporter),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
			attribute.String("environment", environment),
		)),
		// tracesdk.WithSampler(),
		// tracesdk.WithIDGenerator(),
		// tracesdk.WithSpanLimits(),
		// tracesdk.WithSpanProcessor(),
	)

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)

	return tp, nil
}

func OpenTracingIntercepter() trace.Tracer {
	tp, err := InitTracerProvider()
	if err != nil {
		grpclog.Errorln(err)
	}
	tr := tp.Tracer("BroadCastTracer")

	return tr
}
