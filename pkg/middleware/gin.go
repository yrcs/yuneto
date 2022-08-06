package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	thttp "github.com/go-kratos/kratos/v2/transport/http"
)

func GinMiddlewares(m ...middleware.Middleware) gin.HandlerFunc {
	chain := middleware.Chain(m...)
	Middleware := func(h middleware.Handler) middleware.Handler {
		return chain(h)
	}
	return func(c *gin.Context) {
		if _, exists := c.Get("Middleware"); !exists {
			c.Set("Middleware", Middleware)
		}
		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			c.Next()
			var err error
			if c.Writer.Status() >= 400 {
				err = errors.Errorf(c.Writer.Status(), errors.UnknownReason, errors.UnknownReason)
			}
			return c.Writer, err
		}
		next = chain(next)
		ctx := kgin.NewGinContext(c.Request.Context(), c)
		c.Request = c.Request.WithContext(ctx)
		if ginCtx, ok := kgin.FromGinContext(ctx); ok {
			thttp.SetOperation(ctx, ginCtx.FullPath())
		}
		next(c.Request.Context(), c.Request)
	}
}

func GinMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (reply any, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				log.Info("operation: ", tr.Operation())
			}
			return handler(ctx, req)
		}
	}
}
