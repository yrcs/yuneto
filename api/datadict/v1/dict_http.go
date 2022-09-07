package v1

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/kratos/v2/transport/http/binding"
	"github.com/yrcs/yuneto/pkg/result"
)

// This is a compile-time assertion to ensure that this file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type DataDictHTTPServer interface {
	ListChildren(context.Context, *ListChildrenRequest) (*ListChildrenReply, error)
}

func RegisterDataDictHTTPServer(r *gin.Engine, srv DataDictHTTPServer) {
	v1 := r.Group("/v1")
	{
		g1 := v1.Group("/admin/datadict/dicts")
		{
			g1.GET("/:id", ListChildrenHandler(srv))
		}
	}
}

// @Tags        获取子级列表
// @Summary     子级列表
// @Description 获取指定父级 ID 的子级列表
// @Accept      json
// @Produce     json
// @Param       id  path    int true "父级 id" Format(uint64)
// @Success     200 {array} ListChildrenReply
// @Router      /admin/datadict/dicts/{id} [get]
func ListChildrenHandler(srv DataDictHTTPServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		in := ListChildrenRequest{}
		in.Id = uint64(id)

		http.SetOperation(ctx, "/datadict.v1.DataDict/ListChildren")
		ctx.Set("reason", ErrorReason_DICT_INVALID_ARGUMENT.String())
		Middleware := ctx.MustGet("Middleware").(func(handler middleware.Handler) middleware.Handler)
		h := Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ListChildren(ctx, req.(*ListChildrenRequest))
		})

		out, err := h(ctx, &in)
		if err != nil {
			result.Error(ctx, err)
			return
		}
		result.Result(ctx, out)
	}
}
