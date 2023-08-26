package repository

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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

func (ur *userRepository) UpdateUser(ctx *gin.Context,userId string,)