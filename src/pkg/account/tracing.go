/**
 * @Time : 2021/9/17 3:28 PM
 * @Author : solacowa@gmail.com
 * @File : tracing
 * @Software: GoLand
 */

package account

import (
	"context"

	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// 链路追踪中间件
type tracing struct {
	next   Service
	tracer stdopentracing.Tracer
}

func (s *tracing) Namespaces(ctx context.Context, userId, clusterId int64) (res []nsResult, err error) {
	span, ctx := stdopentracing.StartSpanFromContextWithTracer(ctx, s.tracer, "GetNamespaces", stdopentracing.Tag{
		Key:   string(ext.Component),
		Value: "pkg.account",
	})
	defer func() {
		span.LogKV("userId", userId, "clusterId", clusterId, "err", err)
		span.SetTag(string(ext.Error), err != nil)
		span.Finish()
	}()
	return s.next.Namespaces(ctx, userId, clusterId)
}

func (s *tracing) Logout(ctx context.Context, userId int64) (err error) {
	span, ctx := stdopentracing.StartSpanFromContextWithTracer(ctx, s.tracer, "Logout", stdopentracing.Tag{
		Key:   string(ext.Component),
		Value: "package.Account",
	})
	defer func() {
		span.LogKV(
			"userId", userId,
			"err", err)
		span.SetTag(string(ext.Error), err != nil)
		span.Finish()
	}()
	return s.next.Logout(ctx, userId)
}

func (s *tracing) Menus(ctx context.Context, userId int64) (res []userMenuResult, err error) {
	span, ctx := stdopentracing.StartSpanFromContextWithTracer(ctx, s.tracer, "Menus", stdopentracing.Tag{
		Key:   string(ext.Component),
		Value: "package.Account",
	})
	defer func() {
		span.LogKV(
			"userId", userId,
			"err", err)
		span.SetTag(string(ext.Error), err != nil)
		span.Finish()
	}()
	return s.next.Menus(ctx, userId)
}

func (s *tracing) UserInfo(ctx context.Context, userId int64) (res userInfoResult, err error) {
	span, ctx := stdopentracing.StartSpanFromContextWithTracer(ctx, s.tracer, "UserInfo", stdopentracing.Tag{
		Key:   string(ext.Component),
		Value: "package.Account",
	})
	defer func() {
		span.LogKV(
			"userId", userId,
			"err", err)
		span.SetTag(string(ext.Error), err != nil)
		span.Finish()
	}()
	return s.next.UserInfo(ctx, userId)
}

func NewTracing(otTracer stdopentracing.Tracer) Middleware {
	return func(next Service) Service {
		return &tracing{
			next:   next,
			tracer: otTracer,
		}
	}
}
