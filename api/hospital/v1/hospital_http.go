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
}

func RegisterHospitalHTTPServer(r *gin.Engine, srv HospitalHTTPServer) {
	g := r.Group("/v1/admin/hospital")
	{
		// ?page=1&pageSize=10&query[name]=中山大学孙逸仙纪念医院&query[contact_person]=徐瑞华&orderBy[name]=1&orderBy[id]=2
		g.GET("/hospitalSettings", ListHospitalSettingHandler(srv))
		g.POST("/hospitalSettings", AddHospitalSettingHandler(srv))
		g.PUT("/hospitalSettings/:id", EditHospitalSettingHandler(srv))
		g.DELETE("/hospitalSettings/:id", DeleteHospitalSettingHandler(srv))
		g.DELETE("/hospitalSettings", DeleteHospitalSettingsHandler(srv))
	}
}

func ListHospitalSettingHandler(srv HospitalHTTPServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		o := make(map[string]pagination.Order, 0)
		in := pagination.PagingRequest{OrderBy: o}
		p.PackPageData(ctx, &in)

		http.SetOperation(ctx, "/hospital.v1.Hospital/ListHospitalSetting")
		ctx.Set("reason", ErrorReason_HOSPITAL_SETTING_INVALID_ARGUMENT.String())
		validate := ctx.MustGet("validate").(func(handler middleware.Handler) middleware.Handler)
		h := validate(func(ctx context.Context, req any) (any, error) {
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

func AddHospitalSettingHandler(srv HospitalHTTPServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var in AddHospitalSettingRequest
		if err := ctx.ShouldBindJSON(&in); err != nil {
			result.Error(ctx, errors.BadRequest(ErrorReason_HOSPITAL_SETTING_INVALID_DATA_FORMAT.String(), "表单数据格式错误").WithCause(err))
			return
		}

		http.SetOperation(ctx, "/hospital.v1.Hospital/AddHospitalSetting")
		ctx.Set("reason", ErrorReason_HOSPITAL_SETTING_INVALID_ARGUMENT.String())
		validate := ctx.MustGet("validate").(func(handler middleware.Handler) middleware.Handler)
		h := validate(func(ctx context.Context, req any) (any, error) {
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
		validate := ctx.MustGet("validate").(func(handler middleware.Handler) middleware.Handler)
		h := validate(func(ctx context.Context, req any) (any, error) {
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

func DeleteHospitalSettingHandler(srv HospitalHTTPServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		in := DeleteHospitalSettingRequest{Id: uint64(id)}
		http.SetOperation(ctx, "/hospital.v1.Hospital/DeleteHospitalSetting")
		ctx.Set("reason", ErrorReason_HOSPITAL_SETTING_INVALID_ARGUMENT.String())
		validate := ctx.MustGet("validate").(func(handler middleware.Handler) middleware.Handler)
		h := validate(func(ctx context.Context, req any) (any, error) {
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

func DeleteHospitalSettingsHandler(srv HospitalHTTPServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var in DeleteHospitalSettingsRequest
		if err := ctx.ShouldBindJSON(&in); err != nil {
			result.Error(ctx, errors.BadRequest(ErrorReason_HOSPITAL_SETTING_INVALID_DATA_FORMAT.String(), "表单数据格式错误").WithCause(err))
			return
		}

		http.SetOperation(ctx, "/hospital.v1.Hospital/DeleteHospitalSettings")
		ctx.Set("reason", ErrorReason_HOSPITAL_SETTING_INVALID_ARGUMENT.String())
		validate := ctx.MustGet("validate").(func(handler middleware.Handler) middleware.Handler)
		h := validate(func(ctx context.Context, req any) (any, error) {
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
