package middleware

import (
	"context"

	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
)

type validator interface {
	Validate() error
}

func ProtoValidator() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if executed := ctx.Value("ProtoValidatorExecutedOnce"); executed == nil {
				ginCtx, _ := kgin.FromGinContext(ctx)
				ginCtx.Set("ProtoValidatorExecutedOnce", true)
				return handler(ctx, req)
			}
			if v, ok := req.(validator); ok {
				if err := v.Validate(); err != nil {
					reason, _ := ctx.Value("reason").(string)
					return nil, errors.BadRequest(reason, "参数错误").WithCause(err)
				}
			}
			return handler(ctx, req)
		}
	}
}
