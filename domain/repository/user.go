package repository

import (
	"fmt"
	"net/http"
	"thresher/infra/model"
	"thresher/utils/errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IUserRepository interface {
	UpdateUser(ctx *gin.Context, userId string, name *string, description *string) error
	GetFollowing(ctx *gin.Context,userId string)(*[]model.Users,error)
}

type userRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		Db: db,
	}
}

func (ur *userRepository) UpdateUser(ctx *gin.Context, userId string, name *string, description *string) error {
	u := model.Users{
		ID: userId,
	}
	if err := ur.Db.First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(http.StatusNotFound, "user is not found", fmt.Sprintf("/domain/repository/user/UpdateUser\n%d", err))
		}
		return errors.New(http.StatusInternalServerError, "can not get user", fmt.Sprintf("/domain/repository/user/UpdateUser\n%d", err))
	}
	if name != nil {
		u.Name = *name
	}
	if description != nil {
		u.Description = *description
	}
	if err := ur.Db.Save(&u).Error; err != nil {
		return errors.New(http.StatusInternalServerError, "can not update user", fmt.Sprintf("/domain/repository/UpdateUser\n%d", err))
	}
	return nil
}

func (ur *userRepository)GetFollowing(ctx *gin.Context,userId string)(*[]model.Users,error){
	u := model.Users{
		ID: userId,
	}
	if err := ur.Db.First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil,errors.New(http.StatusNotFound, "user is not found", fmt.Sprintf("/domain/repository/user/UpdateUser\n%d", err))
		}
		return nil, errors.New(http.StatusInternalServerError, "can not get user", fmt.Sprintf("/domain/repository/user/UpdateUser\n%d", err))
	}
	var following *[]model.Users
	if err := ur.Db.Model(&u).Association("Following").Find(&following); err != nil {
		return nil,errors.New(http.StatusInternalServerError, "can not get following", fmt.Sprintf("/domain/repository/UpdateUser\n%d", err))
	}
	return following, nil
}