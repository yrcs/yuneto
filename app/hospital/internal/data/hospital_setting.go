package data

import (
	"context"
	"errors"

	"github.com/bwmarrin/snowflake"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/yrcs/yuneto/app/hospital/internal/biz"
	"github.com/yrcs/yuneto/pkg/orm/gorm/po"
	"gorm.io/gorm"
)

var _ biz.HospitalSettingRepo = (*hospitalSettingRepo)(nil)

type hospitalSettingRepo struct {
	data *Data
	log  *log.Helper
}

type HospitalSetting struct {
	po.Base
	Name               string `gorm:"unique;comment:医院名称"`
	RegistrationNumber string `gorm:"unique;comment:医院登记号"`
	ContactPerson      string `gorm:"comment:联系人"`
	ContactMobile      string `gorm:"comment:联系人手机"`
	Locked             uint8  `gorm:"default:0;comment:锁定状态（0:未锁定，1:已锁定）"`
	ApiUrl             string `gorm:"comment:API 基础路径"`
	SignatureKey       string `gorm:"comment:签名密钥"`
}

func NewHospitalSettingRepo(data *Data, logger log.Logger) biz.HospitalSettingRepo {
	return &hospitalSettingRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "hospital/data/HospitalSettingRepo")),
	}
}

func (r *hospitalSettingRepo) Add(ctx context.Context, hs *biz.HospitalSetting) (*biz.HospitalSetting, error) {
	db := r.data.db.WithContext(ctx)

	var hospitalSettings []HospitalSetting
	result := db.Where("name = ?", hs.Name).Or("registration_number = ?", hs.RegistrationNumber).Find(&hospitalSettings)
	if err := result.Error; err == nil {
		if result.RowsAffected > 0 {
			return nil, biz.ErrHospitalSettingSameDataExists
		}
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, biz.ErrHospitalSettingSystemError.WithCause(err)
	}

	var hospitalSetting HospitalSetting
	node, _ := snowflake.NewNode(1)
	hospitalSetting = HospitalSetting{
		Base:               po.Base{Id: uint64(node.Generate().Int64())},
		Name:               hs.Name,
		RegistrationNumber: hs.RegistrationNumber,
		ContactPerson:      hs.ContactPerson,
		ContactMobile:      hs.ContactMobile,
		Locked:             hs.Locked,
		ApiUrl:             hs.ApiUrl,
		SignatureKey:       hs.SignatureKey,
	}
	result = db.Create(&hospitalSetting)
	if err := result.Error; err != nil {
		return nil, biz.ErrHospitalSettingSystemError.WithCause(err)
	}

	return &biz.HospitalSetting{
		Id:        hospitalSetting.Id,
		CreatedAt: hospitalSetting.CreatedAt,
		UpdatedAt: hospitalSetting.UpdatedAt,
	}, nil
}

func (r *hospitalSettingRepo) Edit(ctx context.Context, hs *biz.HospitalSetting) (*biz.HospitalSetting, error) {
	db := r.data.db.WithContext(ctx)

	var hospitalSetting HospitalSetting
	result := db.Where("id = ?", hs.Id).First(&hospitalSetting)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, biz.ErrHospitalSettingNotFound.WithCause(err)
		}
		return nil, biz.ErrHospitalSettingSystemError.WithCause(err)
	}

	var hospitalSettings []HospitalSetting
	result = db.Where("id <> ?", hs.Id).Where(db.Where("name = ?", hs.Name).Or("registration_number = ?", hs.RegistrationNumber)).Find(&hospitalSettings)
	if err := result.Error; err == nil {
		if result.RowsAffected > 0 {
			return nil, biz.ErrHospitalSettingSameDataExists
		}
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, biz.ErrHospitalSettingSystemError.WithCause(err)
	}

	result = db.Model(&hospitalSetting).Updates(
		&HospitalSetting{
			Name:               hs.Name,
			RegistrationNumber: hs.RegistrationNumber,
			ContactPerson:      hs.ContactPerson,
			ContactMobile:      hs.ContactMobile,
			Locked:             hs.Locked,
			ApiUrl:             hs.ApiUrl,
			SignatureKey:       hs.SignatureKey,
		})
	if err := result.Error; err != nil {
		return nil, biz.ErrHospitalSettingSystemError.WithCause(err)
	}

	return &biz.HospitalSetting{
		Id:        hospitalSetting.Id,
		UpdatedAt: hospitalSetting.UpdatedAt,
	}, nil
}

func (r *hospitalSettingRepo) Delete(ctx context.Context, hs *biz.HospitalSetting) error {
	db := r.data.db.WithContext(ctx)
	result := db.Delete(&HospitalSetting{}, hs.Id)
	if err := result.Error; err != nil {
		return biz.ErrHospitalSettingSystemError.WithCause(err)
	}
	return nil
}

func (r *hospitalSettingRepo) DeleteInBatches(ctx context.Context, ids []uint64) error {
	db := r.data.db.WithContext(ctx)
	result := db.Delete(&HospitalSetting{}, ids)
	if err := result.Error; err != nil {
		return biz.ErrHospitalSettingSystemError.WithCause(err)
	}
	return nil
}
