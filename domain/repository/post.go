package repository

import (
	"fmt"
	"net/http"
	"thresher/infra/model"
	"thresher/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IPostRepository interface {
	GetPostById(ctx *gin.Context, id uuid.UUID) (*model.Posts, *int64, error)
}

type postRepository struct {
	Db *gorm.DB
}

func NewPostRepository(db *gorm.DB) IPostRepository {
	return &postRepository{
		Db: db,
	}
}
func (pr *postRepository) GetPostById(ctx *gin.Context, id uuid.UUID) (*model.Posts, *int64, error) {
	p := &model.Posts{ID: id}
	if err := pr.Db.Preload("User").First(&p).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, errors.New(http.StatusNotFound, "post is not found", fmt.Sprintf("/domain/repository/post\n%s", err.Error()))
		}
		return nil, nil, errors.New(http.StatusInternalServerError, "error was occur when get data", fmt.Sprintf("/domain/repository/post\n%s", err.Error()))
	}
	var likeNum int64
	if err := pr.Db.Table("likes").Where("posts_id = ?", id).Count(&likeNum).Error; err != nil {
		return nil, nil, errors.New(http.StatusInternalServerError, "error was occur when like num", fmt.Sprintf("/domain/repository/post\n%s", err.Error()))
	}

	return p, &likeNum, nil
}
