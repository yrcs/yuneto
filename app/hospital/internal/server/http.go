package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"

	"github.com/go-kratos/kratos/v2/transport/http"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	v1 "github.com/yrcs/yuneto/api/hospital/v1"
	"github.com/yrcs/yuneto/api/hospital/v1/docs"
	"github.com/yrcs/yuneto/app/hospital/internal/conf"
	"github.com/yrcs/yuneto/app/hospital/internal/service"
	"github.com/yrcs/yuneto/pkg/middleware"
)

// NewHTTPServer new a HTTP server.
// @title       YunETO (云医通) 后台管理系统 API
// @version     1.0
// @description 本文档描述了后台管理系统微服务接口定义。

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
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
		middleware.GinMiddlewares(
			recovery.Recovery(),
			middleware.GinMiddleware(),
			middleware.Server(logger),
			middleware.ProtoValidator(),
		),
	)
	docs.SwaggerInfo.BasePath = "/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	srv := http.NewServer(opts...)
	srv.HandlePrefix("/", r)
	v1.RegisterHospitalHTTPServer(r, hospital)
	return srv
}
