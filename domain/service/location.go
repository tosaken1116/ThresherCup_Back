package service

import (
	"thresher/domain/repository"
	"thresher/infra/model"

	"github.com/gin-gonic/gin"
)

type ILocationService interface {
	CreateNewLocation(ctx *gin.Context, userId string, lat float32, lon float32) error
	GetMyLocations(ctx *gin.Context, userId string) (*[]model.Location, error)
}

type locationService struct {
	repo repository.ILocationRepository
}

func NewLocationService(lr repository.ILocationRepository) ILocationService {
	return &locationService{
		repo: lr,
	}
}
func (ls *locationService) CreateNewLocation(ctx *gin.Context, userId string, lat float32, lon float32) error {
	return ls.repo.CreateNewLocation(ctx, userId, lat, lon)
}
func (ls *locationService) GetMyLocations(ctx *gin.Context, userId string) (*[]model.Location, error) {
	return ls.repo.GetMyLocations(ctx, userId)
}
