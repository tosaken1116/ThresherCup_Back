package presenter

import (
	e "errors"
	"log"
	"net/http"
	"thresher/usecase/model"
	"thresher/utils/errors"

	"github.com/gin-gonic/gin"
)

type IUserPresenter interface {
	RenderUser(user model.Users)
	RenderUsers(user *[]model.Users)
	RenderError(err error)
	RenderDeleteSuccess()
	RenderUpdateSuccess()
}
type UserPresenter struct {
	ctx *gin.Context
}

func NewUserPresenter(ctx *gin.Context) IUserPresenter {
	return &UserPresenter{
		ctx: ctx,
	}
}
func (up *UserPresenter) RenderUser(user model.Users) {
	up.ctx.JSON(http.StatusOK, user)
}

func (up *UserPresenter) RenderUsers(user *[]model.Users) {
	up.ctx.JSON(http.StatusOK, user)
}


func (up *UserPresenter) RenderError(err error) {
	if e, ok := err.(*errors.Error); ok {
		log.Println(err.Error())
		up.ctx.JSON(e.StatusCode, errors.ErrorResponse{StatusText: e.StatusText, Detail: e.Detail})
		return
	}
	log.Println(err.Error())
	up.ctx.JSON(http.StatusInternalServerError, e.New("unknown error"))

}

func (up *UserPresenter)RenderDeleteSuccess(){
	up.ctx.JSON(http.StatusOK,errors.SuccessResponse{StatusText: "ok"})
}

func (up *UserPresenter)RenderUpdateSuccess(){
	up.ctx.JSON(http.StatusOK,errors.SuccessResponse{StatusText: "ok"})
}
