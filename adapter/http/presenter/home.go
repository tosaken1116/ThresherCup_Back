package presenter

import (
	e "errors"
	"log"
	"net/http"
	"thresher/usecase/model"
	"thresher/utils/errors"

	"github.com/gin-gonic/gin"
)


type IHomePresenter interface{
	RenderHome(home model.Home)
	RenderError(err error)
}

type HomePresenter struct{
	ctx *gin.Context
}

func NewHomePresenter(ctx *gin.Context)IHomePresenter{
	return &HomePresenter{
		ctx: ctx,
	}
}

func (hp *HomePresenter) RenderHome(home model.Home){
	hp.ctx.JSON(http.StatusOK,home)
}

func (hp *HomePresenter) RenderError(err error){
	if e, ok := err.(*errors.Error); ok {
		log.Println(err.Error())
		hp.ctx.JSON(e.StatusCode,gin.H{"error": errors.ErrorResponse{StatusText: e.StatusText, Detail: e.Detail}})
		return
	}
	log.Println(err.Error())
	hp.ctx.JSON(http.StatusInternalServerError, e.New("unknown error"))
}