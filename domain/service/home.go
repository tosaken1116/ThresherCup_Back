package service

import "thresher/domain/repository"

type IHomeService interface {}

type homeService struct {
	repo repository.IHomeRepository
}

func NewHomeService(hr repository.IHomeRepository) IHomeService{
	return &homeService{
		repo:hr,
	}
}