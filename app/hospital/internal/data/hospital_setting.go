package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/yrcs/yuneto/app/hospital/internal/biz"
	p "github.com/yrcs/yuneto/pkg/pagination"
	"github.com/yrcs/yuneto/pkg/repo"
	"github.com/yrcs/yuneto/third_party/pagination"
)

var (
	br *repo.BaseRepo[biz.E, biz.T]
	_  biz.HospitalSettingRepo = (*hospitalSettingRepo)(nil)
)

type hospitalSettingRepo struct {
	*repo.BaseRepo[biz.E, biz.T]
	data *Data
	log  *log.Helper
}

func NewHospitalSettingRepo(data *Data, logger log.Logger) biz.HospitalSettingRepo {
	br = &repo.BaseRepo[biz.E, biz.T]{
		DB:  data.db,
		Log: log.NewHelper(log.With(logger, "module", "yuneto/pkg/repo/BaseRepo")),
	}
	return &hospitalSettingRepo{
		BaseRepo: br,
		data:     data,
		log:      log.NewHelper(log.With(logger, "module", "hospital/data/HospitalSettingRepo")),
	}
}

func (r *hospitalSettingRepo) InitRepo() *repo.BaseRepo[biz.E, biz.T] {
	return br
}

func (r *hospitalSettingRepo) List(ctx context.Context, req *pagination.PagingRequest) ([]biz.E, error) {
	db := r.data.db.WithContext(ctx)

	var hospitalSettings []biz.E
	result := db.Scopes(p.Paginate(req)).Find(&hospitalSettings)
	if err := result.Error; err != nil {
		return nil, err
	}
	return hospitalSettings, nil
}
