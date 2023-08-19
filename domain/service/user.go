package service

import "thresher/domain/repository"

type IUserService interface {}

type userService struct {
	repo repository.IUserRepository
}

func NewUserService(ur repository.IUserRepository) IUserService{
	return &userService{
		repo:ur,
	}
}