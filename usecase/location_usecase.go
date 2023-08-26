package usecase

import (
	"net/http"
	"thresher/domain/service"
	"thresher/usecase/model"
	"thresher/utils/errors"

	"github.com/gin-gonic/gin"
)

type ILocationUsecase interface {
	CreateNewLocation(ctx *gin.Context, input model.InputLocation) error
	GetMyLocations(ctx *gin.Context) (*[]model.Location, error)
}

type locationUsecase struct {
	svc service.ILocationService
}

func NewLocationUsecase(ls service.ILocationService) ILocationUsecase {
	return &locationUsecase{
		svc: ls,
	}
}

func (lu *locationUsecase) CreateNewLocation(ctx *gin.Context, input model.InputLocation) error {
	userId, gErr := ctx.Get("userId")
	if !gErr {
		return errors.New(http.StatusInternalServerError, "cannot get user_id", "/usecase/location_usecase/CreateNewLocation")
	}
	err := lu.svc.CreateNewLocation(ctx, userId.(string), input.Latitude, input.Longitude)
	if err != nil {
		return err
	}
	return nil
}

func (lu *locationUsecase) GetMyLocations(ctx *gin.Context) (*[]model.Location, error) {
	userId, gErr := ctx.Get("userId")
	if !gErr {
		return nil, errors.New(http.StatusInternalServerError, "cannot get user_id", "/usecase/location_usecase/GetMyLocations")
	}
	l, err := lu.svc.GetMyLocations(ctx, userId.(string))
	if err != nil {
		return nil, err
	}
	return model.LocationsFromDomainModels(l), nil
}
