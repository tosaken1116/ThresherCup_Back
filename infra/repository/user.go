package repository

import (
	"thresher/domain/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	Db *gorm.DB
}
func NewUserRepository(db *gorm.DB) repository.IUserRepository {
	return &userRepository{
		Db: db,
	}
}