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
	CreateNewPost(ctx *gin.Context, userId string, description string) (*model.Posts, error)
	DeletePostById(ctx *gin.Context, userId string, postId uuid.UUID) error
	GetMyTimeLine(ctx *gin.Context, userId string) (*[]model.Posts, error)
	GetFollowTimeline(ctx *gin.Context, userId string) (*[]model.Posts, error)
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
	p := new(model.Posts)
	if err := pr.Db.Preload("User").Where("id = ?", id).First(&p).Error; err != nil {
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

func (pr *postRepository) CreateNewPost(ctx *gin.Context, userId string, description string) (*model.Posts, error) {
	p := &model.Posts{
		Description: description,
		UserID:      userId,
	}
	if err := pr.Db.Preload("Users").Create(&p).Error; err != nil {
		return nil, errors.New(http.StatusInternalServerError, "cannot create new post", fmt.Sprintf("/domain/repository/post\n%s", err.Error()))
	}
	return p, nil
}
func (pr *postRepository) DeletePostById(ctx *gin.Context, userId string, postId uuid.UUID) error {
	p, _, err := pr.GetPostById(ctx, postId)
	if err != nil {
		return err
	}
	if p.UserID != userId {
		return errors.New(http.StatusForbidden, "this post is not yours", fmt.Sprintf("/domain/repository/post\n%s", err.Error()))
	}
	if err := pr.Db.Delete(p).Error; err != nil {
		return errors.New(http.StatusInternalServerError, "cannot delete post", fmt.Sprintf("/domain/repository/post\n%s", err.Error()))
	}
	return nil
}

func (pr *postRepository) GetMyTimeLine(ctx *gin.Context, userId string) (*[]model.Posts, error) {
	var posts []model.Posts
	if err := pr.Db.Preload("User").Where("user_id = ?", userId).Find(&posts).Error; err != nil {
		return nil, errors.New(http.StatusInternalServerError, "cannot get posts", fmt.Sprintf("/domain/repository/post\n%s", err.Error()))
	}
	return &posts, nil
}

func (pr *postRepository) GetFollowTimeline(ctx *gin.Context, userId string) (*[]model.Posts, error) {
	var posts []model.Posts
	if err := pr.Db.Preload("User").Joins("JOIN following ON following.followed_id = posts.user_id").Find(&posts).Error; err != nil {
		return nil, errors.New(http.StatusInternalServerError, "cannot get posts", fmt.Sprintf("/domain/repository/post\n%d", err))
	}
	return &posts, nil
}
