package model

import (
	"context"
	"go.opentelemetry.io/otel/propagation"
	oteltrace "go.opentelemetry.io/otel/trace"
	"server/common"
)

// 需要缓存层的时候再把相关的东西分离出来
// 公共方法 以及 缓存层操作
type defaultProductCustomModel struct {
	cache string
}

// tracer 上报链路
func (d *defaultProductCustomModel) tracer(ctx context.Context, name string) (oteltrace.Span, *propagation.HeaderCarrier) {
	return common.Tracer(ctx, name)
}

func (d *defaultProductCustomModel) GetCache() {

}
