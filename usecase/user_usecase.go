package usecase

import (
	"thresher/domain/service"
)

type IUserUsecase interface{
}

type userUsecase struct {
svc service.IUserService
}

func NewUserUsecase(us  service.IUserService) IUserUsecase{
	return &userUsecase{
		svc: us,
	}
}