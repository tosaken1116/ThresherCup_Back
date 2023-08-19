package repository

import (
	"thresher/domain/repository"

	"gorm.io/gorm"
)

type locationRepository struct {
	Db *gorm.DB
}
func NewLocationRepository(db *gorm.DB) repository.ILocationRepository {
	return &locationRepository{
		Db: db,
	}
}