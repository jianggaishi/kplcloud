/**
 * @Time : 8/11/21 4:21 PM
 * @Author : solacowa@gmail.com
 * @File : logging
 * @Software: GoLand
 */

package registry

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type logging struct {
	logger  log.Logger
	next    Service
	traceId string
}

func (s *logging) Secret(ctx context.Context, name string) (err error) {
	panic("implement me")
}

func (s *logging) Update(ctx context.Context, name, host, username, password, remark string) (err error) {
	defer func(begin time.Time) {
		_ = s.logger.Log(
			s.traceId, ctx.Value(s.traceId),
			"method", "Update", "name", name, "host", host, "username", username, "password", password, "remark", remark,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.next.Update(ctx, name, host, username, password, remark)
}

func (s *logging) Delete(ctx context.Context, name string) (err error) {
	defer func(begin time.Time) {
		_ = s.logger.Log(
			s.traceId, ctx.Value(s.traceId),
			"method", "Delete", "name", name,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.next.Delete(ctx, name)
}

func (s *logging) Password(ctx context.Context, name string) (res string, err error) {
	panic("implement me")
}

func (s *logging) Info(ctx context.Context, name string) (res result, err error) {
	panic("implement me")
}

func (s *logging) List(ctx context.Context, query string, page, pageSize int) (res []result, total int, err error) {
	defer func(begin time.Time) {
		_ = s.logger.Log(
			s.traceId, ctx.Value(s.traceId),
			"query", query,
			"page", page,
			"pageSize", pageSize,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.next.List(ctx, query, page, pageSize)
}

func (s *logging) Create(ctx context.Context, name, host, username, password, remark string) (err error) {
	defer func(begin time.Time) {
		_ = s.logger.Log(
			s.traceId, ctx.Value(s.traceId),
			"name", name,
			"host", host,
			"username", username,
			"password", password,
			"remark", remark,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.next.Create(ctx, name, host, username, password, remark)
}

func NewLogging(logger log.Logger, traceId string) Middleware {
	logger = log.With(logger, "registry", "logging")
	return func(next Service) Service {
		return &logging{
			logger:  level.Info(logger),
			next:    next,
			traceId: traceId,
		}
	}
}
