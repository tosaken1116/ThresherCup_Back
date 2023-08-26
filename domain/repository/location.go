package repository

import (
	"net/http"
	"thresher/infra/model"
	"thresher/utils/calc"
	"thresher/utils/errors"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ILocationRepository interface {
	CreateNewLocation(ctx *gin.Context, userId string, lat float32, lon float32) error
	GetMyLocations(ctx *gin.Context, userId string) (*[]model.Location, error)
}

type locationRepository struct {
	Db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) ILocationRepository {
	return &locationRepository{
		Db: db,
	}
}

func (lr *locationRepository) CreateNewLocation(ctx *gin.Context, userId string, lat float32, lon float32) error {
	h := &model.Location{
		Longitude: lon,
		Latitude:  lat,
		UserID:    userId,
	}
	if err := lr.Db.Create(&h).Error; err != nil {
		return errors.New(http.StatusInternalServerError, "cannot create Location", "/domain/repository/Location/CreateNewLocation")
	}
	lr.CreateEncounter(ctx, userId, lat, lon)
	return nil
}

func (lr *locationRepository) GetMyLocations(ctx *gin.Context, userId string) (*[]model.Location, error) {
	h := &[]model.Location{}
	if err := lr.Db.Where("user_id = ?", userId).Find(&h).Error; err != nil {
		return nil, errors.New(http.StatusInternalServerError, "cannot get Location", "/domain/repository/Location/GetMyLocations")
	}
	return h, nil
}
func (lr *locationRepository) GetRecentLocations(minLat, maxLat, minLng, maxLng float32) ([]model.Location, error) {
	fiveMinutesAgo := time.Now().Add(-5 * time.Minute)

	var locations []model.Location
	if err := lr.Db.Where("latitude BETWEEN ? AND ?", minLat, maxLat).
		Where("longitude BETWEEN ? AND ?", minLng, maxLng).
		Where("created_at >= ?", fiveMinutesAgo).
		Find(&locations).Error; err != nil {
		return nil, err
	}

	return locations, nil
}
func (lr *locationRepository) IsHomeRange(userId string, lat float64, lon float64) bool {
	home := model.Home{}
	if err := lr.Db.Preload("").Where("user_id = ?", userId).First(&home); err != nil {
		return true
	}
	if calc.CalculateDistance(float64(home.Latitude), float64(home.Longitude), lat, lon, float64(home.NonPassRange)) {
		return true
	}
	return false
}
func (lr *locationRepository) CreateEncounter(ctx *gin.Context, userId string, lat float32, lon float32) {
	if lr.IsHomeRange(userId, float64(lat), float64(lon)) {
		return
	}
	locations, err := lr.GetRecentLocations(lat-0.0004, lat+0.0004, lon-0.0004, lon+0.0004)
	if err != nil {
		return
	}
	for _, v := range locations {
		if lr.IsHomeRange(userId, float64(v.Latitude), float64(v.Longitude)) {
			return
		}
		e := model.Encounter{
			PassingId: userId,
			PassedId:  v.UserID,
			Latitude:  v.Latitude,
			Longitude: v.Longitude,
		}
		if err := lr.Db.Create(&e).Error; err != nil {
			return
		}
	}
}
