/**
 * @Time : 8/19/21 1:32 PM
 * @Author : solacowa@gmail.com
 * @File : tracing
 * @Software: GoLand
 */

package secrets

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"

	"github.com/kplcloud/kplcloud/src/repository/types"
)

type tracing struct {
	next   Service
	tracer opentracing.Tracer
}

func (s *tracing) FindNsByNames(ctx context.Context, clusterId int64, ns string, names []string) (res []types.Secret, err error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, s.tracer, "FindNsByNames", opentracing.Tag{
		Key:   string(ext.Component),
		Value: "repository.secrets",
	})
	defer func() {
		span.LogKV("clusterId", clusterId, "ns", ns, "names", names, "err", err)
		span.SetTag(string(ext.Error), err != nil)
		span.Finish()
	}()
	return s.next.FindNsByNames(ctx, clusterId, ns, names)
}

func (s *tracing) FindByName(ctx context.Context, name string) (res []types.Secret, err error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, s.tracer, "FindByName", opentracing.Tag{
		Key:   string(ext.Component),
		Value: "repository.secrets",
	})
	defer func() {
		span.LogKV("name", name, "err", err)
		span.SetTag(string(ext.Error), err != nil)
		span.Finish()
	}()
	return s.next.FindByName(ctx, name)
}

func (s *tracing) Delete(ctx context.Context, clusterId int64, ns, name string) (err error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, s.tracer, "Delete", opentracing.Tag{
		Key:   string(ext.Component),
		Value: "repository.Secrets",
	})
	defer func() {
		span.LogKV(
			"clusterId", clusterId,
			"namespace", ns,
			"name", name,
			"error", err,
		)
		span.Finish()
	}()
	return s.next.Delete(ctx, clusterId, ns, name)
}

func (s *tracing) FindBy(ctx context.Context, clusterId int64, ns, name string) (res types.Secret, err error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, s.tracer, "FindBy", opentracing.Tag{
		Key:   string(ext.Component),
		Value: "repository.Secrets",
	})
	defer func() {
		span.LogKV(
			"clusterId", clusterId,
			"namespace", ns,
			"name", name,
			"error", err,
		)
		span.Finish()
	}()
	return s.next.FindBy(ctx, clusterId, ns, name)
}

func (s *tracing) Save(ctx context.Context, secret *types.Secret, data []types.Data) (err error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, s.tracer, "Save", opentracing.Tag{
		Key:   string(ext.Component),
		Value: "repository.Secrets",
	})
	defer func() {
		span.LogKV(
			"error", err,
		)
		span.Finish()
	}()
	return s.next.Save(ctx, secret, data)
}

func NewTracing(otTracer opentracing.Tracer) Middleware {
	return func(next Service) Service {
		return &tracing{
			next:   next,
			tracer: otTracer,
		}
	}
}
