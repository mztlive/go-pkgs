package http

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/mztlive/project-template/pkg/response/httpresponse"
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
		res, err := handler(ctx, req)
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
		res, err := handler(ctx)
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
		err := handler(ctx, req)
		if err != nil {
			httpresponse.SystemError(ctx, err.Error())
			return
		}

		httpresponse.Success(ctx, nil)
	}
}
