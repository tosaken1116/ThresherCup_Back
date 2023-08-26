package usecase

import (
	"net/http"
	"thresher/domain/service"
	"thresher/usecase/model"
	"thresher/utils/errors"

	"github.com/gin-gonic/gin"
)

type ILocationUsecase interface{
	CreateNewLocation(ctx *gin.Context,input model.InputLocation)error
}

type locationUsecase struct{
	svc service.ILocationService
}

func NewLocationUsecase(ls service.ILocationService) ILocationUsecase {
	return &locationUsecase{
		svc: ls,
	}
}

func (lu *locationUsecase) CreateNewLocation(ctx *gin.Context,input model.InputLocation)error{
	userId ,gErr := ctx.Get("userId")
	if !gErr{
		return errors.New(http.StatusInternalServerError,"cannot get user_id","/usecase/location_usecase/CreateNewLocation")
	}
	err := lu.svc.CreateNewLocation(ctx, userId.(string),input.Latitude,input.Longitude)
	if err != nil {
		return err
	}
	return nil
}