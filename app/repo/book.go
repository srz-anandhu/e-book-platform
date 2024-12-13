package repo

import (
	"fmt"

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

func (bookModel *Book) CreateBook(db *gorm.DB) (bookID int64, err error) {
	result := db.Create(&bookModel)
	if result.Error != nil {
		return 0, result.Error
	}
	return bookModel.ID, nil
}

func GetOneBook(db *gorm.DB, id int64) (*Book, error) {
	book := &Book{}
	result := db.Unscoped().Where("id=? AND status IN (1,2)", id).First(book)
	if result.Error != nil {
		return nil, result.Error
	}

	if book.Status == 3 {
		return nil, fmt.Errorf("book was deleted")
	}

	return book, nil
}
