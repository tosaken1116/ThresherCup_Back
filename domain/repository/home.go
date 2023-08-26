package repository

import (
	"net/http"
	"thresher/infra/model"
	"thresher/utils/errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IHomeRepository interface {
	CreateNewHome(ctx *gin.Context, userId string, lat float32, lon float32, npr uint16) (*model.Home, error)
	GetMyHome(ctx *gin.Context, userId string) (*model.Home, error)
}

type homeRepository struct {
	Db *gorm.DB
}

func NewHomeRepository(db *gorm.DB) IHomeRepository {
	return &homeRepository{
		Db: db,
	}
}
func (hr *homeRepository) CreateNewHome(ctx *gin.Context, userId string, lat float32, lon float32, npr uint16) (*model.Home, error) {
	if homeExist := hr.Db.Where("user_id = ?", userId).First(&model.Home{}).Error; !errors.Is(homeExist, gorm.ErrRecordNotFound) {
		return nil, errors.New(http.StatusConflict, "your home is already exist", "domain/repository/home/CreateNewHome")
	}
	h := &model.Home{
		NonPassRange: npr,
		Longitude:    lon,
		Latitude:     lat,
		UserID:       userId,
	}
	if err := hr.Db.Preload("User").Create(&h).Error; err != nil {
		return nil, errors.New(http.StatusInternalServerError, "cannot create home", "/domain/repository/home/CreateNewHome")
	}
	return h, nil
}

func (hr *homeRepository) GetMyHome(ctx *gin.Context, userId string) (*model.Home, error) {
	h := &model.Home{}
	if err := hr.Db.Where("user_id = ?", userId).Preload("User").First(&h).Error; err != nil {
		return nil, errors.New(http.StatusInternalServerError, "cannot get home", "/domain/repository/home/GetMyHome")
	}
	return h, nil
}
