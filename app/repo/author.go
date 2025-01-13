package repo

import (
	"fmt"
	"log"
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
	// if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	// 	return nil, fmt.Errorf("author not found with ID : %d : %v", id, result.Error)

	// }
	if result.Error != nil {
		return nil, result.Error
	}
	if !author.Status {
		return nil, fmt.Errorf("author with ID : %d already deleted", id)
	}

	return author, nil
}

func DeleteAuthor(db *gorm.DB, id, userID int64) error {
	// author := &Author{}
	// result := db.Where("id=?", id).Delete(author)
	// if result.Error != nil {
	// 	return result.Error
	// }

	// Update status to false and deleted by
	updates := map[string]interface{}{
		"status":     false,
		"deleted_by": userID,
		"deleted_at": time.Now().UTC(),
	}
	updateResult := db.Table("authors").Where("id=?", id).Updates(updates)
	if updateResult.Error != nil {
		return updateResult.Error
	}
	log.Println("Author deleted successfully")
	return nil
}

func UpdateAuthor(db *gorm.DB, newName string, id, userID int64) error {
	result := db.Table("authors").Where("id=?", id).Updates(map[string]interface{}{
		"name":       newName,
		"updated_by": userID,
		"updated_at": time.Now().UTC(),
	})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no active author found with ID : %d to update", id)
	}
	log.Printf(" Author name updated successfully by user : %d", userID)
	return nil
}

func GetAllAuthors(db *gorm.DB) ([]*Author, error) {
	var authors []*Author
	result := db.Where("status=?", true).Find(&authors)
	if result.Error != nil {
		return nil, result.Error
	}
	return authors, nil
}
