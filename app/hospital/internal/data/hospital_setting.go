package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/yrcs/yuneto/app/hospital/internal/biz"
	"github.com/yrcs/yuneto/app/hospital/internal/pkg/do"
	"github.com/yrcs/yuneto/app/hospital/internal/pkg/po"
	p "github.com/yrcs/yuneto/pkg/pagination"
	"github.com/yrcs/yuneto/pkg/repo"
	"github.com/yrcs/yuneto/third_party/pagination"
)

var _ biz.HospitalSettingRepo[*do.HospitalSetting, *po.HospitalSetting] = (*hospitalSettingRepo[
	*do.HospitalSetting, *po.HospitalSetting])(nil)

type hospitalSettingRepo[E *do.HospitalSetting, T *po.HospitalSetting] struct {
	repo.BaseRepo[E, T]
	data *Data
	log  *log.Helper
}

func NewHospitalSettingRepo(data *Data, logger log.Logger) biz.HospitalSettingRepo[*do.HospitalSetting, *po.HospitalSetting] {
	return &hospitalSettingRepo[*do.HospitalSetting, *po.HospitalSetting]{
		BaseRepo: repo.BaseRepo[*do.HospitalSetting, *po.HospitalSetting]{DB: data.db},
		data:     data,
		log:      log.NewHelper(log.With(logger, "module", "hospital/data/HospitalSettingRepo")),
	}
}

func (r *hospitalSettingRepo[E, T]) List(ctx context.Context, req *pagination.PagingRequest) ([]E, error) {
	db := r.data.db.WithContext(ctx)

	var hospitalSettings []E
	result := db.Scopes(p.Paginate(req)).Find(&hospitalSettings)
	if err := result.Error; err != nil {
		return nil, err
	}
	return hospitalSettings, nil
}
