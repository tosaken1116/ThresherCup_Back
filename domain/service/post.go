package service

import "thresher/domain/repository"

type IPostService interface {}

type postService struct {
	repo repository.IPostRepository
}

func NewPostService(pr repository.IPostRepository) IPostService{
	return &postService{
		repo:pr,
	}
}