package repository

import (
	"thresher/domain/repository"

	"gorm.io/gorm"
)

type postRepository struct {
	Db *gorm.DB
}
func NewPostRepository(db *gorm.DB) repository.IPostRepository {
	return &postRepository{
		Db: db,
	}
}