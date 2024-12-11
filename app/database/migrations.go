package database

import (
	"ebook/app/repo"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func AutoMigrateModels(db *gorm.DB) error {
	if err := db.AutoMigrate(&repo.User{}); err != nil {
		return fmt.Errorf("user model migration failed due to : %v", err)
	}
	if err := db.AutoMigrate(&repo.Author{}); err != nil {
		return fmt.Errorf("author model migration failed due to : %v", err)
	}
	log.Println("migration successful..")
	return nil
}
