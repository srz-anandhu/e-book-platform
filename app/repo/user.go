package repo

import (
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64          `gorm:"primaryKey"`
	Mail      string         `gorm:"column:mail;unique;not null"`
	Username  string         `gorm:"column:username;unique;not null"`
	Password  string         `gorm:"column:password;not null"`
	Salt      string         `gorm:"column:salt;not null"`
	IsDeleted bool           `gorm:"column:is_deleted;default:false"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	DeletedBy *int64         `gorm:"column:deleted_by"`
	Active    string         `gorm:"column:active"`
}

func (userModel *User) CreateUser(db *gorm.DB) (userID int64, err error) {
	result := db.Create(&userModel)

	if result.Error != nil {
		return 0, result.Error
	}

	return userModel.ID, nil
}

func GetOneUser(db *gorm.DB, id int64) (*User, error) {
	user := &User{}

	result := db.Unscoped().Where("id = ?", id).First(user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {

		return nil, fmt.Errorf("user not found with ID=%d due to : %v", id, result.Error)

	} else if result.Error != nil {
		return nil, result.Error
	}

	if user.IsDeleted {
		return nil, fmt.Errorf("user with ID : %d already deleted", id)
	}

	return user, nil
}

func DeleteUser(db *gorm.DB, id int64) error {
	// user := &User{}

	// result := db.Where("id=?", id).Delete(user)

	// if errors.Is(result.Error, gorm.ErrInvalidData) {
	// 	return result.Error
	// } else if result.Error != nil {
	// 	return result.Error
	// }

	// update is_deleted field to true
	updateResult := db.Table("users").Where("id=?", id).Updates(map[string]interface{}{
		"is_deleted": true,
		"deleted_at": time.Now().UTC(),
	})
	if updateResult.Error != nil {
		return updateResult.Error
	}

	log.Println("user deleted successfully...")
	return nil
}

func UpdateUser(db *gorm.DB, id int64, newPassword string) error {
	result := db.Table("users").Where("id=? AND is_deleted=?", id, false).Updates(map[string]interface{}{
		"password":   newPassword,
		"updated_at": time.Now().UTC(),
	})

	if result.Error != nil {
		return result.Error
	}
	// Check if any rows were updated
	if result.RowsAffected == 0 {
		return fmt.Errorf("no active user found with ID %d to update", id)
	}

	log.Println("user password updated successfully..")
	return nil
}

func GetAllUsers(db *gorm.DB) ([]*User, error) {
	var user []*User
	result := db.Where("is_deleted", false).Find(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
