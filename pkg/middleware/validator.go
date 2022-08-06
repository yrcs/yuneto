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
		return func(ctx context.Context, req any) (reply any, err error) {
			ginCtx, _ := kgin.FromGinContext(ctx)
			validate := func(handler middleware.Handler) middleware.Handler {
				return func(ctx context.Context, req any) (reply any, err error) {
					if v, ok := req.(validator); ok {
						if err := v.Validate(); err != nil {
							reason := ginCtx.MustGet("reason").(string)
							return nil, errors.BadRequest(reason, "参数错误").WithCause(err)
						}
					}
					return handler(ctx, req)
				}
			}
			ginCtx.Set("validate", validate)
			return
		}
	}
}
