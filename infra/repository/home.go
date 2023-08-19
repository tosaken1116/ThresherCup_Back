package repository

import (
	"thresher/domain/repository"

	"gorm.io/gorm"
)

type homeRepository struct {
	Db *gorm.DB
}
func NewHomeRepository(db *gorm.DB) repository.IHomeRepository {
	return &homeRepository{
		Db: db,
	}
}