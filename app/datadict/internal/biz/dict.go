package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/yrcs/yuneto/api/datadict/v1"
	"github.com/yrcs/yuneto/app/datadict/internal/pkg/do"
	"github.com/yrcs/yuneto/app/datadict/internal/pkg/po"
	"github.com/yrcs/yuneto/pkg/repo"
	"github.com/yrcs/yuneto/pkg/usecase"
)

var (
	ErrDictSystemError = errors.InternalServer(v1.ErrorReason_DICT_SYSTEM_ERROR.String(), "系统错误")
)

type E *do.Dict
type T *po.Dict

type DictRepo interface {
	repo.Repo[E, T]
	InitRepo() *repo.BaseRepo[E, T]
	FindChildren(ctx context.Context, id uint64) ([]E, error)
}

type DictUsecase struct {
	usecase.BaseUsecase[E, T]
	dictRepo DictRepo
	tm       Transaction
	log      *log.Helper
}

func NewDictUsecase(dictRepo DictRepo, tm Transaction, logger log.Logger) *DictUsecase {
	return &DictUsecase{
		BaseUsecase: usecase.BaseUsecase[E, T]{Repo: dictRepo.InitRepo()},
		dictRepo:    dictRepo,
		tm:          tm,
		log:         log.NewHelper(log.With(logger, "module", "datadict/biz/DictUsecase")),
	}
}

func (u *DictUsecase) FindChildren(ctx context.Context, id uint64) ([]E, error) {
	return u.dictRepo.FindChildren(ctx, id)
}
