package usecase

import (
	"net/http"
	"thresher/domain/service"
	"thresher/usecase/model"
	"thresher/utils/cast"
	"thresher/utils/errors"

	"github.com/gin-gonic/gin"
)

type IPostUsecase interface {
	GetPostById(ctx *gin.Context, id string) (*model.Post, error)
}

type postUsecase struct {
	svc service.IPostService
}

func NewPostUsecase(ps service.IPostService) IPostUsecase {
	return &postUsecase{
		svc: ps,
	}
}

func (pu *postUsecase) GetPostById(ctx *gin.Context, id string) (*model.Post, error) {
	uuid, err := cast.CastStringToUUID(id)
	if err != nil {
		return nil, errors.New(http.StatusBadRequest, "id cannot be casted correctly", "/usecase/post_usecase")
	}
	p, likedNum, err := pu.svc.GetPostById(ctx, *uuid)
	if err != nil {
		return nil, err
	}
	return model.PostFromDomainModel(p, *likedNum), nil
}
