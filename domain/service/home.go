package service

import "thresher/domain/repository"

type IHomeService interface {}

type homeService struct {
	repo repository.IHomeRepository
}

func NewHomeService(ur repository.IHomeRepository) IHomeService{
	return &homeService{
		repo:ur,
	}
}