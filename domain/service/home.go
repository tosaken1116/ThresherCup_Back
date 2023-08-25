package service

import (
	"thresher/domain/repository"
	"thresher/infra/model"

	"github.com/gin-gonic/gin"
)

type IHomeService interface {
	CreateNewHome(ctx *gin.Context,userId string,lat float32,lon float32,npr uint16)(*model.Home,error)
}

type homeService struct {
	repo repository.IHomeRepository
}

func NewHomeService(hr repository.IHomeRepository) IHomeService{
	return &homeService{
		repo:hr,
	}
}
func (hs *homeService) CreateNewHome(ctx *gin.Context,userId string,lat float32,lon float32,npr uint16)(*model.Home,error){
	h, err := hs.repo.CreateNewHome(ctx,userId,lat,lon,npr);
	if err != nil{
		return nil,err
	}
	return h,nil

}