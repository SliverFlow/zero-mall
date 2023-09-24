package common

import (
	"context"
	"github.com/zeromicro/go-zero/core/trace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	oteltrace "go.opentelemetry.io/otel/trace"
)

// Tracer
// Author [SliverFlow]
// @desc 上报链路
func Tracer(ctx context.Context, name string) (oteltrace.Span, *propagation.HeaderCarrier) {
	tracer := otel.GetTracerProvider().Tracer(trace.TraceName)
	spanCtx, span := tracer.Start(ctx, name, oteltrace.WithSpanKind(oteltrace.SpanKindProducer))
	carrier := &propagation.HeaderCarrier{}
	otel.GetTextMapPropagator().Inject(spanCtx, carrier)
	return span, carrier
}
