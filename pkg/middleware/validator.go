package middleware

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
)

type validator interface {
	Validate() error
}

func ProtoValidator() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if v, ok := req.(validator); ok {
				if err := v.Validate(); err != nil {
					reason, _ := ctx.Value("reason").(string)
					return nil, errors.BadRequest(reason, err.Error()).WithCause(err)
				}
			}
			return handler(ctx, req)
		}
	}
}
