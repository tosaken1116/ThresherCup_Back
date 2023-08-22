package usecase

import (
	"log"
	"net/http"
	"thresher/domain/service"
	"thresher/usecase/model"
	"thresher/utils/cast"
	"thresher/utils/errors"

	"github.com/gin-gonic/gin"
)

type IPostUsecase interface {
	GetPostById(ctx *gin.Context, id string) (*model.Post, error)
	CreateNewPost(ctx *gin.Context,input model.InputPost)(*model.Post,error)
	DeletePostById(ctx *gin.Context, id string)error
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
		return nil, errors.New(http.StatusBadRequest, "id cannot be casted correctly", "/usecase/post_usecase/GetPostById")
	}
	p, likedNum, err := pu.svc.GetPostById(ctx, *uuid)
	if err != nil {
		return nil, err
	}
	return model.PostFromDomainModel(p, *likedNum), nil
}

func (pu *postUsecase) CreateNewPost(ctx *gin.Context,input model.InputPost)(*model.Post,error){
	log.Println(input)
	if input.Description ==""{
		return nil,errors.New(http.StatusBadRequest,"description is null string","/usecase/post_usecase/CreateNewPost")
	}
	userId ,gErr := ctx.Get("userId")
	if !gErr{
		return nil ,errors.New(http.StatusInternalServerError,"cannot get user_id","/usecase/post_usecase/CreateNewPost")
	}
	// cUserId,err := cast.CastStringToUUID(userId.(string))
	// if err != nil {
	// 	return nil, errors.New(http.StatusBadRequest, "id cannot be casted correctly", "/usecase/post_usecase/CreateNewPost")
	// }
	p,err := pu.svc.CreateNewPost(ctx, userId.(string),input.Description)
	if err != nil {
		return nil,err
	}
	return model.PostFromDomainModel(p,0),nil
}

func (pu *postUsecase) DeletePostById(ctx *gin.Context, id string)error{
	userId ,gErr := ctx.Get("userId")
	if !gErr{
		return errors.New(http.StatusInternalServerError,"cannot get user_id","/usecase/post_usecase/DeletePostById")
	}
	uuid, err := cast.CastStringToUUID(id)
	if err != nil {
		return  errors.New(http.StatusBadRequest, "id cannot be casted correctly", "/usecase/post_usecase/DeletePostById")
	}
	err = pu.svc.DeletePostById(ctx,userId.(string), *uuid)
	if err != nil {
		return err
	}
	return  nil
}