package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "github.com/yrcs/yuneto/api/hospital/v1"
	"github.com/yrcs/yuneto/app/hospital/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewHospitalService)

type HospitalService struct {
	v1.UnimplementedHospitalServer

	hsu *biz.HospitalSettingUsecase[biz.E, biz.T]
	log *log.Helper
}

func NewHospitalService(hsu *biz.HospitalSettingUsecase[biz.E, biz.T], logger log.Logger) *HospitalService {
	return &HospitalService{
		hsu: hsu,
		log: log.NewHelper(log.With(logger, "module", "hospital/service")),
	}
}
