package service

import "thresher/domain/repository"

type IEncounterService interface {}

type encounterService struct {
	repo repository.IEncounterRepository
}

func NewEncounterService(er repository.IEncounterRepository) IEncounterService{
	return &encounterService{
		repo:er,
	}
}