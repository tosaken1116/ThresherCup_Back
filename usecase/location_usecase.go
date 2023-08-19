package usecase

import "thresher/domain/service"

type ILocationUsecase interface{}

type locationUsecase struct{
	svc service.ILocationService
}

func NewLocationUsecase(ls service.ILocationService) ILocationUsecase {
	return &locationUsecase{
		svc: ls,
	}
}