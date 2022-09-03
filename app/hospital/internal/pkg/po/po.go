package po

import "github.com/yrcs/yuneto/pkg/repo"

type HospitalSetting struct {
	repo.Base
	Name               string `gorm:"unique;comment:医院名称"`
	RegistrationNumber string `gorm:"unique;comment:医院登记号"`
	ContactPerson      string `gorm:"comment:联系人"`
	ContactMobile      string `gorm:"comment:联系人手机"`
	Locked             uint8  `gorm:"default:0;comment:锁定状态（0:未锁定，1:已锁定）"`
	ApiUrl             string `gorm:"comment:API 基础路径"`
	SignatureKey       string `gorm:"comment:签名密钥"`
}
