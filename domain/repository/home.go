package repository

import "gorm.io/gorm"

type IHomeRepository interface {
}

type homeRepository struct {
	Db *gorm.DB
}
func NewHomeRepository(db *gorm.DB) IHomeRepository {
	return &homeRepository{
		Db: db,
	}
}