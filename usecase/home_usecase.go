package usecase

import (
	"net/http"
	"thresher/domain/service"
	"thresher/usecase/model"
	"thresher/utils/errors"

	"github.com/gin-gonic/gin"
)

type IHomeUsecase interface{
	CreateNewHome(ctx *gin.Context,input model.InputHome)(*model.Home,error)
}

type homeUsecase struct{
	svc service.IHomeService
}

func NewHomeUsecase(hs service.IHomeService) IHomeUsecase {
	return &homeUsecase{
		svc: hs,
	}
}


func (hu *homeUsecase) CreateNewHome(ctx *gin.Context,input model.InputHome)(*model.Home,error){
	userId ,gErr := ctx.Get("userId")
	if !gErr{
		return nil ,errors.New(http.StatusInternalServerError,"cannot get user_id","/usecase/home_usecase/CreateNewHome")
	}
	h,err := hu.svc.CreateNewHome(ctx, userId.(string),input.Latitude,input.Longitude,input.NonPassRange)
	if err != nil {
		return nil,err
	}
	return model.HomeFromDomainModel(h), nil
}
