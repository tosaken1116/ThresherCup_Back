package repository

import (
	"net/http"
	"thresher/infra/model"
	"thresher/utils/errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ILocationRepository interface {
	CreateNewLocation(ctx *gin.Context,userId string,lat float32,lon float32)error
}

type locationRepository struct {
	Db *gorm.DB
}
func NewLocationRepository(db *gorm.DB) ILocationRepository {
	return &locationRepository{
		Db: db,
	}
}

func (lr *locationRepository) CreateNewLocation(ctx *gin.Context,userId string,lat float32,lon float32)error{
	h := &model.Location{
		Longitude: lon,
		Latitude: lat,
		UserID: userId,
	}
	if err := lr.Db.Create(&h).Error; err!=nil{
		return  errors.New(http.StatusInternalServerError,"cannot create Location","/domain/repository/Location/CreateNewLocation")
	}
	return nil
}