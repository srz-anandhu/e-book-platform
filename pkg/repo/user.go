package repo

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	DeletedBy *int64         `gorm:"column:deleted_by"`
}

type User struct {
	ID       int64  `gorm:"primaryKey"`
	Mail     string `gorm:"column:mail;unique;not null"`
	Username string `gorm:"column:username;unique;not null"`
	Password string `gorm:"column:password;not null"`
	Salt     string `gorm:"column:salt;not null"`
	IsDeleted bool   `gorm:"column:is_deleted;default:false"`
	BaseModel
}

func (userModel *User) CreateUser(db *gorm.DB) (userID int64, err error) {
	result := db.Create(&userModel)

	if result.Error != nil {
		return 0, result.Error
	}

	return userModel.ID, nil
}

func GetOneUser(db *gorm.DB, id int64) (user *User, err error) {

	//result := db.Where("id=?", id).First(&user)
	result := db.First(&user, id).Where("is_deleted=?", false)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("user not found with ID=%d due to : %v", id, result.Error)
	} else if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func DeleteUser(db *gorm.DB, id int64) (err error) {
	//var userModel *User
	isDeleted := map[string]interface{}{
		"is_deleted": true,
	}
	result := db.Table("users").Where("id=?", id).Updates(isDeleted)

	if errors.Is(result.Error, gorm.ErrInvalidData) {
		return err
	} else if result.Error != nil {
		return err
	} else {
		fmt.Println("user deleted successfully")
	}

	return nil
}
