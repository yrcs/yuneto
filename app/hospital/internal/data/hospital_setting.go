package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/yrcs/yuneto/app/hospital/internal/biz"
	p "github.com/yrcs/yuneto/pkg/pagination"
	"github.com/yrcs/yuneto/pkg/repo"
	"github.com/yrcs/yuneto/third_party/pagination"
)

var _ biz.HospitalSettingRepo[biz.E, biz.T] = (*hospitalSettingRepo[biz.E, biz.T])(nil)

type hospitalSettingRepo[E biz.E, T biz.T] struct {
	repo.BaseRepo[E, T]
	data *Data
	log  *log.Helper
}

func NewHospitalSettingRepo(data *Data, logger log.Logger) biz.HospitalSettingRepo[biz.E, biz.T] {
	return &hospitalSettingRepo[biz.E, biz.T]{
		BaseRepo: repo.BaseRepo[biz.E, biz.T]{DB: data.db},
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
