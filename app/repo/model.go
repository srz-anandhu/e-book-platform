package repo

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	CreatedBy int64          `gorm:"column:created_by"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	UpdatedBy *int64         `gorm:"column:updated_by"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	DeletedBy *int64         `gorm:"column:deleted_by"`
}
