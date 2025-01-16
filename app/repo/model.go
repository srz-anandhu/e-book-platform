package repo

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	CreatedBy int            `gorm:"column:created_by"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	UpdatedBy *int           `gorm:"column:updated_by"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	DeletedBy *int           `gorm:"column:deleted_by"`
}
