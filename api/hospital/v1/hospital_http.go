package v1

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/kratos/v2/transport/http/binding"
	"google.golang.org/protobuf/types/known/emptypb"

	p "github.com/yrcs/yuneto/pkg/pagination"
	"github.com/yrcs/yuneto/pkg/result"
	"github.com/yrcs/yuneto/third_party/pagination"
)

// This is a compile-time assertion to ensure that this file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type HospitalHTTPServer interface {
	ListHospitalSetting(context.Context, *pagination.PagingRequest) (*pagination.PagingReply, error)
	AddHospitalSetting(context.Context, *AddHospitalSettingRequest) (*CommonAddReply, error)
	EditHospitalSetting(context.Context, *EditHospitalSettingRequest) (*CommonEditReply, error)
	DeleteHospitalSetting(context.Context, *DeleteHospitalSettingRequest) (*emptypb.Empty, error)
	DeleteHospitalSettings(context.Context, *DeleteHospitalSettingsRequest) (*emptypb.Empty, error)
	LockHospitalSetting(context.Context, *LockHospitalSettingRequest) (*CommonEditReply, error)
}

func RegisterHospitalHTTPServer(r *gin.Engine, srv HospitalHTTPServer) {
	v1 := r.Group("/v1")
	{
		g1 := v1.Group("/admin/hospital/hospitalSettings")
		{
			// ?page=1&pageSize=10&query[name]=中山大学附属第一医院&orderBy[name]=1&orderBy[id]=2
			g1.GET("", ListHospitalSettingHandler(srv))
			g1.POST("", AddHospitalSettingHandler(srv))
			g1.PUT("/:id", EditHospitalSettingHandler(srv))
			g1.DELETE("/:id", DeleteHospitalSettingHandler(srv))
			g1.DELETE("", DeleteHospitalSettingsHandler(srv))
			g1.PUT("/:id/:locked", LockHospitalSettingHandler(srv))
		}
	}
}

// @Tags        获取医院设置列表
// @Summary     医院设置列表
// @Description 医院设置分页列表
// @Accept      json
// @Produce     json
// @Param       page          query    int    false "页码"    Format(uint32)
// @Param       pageSize      query    int    false "每页条目数" Format(uint32)
// @Param       query[name]   query    string false "查询参数"
// @Param       orderBy[name] query    int    false "排序参数" Enums(0, 1, 2)
// @Success     200           {object} pagination.PagingReply
// @Router      /admin/hospital/hospitalSettings [get]
func ListHospitalSettingHandler(srv HospitalHTTPServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		o := make(map[string]pagination.Order, 0)
		in := pagination.PagingRequest{OrderBy: o}
		p.PackPagingData(ctx, &in)

		http.SetOperation(ctx, "/hospital.v1.Hospital/ListHospitalSetting")
		ctx.Set("reason", ErrorReason_HOSPITAL_SETTING_INVALID_ARGUMENT.String())
		Middleware := ctx.MustGet("Middleware").(func(handler middleware.Handler) middleware.Handler)
		h := Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ListHospitalSetting(ctx, req.(*pagination.PagingRequest))
		})

		out, err := h(ctx, &in)
		if err != nil {
			result.Error(ctx, err)
			return
		}
		result.Result(ctx, out)
	}
}

// @Tags        添加医院设置
// @Summary     添加设置
// @Description 接收 json 以添加医院设置
// @Accept      json
// @Produce     json
// @Param       message body     AddHospitalSettingRequest true "医院设置属性"
// @Success     200     {object} CommonAddReply
// @Router      /admin/hospital/hospitalSettings [post]
func AddHospitalSettingHandler(srv HospitalHTTPServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var in AddHospitalSettingRequest
		if err := ctx.ShouldBindJSON(&in); err != nil {
			result.Error(ctx, errors.BadRequest(ErrorReason_HOSPITAL_SETTING_INVALID_DATA_FORMAT.String(), "表单数据格式错误").WithCause(err))
			return
		}

		http.SetOperation(ctx, "/hospital.v1.Hospital/AddHospitalSetting")
		ctx.Set("reason", ErrorReason_HOSPITAL_SETTING_INVALID_ARGUMENT.String())
		Middleware := ctx.MustGet("Middleware").(func(handler middleware.Handler) middleware.Handler)
		h := Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.AddHospitalSetting(ctx, req.(*AddHospitalSettingRequest))
		})

		out, err := h(ctx, &in)
		if err != nil {
			result.Error(ctx, err)
			return
		}
		result.Result(ctx, out)
	}
}

