package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"github.com/yrcs/yuneto/app/datadict/internal/biz"
	"github.com/yrcs/yuneto/pkg/repo"
)

var (
	br *repo.BaseRepo[biz.E, biz.T]
	_  biz.DictRepo = (*dictRepo)(nil)
)

type dictRepo struct {
	*repo.BaseRepo[biz.E, biz.T]
	data *Data
	log  *log.Helper
}

func NewDictRepo(data *Data, logger log.Logger) biz.DictRepo {
	br = &repo.BaseRepo[biz.E, biz.T]{
		DB:  data.db,
		Log: log.NewHelper(log.With(logger, "module", "yuneto/pkg/repo/BaseRepo")),
	}
	return &dictRepo{
		BaseRepo: br,
		data:     data,
		log:      log.NewHelper(log.With(logger, "module", "datadict/data/DictRepo")),
	}
}

func (r *dictRepo) InitRepo() *repo.BaseRepo[biz.E, biz.T] {
	return br
}

func (r *dictRepo) FindChildren(ctx context.Context, id uint64) ([]biz.E, error) {
	db := r.data.db.WithContext(ctx)

	var pos []biz.T
	result := db.Where("ParentId = ?", id).Find(&pos)
	if err := result.Error; err != nil {
		return nil, err
	}
	for _, v := range pos {
		if hasChildren, err := r.Exists(ctx, "ParentId", id); err == nil {
			v.HasChildren = hasChildren
		} else {
			return nil, err
		}
	}
	var dos []biz.E
	copier.CopyWithOption(&dos, pos, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	return dos, nil
}
