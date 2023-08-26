package presenter

import (
	e "errors"
	"log"
	"net/http"
	"thresher/usecase/model"
	"thresher/utils/errors"

	"github.com/gin-gonic/gin"
)

type IPostPresenter interface {
	RenderPost(post model.Post)
	RenderPosts(post *[]model.Post)
	RenderError(err error)
	RenderDeleteSuccess()
}
type PostPresenter struct {
	ctx *gin.Context
}

func NewPostPresenter(ctx *gin.Context) IPostPresenter {
	return &PostPresenter{
		ctx: ctx,
	}
}
func (pp *PostPresenter) RenderPost(post model.Post) {
	pp.ctx.JSON(http.StatusOK, post)
}

func (pp *PostPresenter) RenderPosts(posts *[]model.Post) {
	pp.ctx.JSON(http.StatusOK, posts)
}
func (pp *PostPresenter) RenderError(err error) {
	if e, ok := err.(*errors.Error); ok {
		log.Println(err.Error())
		pp.ctx.JSON(e.StatusCode, errors.ErrorResponse{StatusText: e.StatusText, Detail: e.Detail})
		return
	}
	log.Println(err.Error())
	pp.ctx.JSON(http.StatusInternalServerError, e.New("unknown error"))

}

func (pp *PostPresenter) RenderDeleteSuccess() {
	pp.ctx.JSON(http.StatusOK, errors.SuccessResponse{StatusText: "ok"})
}
