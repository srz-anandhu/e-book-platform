package repo

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID        int64          `gorm:"primaryKey"`
	Name      string         `gorm:"column:name"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	CreatedBy int64          `gorm:"column:created_by"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoCreateTime"`
	UpdatedBy int64          `gorm:"column:updated_by"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	DeletedBy *int64         `gorm:"column:deleted_by"`
}

func (authorModel *Author) CreateAuthor(db *gorm.DB) (authorID int64, err error) {
	result := db.Create(&authorModel)
	if result.Error != nil {
		return 0, result.Error
	}
	return authorModel.ID, nil
}
