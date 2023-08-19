package usecase

import "thresher/domain/service"

type IPostUsecase interface{}

type postUsecase struct{
	svc service.IPostService
}

func NewPostUsecase(ps service.IPostService) IPostUsecase {
	return &postUsecase{
		svc: ps,
	}
}