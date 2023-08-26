package service

import (
	"thresher/domain/repository"
	"thresher/infra/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IPostService interface {
	GetPostById(ctx *gin.Context, id uuid.UUID) (*model.Posts, *int64, error)
	CreateNewPost(ctx *gin.Context, userId string, description string) (*model.Posts, error)
	DeletePostById(ctx *gin.Context, userId string, postId uuid.UUID) error
	GetMyTimeLine(ctx *gin.Context, userId string) (*[]model.Posts, error)
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

func (ps postService) CreateNewPost(ctx *gin.Context, userId string, description string) (*model.Posts, error) {
	p, err := ps.repo.CreateNewPost(ctx, userId, description)
	if err != nil {
		return nil, err
	}
	return p, nil
}
func (ps postService) DeletePostById(ctx *gin.Context, userId string, postId uuid.UUID) error {
	err := ps.repo.DeletePostById(ctx, userId, postId)
	if err != nil {
		return err
	}
	return nil
}
func (ps postService) GetMyTimeLine(ctx *gin.Context, userId string) (*[]model.Posts, error) {
	p, err := ps.repo.GetMyTimeLine(ctx, userId)
	if err != nil {
		return nil, err
	}
	return p, nil
}
