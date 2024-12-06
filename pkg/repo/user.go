package repo

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time  `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
	DeletedBy *int64     `gorm:"column:deleted_by"`
}

type User struct {
	ID        int64  `gorm:"primaryKey"`
	Mail      string `gorm:"column:mail;unique;not null"`
	Username  string `gorm:"column:username;unique;not null"`
	Password  string `gorm:"column:password;not null"`
	Salt      string `gorm:"column:salt;not null"`
	IsDeleted bool   `gorm:"column:is_deleted;default:false"`
	BaseModel
}

func (userRepo *User) CreateUser(db *gorm.DB) (userID int64, err error) {
	result := db.Create(&userRepo)

	if result.Error != nil {
		return 0, result.Error
	}

	return userRepo.ID, nil
}

func GetOneUser(db *gorm.DB, id int64) (user *User, err error) {

	result := db.First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found with id")
	} else if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
