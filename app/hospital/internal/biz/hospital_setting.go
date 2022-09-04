package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/yrcs/yuneto/api/hospital/v1"
	"github.com/yrcs/yuneto/app/hospital/internal/pkg/do"
	"github.com/yrcs/yuneto/app/hospital/internal/pkg/po"
	"github.com/yrcs/yuneto/pkg/repo"
	"github.com/yrcs/yuneto/pkg/usecase"
	"github.com/yrcs/yuneto/third_party/pagination"
)

var (
	ErrHospitalSettingSameDataExists = errors.BadRequest(v1.ErrorReason_HOSPITAL_SETTING_SAME_DATA_EXISTS.String(), "医院名称或登记号已存在")
	ErrHospitalSettingNotFound       = errors.NotFound(v1.ErrorReason_HOSPITAL_SETTING_NOT_FOUND.String(), "医院设置不存在或已删除")
	ErrHospitalSettingSystemError    = errors.InternalServer(v1.ErrorReason_HOSPITAL_SETTING_SYSTEM_ERROR.String(), "系统错误")
)

type E *do.HospitalSetting
type T *po.HospitalSetting

type HospitalSettingRepo interface {
	repo.Repo[E, T]
	InitRepo() *repo.BaseRepo[E, T]
	List(context.Context, *pagination.PagingRequest) ([]E, error)
}

type HospitalSettingUsecase struct {
	usecase.BaseUsecase[E, T]
	hsRepo HospitalSettingRepo
	tm     Transaction
	log    *log.Helper
}

func NewHospitalSettingUsecase(hsRepo HospitalSettingRepo, tm Transaction, logger log.Logger) *HospitalSettingUsecase {
	return &HospitalSettingUsecase{
		hsRepo:      hsRepo,
		BaseUsecase: usecase.BaseUsecase[E, T]{Repo: hsRepo.InitRepo()},
		tm:          tm,
		log:         log.NewHelper(log.With(logger, "module", "hospital/biz/HospitalSettingUsecase")),
	}
}

func (hsu *HospitalSettingUsecase) NameExists(ctx context.Context, name string) (bool, error) {
	return hsu.hsRepo.Exists(ctx, "Name", name)
}

func (hsu *HospitalSettingUsecase) RegistrationNumberExists(ctx context.Context, registrationNumber string) (bool, error) {
	return hsu.hsRepo.Exists(ctx, "RegistrationNumber", registrationNumber, true)
}

func (hsu *HospitalSettingUsecase) NameUnique(ctx context.Context, id uint64, name string) (bool, error) {
	return hsu.hsRepo.Unique(ctx, id, "Name", name)
}

func (hsu *HospitalSettingUsecase) RegistrationNumberUnique(ctx context.Context, id uint64, registrationNumber string) (bool, error) {
	return hsu.hsRepo.Unique(ctx, id, "RegistrationNumber", registrationNumber, true)
}

func (hsu *HospitalSettingUsecase) List(ctx context.Context, req *pagination.PagingRequest) ([]E, error) {
	return hsu.hsRepo.List(ctx, req)
}
