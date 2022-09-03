package repo

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/plugin/optimisticlock"
)

type Base struct {
	Id        uint64                 `gorm:"primaryKey;autoIncrement:false;comment:分布式全局唯一 ID"`
	CreatedAt time.Time              `gorm:"type:datetime;comment:创建时间"`
	UpdatedAt time.Time              `gorm:"type:datetime;comment:更新时间"`
	DeletedAt gorm.DeletedAt         `gorm:"type:datetime;index;comment:删除时间"`
	Version   optimisticlock.Version `gorm:"default:1;comment:版本号（乐观锁专用）"`
}
