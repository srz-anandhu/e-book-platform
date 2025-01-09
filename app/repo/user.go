package repo

import (
	"ebook/app/dto"
	"errors"
	"fmt"
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

type UserRepo interface {
	CreateUser(userReq *dto.UserCreateRequest) (int64, error)
	GetUser(id int) (userResp *dto.UserResponse, err error)
	DeleteUser(id int) error
	UpdateUser(updateReq *dto.UserUpdateRequest) error
	GetAllUsers() ([]*dto.UserResponse, error)
}

type UserRepoImpl struct {
	db *gorm.DB
}

// For checking implementation of Repo interface
var _ UserRepo = (*UserRepoImpl)(nil)

func NewUserRepo(db *gorm.DB) UserRepo {
	return &UserRepoImpl{
		db: db,
	}
}

var user *User

func (r *UserRepoImpl) CreateUser(userReq *dto.UserCreateRequest) (int64, error) {
	result := r.db.Create(&user)

	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}

func (r *UserRepoImpl) GetUser(id int) (userResp *dto.UserResponse, err error) {
	userResp = &dto.UserResponse{}
	result := r.db.Unscoped().Where("id = ?", id).First(userResp)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {

		return nil, result.Error

	} else if result.Error != nil {
		return nil, result.Error
	}

	if userResp.IsDeleted {
		return nil, fmt.Errorf("user with id %d is marked as deleted", id)
	}
	return userResp, nil
}

// soft delete
func (r *UserRepoImpl) DeleteUser(id int) error {
	updateResult := r.db.Table("users").Where("id = ? AND is_deleted = ?", id, false).Updates(map[string]interface{}{
		"deleted_at": time.Now().UTC(),
		"is_deleted": true,
	})
	if updateResult.Error != nil {
		return updateResult.Error
	}
	return nil
}

func (r *UserRepoImpl) UpdateUser(updateReq *dto.UserUpdateRequest) error {
	result := r.db.Table("users").Where("id = ? AND is_deleted = ?", updateReq.ID, false).Updates(map[string]interface{}{
		"username":   updateReq.NewUsername,
		"mail":       updateReq.NewMail,
		"password":   updateReq.NewPassword,
		"updated_at": time.Now().UTC(),
	})

	if result.Error != nil {
		return result.Error
	}
	// check if any rows were affected
	if result.RowsAffected == 0 {
		return fmt.Errorf("no active user found with ID %d to update", updateReq.ID)
	}
	return nil
}

func (r *UserRepoImpl) GetAllUsers() ([]*dto.UserResponse, error) {
	var user []*dto.UserResponse
	result := r.db.Where("is_deleted", false).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