// @Tags        编辑医院设置
// @Summary     编辑设置
// @Description 接收 id 与 json 以编辑医院设置
// @Accept      json
// @Produce     json
// @Param       id      path     int                        true  "医院设置 id" Format(uint64)
// @Param       message body     EditHospitalSettingRequest false "医院设置属性"
// @Success     200    {object} CommonEditReply
// @Router      /admin/hospital/hospitalSettings/{id} [put]
func EditHospitalSettingHandler(srv HospitalHTTPServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		in := EditHospitalSettingRequest{Id: uint64(id)}
		if err := ctx.ShouldBindJSON(&in); err != nil {
			result.Error(ctx, errors.BadRequest(ErrorReason_HOSPITAL_SETTING_INVALID_DATA_FORMAT.String(), "表单数据格式错误").WithCause(err))
			return
		}

		http.SetOperation(ctx, "/hospital.v1.Hospital/EditHospitalSetting")
		ctx.Set("reason", ErrorReason_HOSPITAL_SETTING_INVALID_ARGUMENT.String())
		Middleware := ctx.MustGet("Middleware").(func(handler middleware.Handler) middleware.Handler)
		h := Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.EditHospitalSetting(ctx, req.(*EditHospitalSettingRequest))
		})

		out, err := h(ctx, &in)
		if err != nil {
			result.Error(ctx, err)
			return
		}
		result.Result(ctx, out)
	}
}

// @Tags        删除医院设置
// @Summary     删除设置
// @Description 接收 id 以删除医院设置
// @Accept      json
// @Produce     json
// @Param       id path int true "医院设置 id" Format(uint64)
// @Success     200
// @Router      /admin/hospital/hospitalSettings/{id} [delete]
func DeleteHospitalSettingHandler(srv HospitalHTTPServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		in := DeleteHospitalSettingRequest{Id: uint64(id)}
		http.SetOperation(ctx, "/hospital.v1.Hospital/DeleteHospitalSetting")
		ctx.Set("reason", ErrorReason_HOSPITAL_SETTING_INVALID_ARGUMENT.String())
		Middleware := ctx.MustGet("Middleware").(func(handler middleware.Handler) middleware.Handler)
		h := Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.DeleteHospitalSetting(ctx, req.(*DeleteHospitalSettingRequest))
		})

		out, err := h(ctx, &in)
		if err != nil {
			result.Error(ctx, err)
			return
		}
		result.Result(ctx, out)
	}
}

// @Tags        批量删除医院设置
// @Summary     批量删除设置
// @Description 接收 json 以批量删除医院设置
// @Accept      json
// @Produce     json
// @Param       message body DeleteHospitalSettingsRequest true "医院设置 id 数组"
// @Success     200
// @Router      /admin/hospital/hospitalSettings [delete]
func DeleteHospitalSettingsHandler(srv HospitalHTTPServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var in DeleteHospitalSettingsRequest
		if err := ctx.ShouldBindJSON(&in); err != nil {
			result.Error(ctx, errors.BadRequest(ErrorReason_HOSPITAL_SETTING_INVALID_DATA_FORMAT.String(), "表单数据格式错误").WithCause(err))
			return
		}

		http.SetOperation(ctx, "/hospital.v1.Hospital/DeleteHospitalSettings")
		ctx.Set("reason", ErrorReason_HOSPITAL_SETTING_INVALID_ARGUMENT.String())
		Middleware := ctx.MustGet("Middleware").(func(handler middleware.Handler) middleware.Handler)
		h := Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.DeleteHospitalSettings(ctx, req.(*DeleteHospitalSettingsRequest))
		})

		out, err := h(ctx, &in)
		if err != nil {
			result.Error(ctx, err)
			return
		}
		result.Result(ctx, out)
	}
}

// @Tags        锁定/解锁医院设置
// @Summary     锁定/解锁设置
// @Description 接收 id 与 locked 以锁定/解锁医院设置
// @Accept      json
// @Produce     json
// @Param       id     path     int true "锁定/解锁医院设置" Format(uint64)
// @Param       locked path     int true "锁定/解锁状态"   Enums(0, 1)
// @Success     200     {object} CommonEditReply
// @Router      /admin/hospital/hospitalSettings/{id}/{locked} [put]
func LockHospitalSettingHandler(srv HospitalHTTPServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		locked, _ := strconv.Atoi(ctx.Param("locked"))
		in := LockHospitalSettingRequest{Id: uint64(id), Locked: uint32(locked)}
		http.SetOperation(ctx, "/hospital.v1.Hospital/LockHospitalSetting")
		ctx.Set("reason", ErrorReason_HOSPITAL_SETTING_INVALID_ARGUMENT.String())
		Middleware := ctx.MustGet("Middleware").(func(handler middleware.Handler) middleware.Handler)
		h := Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.LockHospitalSetting(ctx, req.(*LockHospitalSettingRequest))
		})

		out, err := h(ctx, &in)
		if err != nil {
			result.Error(ctx, err)
			return
		}
		result.Result(ctx, out)
	}
}

// type GreeterHTTPClient interface {
// 	SayHello(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
// }

// type GreeterHTTPClientImpl struct {
// 	cc *http.Client
// }

// func NewGreeterHTTPClient(client *http.Client) GreeterHTTPClient {
// 	return &GreeterHTTPClientImpl{client}
// }

// func (c *GreeterHTTPClientImpl) SayHello(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
// 	var out HelloReply
// 	pattern := "/helloworld/{name}"
// 	path := binding.EncodeURL(pattern, in, true)
// 	opts = append(opts, http.Operation("/hospital.v1.Hospital/ListHospitalSetting"))
// 	opts = append(opts, http.PathTemplate(pattern))
// 	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &out, err
// }
