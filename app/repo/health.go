package repo

import (
	"fmt"

	"gorm.io/gorm"
)

type HealthRepo interface {
	CheckPing() error
}

type HealthRepoImpl struct {
	db *gorm.DB
}

func NewHealthRepo(db *gorm.DB) HealthRepo {
	return &HealthRepoImpl{
		db: db,
	}
}

func (r *HealthRepoImpl) CheckPing() error {
	if r.db == nil {
		return fmt.Errorf("db connection is not initialised")
	}
	sqlDB, err := r.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}