package repository

import "gorm.io/gorm"

type ILocationRepository interface {
}

type locationRepository struct {
	Db *gorm.DB
}
func NewLocationRepository(db *gorm.DB) ILocationRepository {
	return &locationRepository{
		Db: db,
	}
}