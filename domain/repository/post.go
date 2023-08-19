package repository

import (
	"gorm.io/gorm"
)

type IPostRepository interface {
}

type postRepository struct {
	Db *gorm.DB
}
func NewPostRepository(db *gorm.DB) IPostRepository {
	return &postRepository{
		Db: db,
	}
}

