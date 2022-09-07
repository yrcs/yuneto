package po

import "github.com/yrcs/yuneto/pkg/repo"

type Dict struct {
	repo.Base
	Id          uint64 `gorm:"primaryKey;autoIncrement:false;comment:ID"`
	ParentId    uint64 `gorm:"index;autoIncrement:false;comment:上级ID"`
	Name        string `gorm:"not null;comment:名称"`
	Value       string `gorm:"comment:值"`
	DictCode    string `gorm:"index;comment:编码"`
	HasChildren bool   `gorm:"-"`
}
