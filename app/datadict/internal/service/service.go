package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	v1 "github.com/yrcs/yuneto/api/datadict/v1"
	"github.com/yrcs/yuneto/app/datadict/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewDataDictService)

type DataDictService struct {
	v1.UnimplementedDataDictServer

	u   *biz.DictUsecase
	log *log.Helper
}

func NewDataDictService(u *biz.DictUsecase, logger log.Logger) *DataDictService {
	return &DataDictService{
		u:   u,
		log: log.NewHelper(log.With(logger, "module", "datadict/service")),
	}
}
