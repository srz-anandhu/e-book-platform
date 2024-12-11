package repo

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID        int64          `gorm:"primaryKey"`
	Name      string         `gorm:"column:name"`
	Status    bool           `gorm:"column:status;default:true"`
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

func GetOneAuthor(db *gorm.DB, id int64) (*Author, error) {
	author := &Author{}
	result := db.Unscoped().Where("id=?", id).First(author)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("author not found with ID : %d : %v", id, result.Error)

	} 

	if result.Error != nil {
		return nil, result.Error
	}

	if !author.Status {
		return nil, fmt.Errorf("author with ID : %d already deleted", id)
	}

	return author, nil
}
