package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/yrcs/yuneto/api/hospital/v1"
)

var (
	ErrHospitalSettingSameDataExists = errors.BadRequest(v1.ErrorReason_HOSPITAL_SETTING_SAME_DATA_EXISTS.String(), "医院名称或登记号已存在")
	ErrHospitalSettingNotFound       = errors.NotFound(v1.ErrorReason_HOSPITAL_SETTING_NOT_FOUND.String(), "医院设置不存在或已删除")
	ErrHospitalSettingSystemError    = errors.InternalServer(v1.ErrorReason_HOSPITAL_SETTING_SYSTEM_ERROR.String(), "系统错误")
)

type HospitalSetting struct {
	Id                 uint64
	CreatedAt          time.Time
	UpdatedAt          time.Time
	Name               string
	RegistrationNumber string
	ContactPerson      string
	ContactMobile      string
	Locked             uint8
	ApiUrl             string
	SignatureKey       string
}

type HospitalSettingRepo interface {
	Add(context.Context, *HospitalSetting) (*HospitalSetting, error)
	Edit(context.Context, *HospitalSetting) (*HospitalSetting, error)
	Delete(context.Context, *HospitalSetting) error
	DeleteInBatches(context.Context, []uint64) error
}

type HospitalSettingUsecase struct {
	hsRepo HospitalSettingRepo
	log    *log.Helper
}

func NewHospitalSettingUsecase(hsRepo HospitalSettingRepo, logger log.Logger) *HospitalSettingUsecase {
	return &HospitalSettingUsecase{
		hsRepo: hsRepo,
		log:    log.NewHelper(log.With(logger, "module", "hospital/biz/HospitalSettingUsecase")),
	}
}

func (hsu *HospitalSettingUsecase) Add(ctx context.Context, hs *HospitalSetting) (*HospitalSetting, error) {
	return hsu.hsRepo.Add(ctx, hs)
}

func (hsu *HospitalSettingUsecase) Edit(ctx context.Context, hs *HospitalSetting) (*HospitalSetting, error) {
	return hsu.hsRepo.Edit(ctx, hs)
}

func (hsu *HospitalSettingUsecase) Delete(ctx context.Context, hs *HospitalSetting) error {
	return hsu.hsRepo.Delete(ctx, hs)
}

func (hsu *HospitalSettingUsecase) DeleteInBatches(ctx context.Context, ids []uint64) error {
	return hsu.hsRepo.DeleteInBatches(ctx, ids)
}
