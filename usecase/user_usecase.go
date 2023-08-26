package usecase

import (
	"net/http"
	"thresher/domain/service"
	"thresher/usecase/model"
	"thresher/utils/errors"

	"github.com/gin-gonic/gin"
)

type IUserUsecase interface {
	UpdateUser(ctx *gin.Context, input model.UpdateUser) error
	GetFollowing(ctx *gin.Context) (*[]model.Users, error)
	GetFollowed(ctx *gin.Context) (*[]model.Users, error)
	CreateFollow(ctx *gin.Context, targetId string) error
	DeleteFollow(ctx *gin.Context, targetId string) error
}

type userUsecase struct {
	svc service.IUserService
}

func NewUserUsecase(us service.IUserService) IUserUsecase {
	return &userUsecase{
		svc: us,
	}
}

func (uu *userUsecase) UpdateUser(ctx *gin.Context, input model.UpdateUser) error {
	userId, gErr := ctx.Get("userId")
	if !gErr {
		return errors.New(http.StatusInternalServerError, "cannot get user_id", "/usecase/user_usecase/UpdateUser")
	}
	err := uu.svc.UpdateUser(ctx, userId.(string), input.Name, input.Description)
	if err != nil {
		return err
	}
	return nil
}

func (uu *userUsecase) GetFollowing(ctx *gin.Context) (*[]model.Users, error) {
	userId, gErr := ctx.Get("userId")
	if !gErr {
		return nil, errors.New(http.StatusInternalServerError, "cannot get user_id", "/usecase/user_usecase/GetFollowing")
	}
	u, err := uu.svc.GetFollowing(ctx, userId.(string))
	if err != nil {
		return nil, err
	}
	return model.UsersFromDomainModels(u), nil
}

func (uu *userUsecase) GetFollowed(ctx *gin.Context) (*[]model.Users, error) {
	userId, gErr := ctx.Get("userId")
	if !gErr {
		return nil, errors.New(http.StatusInternalServerError, "cannot get user_id", "/usecase/user_usecase/GetFollowed")
	}
	u, err := uu.svc.GetFollowed(ctx, userId.(string))
	if err != nil {
		return nil, err
	}
	return model.UsersFromDomainModels(u), nil
}
func (uu *userUsecase) CreateFollow(ctx *gin.Context, targetId string) error {
	userId, gErr := ctx.Get("userId")
	if !gErr {
		return errors.New(http.StatusInternalServerError, "cannot get user_id", "/usecase/user_usecase/CreateFollow")
	}
	err := uu.svc.CreateFollow(ctx, userId.(string), targetId)
	if err != nil {
		return err
	}
	return nil
}
func (uu *userUsecase) DeleteFollow(ctx *gin.Context, targetId string) error {
	userId, gErr := ctx.Get("userId")
	if !gErr {
		return errors.New(http.StatusInternalServerError, "cannot get user_id", "/usecase/user_usecase/DeleteFollow")
	}
	err := uu.svc.DeleteFollow(ctx, userId.(string), targetId)
	if err != nil {
		return err
	}
	return nil
}
