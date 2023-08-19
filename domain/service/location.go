package service

import "thresher/domain/repository"

type ILocationService interface {}

type locationService struct {
	repo repository.ILocationRepository
}

func NewLocationService(lr repository.ILocationRepository) ILocationService{
	return &locationService{
		repo:lr,
	}
}