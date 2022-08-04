package server

import (
	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"

	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/yrcs/yuneto/api/hospital/v1"
	"github.com/yrcs/yuneto/app/hospital/internal/conf"
	"github.com/yrcs/yuneto/app/hospital/internal/service"
	mgin "github.com/yrcs/yuneto/pkg/middleware/gin"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, hospital *service.HospitalService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	r := gin.Default()
	r.Use(
		kgin.Middlewares(
			recovery.Recovery(),
			mgin.GinMiddleware(),
			logging.Server(logger),
			mgin.ProtoValidator(),
		),
	)

	srv := http.NewServer(opts...)
	srv.HandlePrefix("/", r)
	v1.RegisterHospitalHTTPServer(r, hospital)
	return srv
}