package service

import (
	"thresher/domain/repository"
	"thresher/infra/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IPostService interface {
	GetPostById(ctx *gin.Context, id uuid.UUID) (*model.Posts, *int64, error)
}

type postService struct {
	repo repository.IPostRepository
}

func NewPostService(pr repository.IPostRepository) IPostService {
	return &postService{
		repo: pr,
	}
}

func (ps postService) GetPostById(ctx *gin.Context, id uuid.UUID) (*model.Posts, *int64, error) {
	p, likedNum, err := ps.repo.GetPostById(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	return p, likedNum, nil
}
