package usecase

import (
	"net/http"
	"thresher/domain/service"
	"thresher/usecase/model"
	"thresher/utils/errors"

	"github.com/gin-gonic/gin"
)

type IEncounterUsecase interface{
	GetEncounter(ctx *gin.Context)(*[]model.Encounter,error)
}

type encounterUsecase struct{
	svc service.IEncounterService
}

func NewEncounterUsecase(es service.IEncounterService) IEncounterUsecase {
	return &encounterUsecase{
		svc: es,
	}
}

func (eu *encounterUsecase) GetEncounter(ctx *gin.Context)(*[]model.Encounter,error){
	userId ,gErr := ctx.Get("userId")
	if !gErr{
		return nil ,errors.New(http.StatusInternalServerError,"cannot get user_id","/usecase/home_usecase/CreateNewHome")
	}
	h,err := eu.svc.GetEncounter(ctx, userId.(string))
	if err != nil {
		return nil,err
	}
	if h == nil{
		return nil,nil
	}
	return model.EncountersFromDomainModels(h), nil
}