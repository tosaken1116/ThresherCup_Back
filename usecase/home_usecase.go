package usecase

import "thresher/domain/service"

type IHomeUsecase interface{}

type homeUsecase struct{
	svc service.IHomeService
}

func NewHomeUsecase(hs service.IHomeService) IHomeUsecase {
	return &homeUsecase{
		svc: hs,
	}
}