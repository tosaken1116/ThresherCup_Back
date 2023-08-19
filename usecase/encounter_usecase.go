package usecase

import "thresher/domain/service"

type IEncounterUsecase interface{}

type encounterUsecase struct{
	svc service.IEncounterService
}

func NewEncounterUsecase(es service.IEncounterService) IEncounterUsecase {
	return &encounterUsecase{
		svc: es,
	}
}