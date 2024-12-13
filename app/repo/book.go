package repo

import (
	"fmt"
	"log"
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

func UpdateBook(db *gorm.DB, id int64, title, content string, userID, authorID int64, status int) error {
	result := db.Table("books").Where("id=? AND status IN (1,2)", id).Updates(map[string]interface{}{
		"title":      title,
		"content":    content,
		"author_id":  authorID,
		"updated_by": userID,
		"status":     status,
		"updated_at": time.Now().UTC(),
	})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no book found with ID : %d to update", id)
	}
	log.Printf("Book updated successfully by user : %d", userID)
	return nil
}
