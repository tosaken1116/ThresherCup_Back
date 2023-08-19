package repository

import "gorm.io/gorm"

type IUserRepository interface {
}


type userRepository struct {
	Db *gorm.DB
}
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		Db: db,
	}
}