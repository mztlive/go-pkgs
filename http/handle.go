package http

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/mztlive/go-pkgs/response/httpresponse"
	"go.opentelemetry.io/otel"
)

type ServiceAnyHandlerWithArg[R, S any] func(ctx context.Context, req R) (S, error)

type ServiceAnyHandler[S any] func(ctx context.Context) (S, error)

type ServicePostHandlerNotResp[R any] func(ctx context.Context, req R) error

type ServicePostHandlerNotRespWithArg[R any] func(ctx context.Context, req R) error

// PostHandle is a gin handler wrapper for service handler
func PostHandle[R, S any](handler ServiceAnyHandlerWithArg[R, S]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req R
		if err := ctx.ShouldBind(&req); err != nil {
			httpresponse.BadRequest(ctx, err.Error())
			return
		}

		// 获取一个tracer
		tracer := otel.Tracer(ctx.Request.URL.Path)
		// 开始一个新的Span
		newCtx, span := tracer.Start(ctx, "service_handle")
		defer span.End()

		res, err := handler(newCtx, req)
		if err != nil {
			httpresponse.SystemError(ctx, err.Error())
			return
		}

		httpresponse.Success(ctx, res)
	}
}

// AnyHandle is a gin handler wrapper for service handler
func AnyHandle[S any](handler ServiceAnyHandler[S]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取一个tracer
		tracer := otel.Tracer(ctx.Request.URL.Path)
		// 开始一个新的Span
		newCtx, span := tracer.Start(ctx, "service_handle")

		defer span.End()
		res, err := handler(newCtx)
		if err != nil {
			httpresponse.SystemError(ctx, err.Error())
			return
		}

		httpresponse.Success(ctx, res)
	}
}

// PostHandleNotResp is a gin handler wrapper for service handler
func PostHandleNotResp[R any](handler ServicePostHandlerNotResp[R]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req R
		if err := ctx.ShouldBind(&req); err != nil {
			httpresponse.BadRequest(ctx, err.Error())
			return
		}

		// 获取一个tracer
		tracer := otel.Tracer(ctx.Request.URL.Path)
		// 开始一个新的Span
		newCtx, span := tracer.Start(ctx, "service_handle")

		defer span.End()

		err := handler(newCtx, req)
		if err != nil {
			httpresponse.SystemError(ctx, err.Error())
			return
		}

		httpresponse.Success(ctx, nil)
	}
}
