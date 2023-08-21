package controller

import (
	"thresher/adapter/http/presenter"
	"thresher/usecase"

	"github.com/gin-gonic/gin"
)

type IPostController interface {
	GetPostById(ctx *gin.Context)
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
