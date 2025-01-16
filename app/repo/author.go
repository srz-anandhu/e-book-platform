package repo

import (
	"ebook/app/dto"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID     int64  `gorm:"primaryKey"`
	Name   string `gorm:"column:name"`
	Status bool   `gorm:"column:status;default:true"` // false if deleted
	BaseModel
}

type AuthorRepo interface {
	GetAuthor(id int) (*Author, error)
	CreateAuthor(authorReq *dto.AuthorCreateRequest) (int64, error)
	UpdateAuthor(updateReq *dto.AuthorUpdateRequest) error
	DeleteAuthor(deleteReq *dto.AuthorDeleteReq) error
	GetAllAuthors() ([]*Author, error)
}

type AuthorRepoImpl struct {
	db *gorm.DB
}

// For checking implementation of Repo interface
var _ AuthorRepo = (*AuthorRepoImpl)(nil)

func NewAuthorRepo(db *gorm.DB) AuthorRepo {
	return &AuthorRepoImpl{
		db: db,
	}
}

func (r *AuthorRepoImpl) GetAuthor(id int) (*Author, error) {
	authorResp := &Author{}
	result := r.db.Unscoped().Where("id = ?", id).First(authorResp)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	} else if result.Error != nil {
		return nil, result.Error
	}
	if !authorResp.Status {
		return nil, fmt.Errorf("author with id %d is marked as deleted", id)
	}
	return authorResp, nil
}

func (r *AuthorRepoImpl) CreateAuthor(authorReq *dto.AuthorCreateRequest) (int64, error) {
	author := &Author{
		ID:   int64(authorReq.ID),
		Name: authorReq.Name,
	}

	result := r.db.Create(&author)
	if result.Error != nil {
		return 0, result.Error
	}
	return author.ID, nil
}

func (r *AuthorRepoImpl) UpdateAuthor(updateReq *dto.AuthorUpdateRequest) error {
	result := r.db.Table("authors").Where("id=?", updateReq.ID).Updates(map[string]interface{}{
		"name":       updateReq.NewName,
		"updated_by": updateReq.UpdatedBy,
		"updated_at": time.Now().UTC(),
	})

	if result.Error != nil {
		return result.Error
	}
	// check if any rows were affected
	if result.RowsAffected == 0 {
		return fmt.Errorf("no active author found with ID : %d to update", updateReq.ID)
	}
	return nil
}

func (r *AuthorRepoImpl) DeleteAuthor(deleteReq *dto.AuthorDeleteReq) error {
	// Update status to false and deleted by
	updates := map[string]interface{}{
		"status":     false,
		"deleted_by": deleteReq.DeletedBy, // userID
		"deleted_at": time.Now().UTC(),
	}
	updateResult := r.db.Table("authors").Where("id=?", deleteReq.ID).Updates(updates)
	if updateResult.Error != nil {
		return updateResult.Error
	}

	return nil
}

func (r *AuthorRepoImpl) GetAllAuthors() ([]*Author, error) {
	var authors []*Author
	result := r.db.Where("status=?", true).Find(&authors)
	if result.Error != nil {
		return nil, result.Error
	}
	return authors, nil
}
