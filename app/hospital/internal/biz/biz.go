package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/yrcs/yuneto/app/hospital/internal/pkg/do"
	"github.com/yrcs/yuneto/app/hospital/internal/pkg/po"
	"github.com/yrcs/yuneto/pkg/repo"
	"gorm.io/gorm"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewBaseRepo, NewHospitalSettingUsecase)

func NewBaseRepo(db *gorm.DB, logger log.Logger) *repo.BaseRepo[*do.HospitalSetting, *po.HospitalSetting] {
	return &repo.BaseRepo[*do.HospitalSetting, *po.HospitalSetting]{
		DB:  db,
		Log: log.NewHelper(log.With(logger, "module", "hospital/biz/NewBaseRepo")),
	}
}

type Transaction interface {
	InTx(context.Context, func(ctx context.Context) error) error
}
