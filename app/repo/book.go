package repo

import (
	"ebook/app/dto"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID       int64  `gorm:"primaryKey"`
	Title    string `gorm:"column:title"`
	Content  string `gorm:"column:content"`
	AuthorID int64  `gorm:"column:author_id"`
	Status   int    `gorm:"column:status"` // 1 - Drafted, 2 - Published, 3 - Deleted
	BaseModel
}

type BookRepo interface {
	GetBook(id int) (*dto.BookResponse, error)
	CreateBook(*dto.BookCreateRequest) (int64, error)
	UpdateBook(bookUpdateReq *dto.BookUpdateRequest) error
	DeleteBook(bookDeleteReq *dto.BookDeleteRequest) error
	GetAllBooks() ([]*dto.BookResponse, error)
}

type BookRepoImpl struct {
	db *gorm.DB
}

// For checking implementation of BookRepi interface
var _ BookRepo = (*BookRepoImpl)(nil)

func NewBookRepo(db *gorm.DB) BookRepo {
	return &BookRepoImpl{
		db: db,
	}
}

func (r *BookRepoImpl) GetBook(id int) (*dto.BookResponse, error) {
	bookRes := &dto.BookResponse{}
	result := r.db.Unscoped().Where("id=? AND status IN (1,2)", id).First(bookRes)
	if result.Error != nil {
		return nil, result.Error
	}

	if bookRes.Status == 3 {
		return nil, fmt.Errorf("book was deleted")
	}

	return bookRes, nil
}

var book *Book

func (r *BookRepoImpl) CreateBook(*dto.BookCreateRequest) (int64, error) {
	result := r.db.Create(&book)
	if result.Error != nil {
		return 0, result.Error
	}
	return book.ID, nil
}

func (r *BookRepoImpl) UpdateBook(bookUpdateReq *dto.BookUpdateRequest) error {
	result := r.db.Table("books").Where("id=? AND status IN (1,2)", bookUpdateReq.ID).Updates(map[string]interface{}{
		"title":      bookUpdateReq.Title,
		"content":    bookUpdateReq.Content,
		"author_id":  bookUpdateReq.AuthorID,
		"updated_by": bookUpdateReq.UpdatedBy,
		"status":     bookUpdateReq.Status,
		"updated_at": time.Now().UTC(),
	})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no book found with ID : %d to update", bookUpdateReq.ID)
	}
	return nil
}

func (r *BookRepoImpl) DeleteBook(bookDeleteReq *dto.BookDeleteRequest) error {
	updateResult := r.db.Table("books").Where("id=? AND status IN (1,2)", bookDeleteReq.ID).Updates(map[string]interface{}{
		"status":     3,
		"deleted_by": bookDeleteReq.DeletedBy,
		"deleted_at": time.Now().UTC(),
	})
	if updateResult.Error != nil {
		return updateResult.Error
	}

	return nil
}

func (r *BookRepoImpl) GetAllBooks() ([]*dto.BookResponse, error) {
	var books []*dto.BookResponse
	result := r.db.Where("status IN (1,2)").Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}
