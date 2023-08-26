package service

import (
	"thresher/domain/repository"

	"github.com/gin-gonic/gin"
)

type ILocationService interface {
	CreateNewLocation(ctx *gin.Context,userId string,lat float32,lon float32 )error
}

type locationService struct {
	repo repository.ILocationRepository
}

func NewLocationService(lr repository.ILocationRepository) ILocationService{
	return &locationService{
		repo:lr,
	}
}
func (ls *locationService)CreateNewLocation(ctx *gin.Context,userId string,lat float32,lon float32 )error{
	return ls.repo.CreateNewLocation(ctx,userId,lat,lon)
}