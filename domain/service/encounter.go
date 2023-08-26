package service

import (
	"thresher/domain/repository"
	"thresher/infra/model"

	"github.com/gin-gonic/gin"
)

type IEncounterService interface {
	GetEncounter(ctx *gin.Context,userId string)(*[]model.Encounter,error)
}

type encounterService struct {
	repo repository.IEncounterRepository
}

func NewEncounterService(er repository.IEncounterRepository) IEncounterService{
	return &encounterService{
		repo:er,
	}
}

func (es *encounterService) GetEncounter(ctx *gin.Context,userId string)(*[]model.Encounter,error){
	return es.repo.GetEncounter(ctx,userId)
}