package service

import (
	"thresher/domain/repository"
	"thresher/infra/model"

	"github.com/gin-gonic/gin"
)

type IUserService interface {
	UpdateUser(ctx *gin.Context, userId string, name *string, description *string) error
	GetFollowing(ctx *gin.Context, userId string) (*[]model.Users, error)
	GetFollowed(ctx *gin.Context, userId string) (*[]model.Users, error)
}

type userService struct {
	repo repository.IUserRepository
}

func NewUserService(ur repository.IUserRepository) IUserService {
	return &userService{
		repo: ur,
	}
}

func (us *userService) UpdateUser(ctx *gin.Context, userId string, name *string, description *string) error {
	return us.repo.UpdateUser(ctx, userId, name, description)
}

func (us *userService) GetFollowing(ctx *gin.Context, userId string) (*[]model.Users, error) {
	return us.repo.GetFollowing(ctx, userId)
}
func(us *userService) GetFollowed(ctx *gin.Context, userId string) (*[]model.Users, error){
	return us.repo.GetFollowed(ctx,userId)
}