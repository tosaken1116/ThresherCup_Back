package controller

import (
	"net/http"
	"thresher/adapter/http/presenter"
	"thresher/usecase"
	"thresher/usecase/model"
	"thresher/utils/errors"

	"github.com/gin-gonic/gin"
)

type IPostController interface {
	GetPostById(ctx *gin.Context)
	CreateNewPost(ctx *gin.Context)
	DeletePostById(ctx *gin.Context)
}

type postController struct {
	usc usecase.IPostUsecase
}

func NewPostController(pu usecase.IPostUsecase) IPostController {
	return &postController{
		usc: pu,
	}
}

// @Summary 投稿の取得
// @Tags post
// @Accept  json
// @Produce  json
// @Security Bearer
// @Param       id   path   string   true  "ID"
// @Success 200 {object} model.Post
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /posts/{id} [get]
func (pc *postController) GetPostById(ctx *gin.Context) {
	presenter := presenter.NewPostPresenter(ctx)
	id := ctx.Param("id")
	p, err := pc.usc.GetPostById(ctx, id)
	if err != nil {
		presenter.RenderError(err)
		return
	}
	presenter.RenderPost(*p)
}

// @Summary 投稿の作成
// @Tags post
// @Accept  json
// @Produce  json
// @Security Bearer
// @Param       description   body   string   true  "Description"
// @Success 200 {object} model.Post
// @Failure 400 {object} errors.ErrorResponse
// @Failure 403 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /posts [post]
func (pc *postController) CreateNewPost(ctx *gin.Context) {
	presenter := presenter.NewPostPresenter(ctx)
	input := model.InputPost{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		presenter.RenderError(errors.New(http.StatusBadRequest,"Invalid Post param","/adapter/http/controller/post/CreateNewPost"))
	}
	p, err := pc.usc.CreateNewPost(ctx,input)
	if err != nil {
		presenter.RenderError(err)
		return
	}
	presenter.RenderPost(*p)
}


// @Summary 投稿の削除
// @Tags post
// @Accept  json
// @Produce  json
// @Security Bearer
// @Param       id   path   string   true  "ID"
// @Success 200 {object} errors.SuccessResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 401 {object} errors.ErrorResponse
// @Failure 403 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /posts/{id} [delete]
func (pc *postController) DeletePostById(ctx *gin.Context) {
	presenter := presenter.NewPostPresenter(ctx)
	id := ctx.Param("id")
	err := pc.usc.DeletePostById(ctx, id)
	if err != nil {
		presenter.RenderError(err)
		return
	}
	presenter.RenderDeleteSuccess()
}
