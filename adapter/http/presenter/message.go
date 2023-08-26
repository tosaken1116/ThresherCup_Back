package presenter

import (
	e "errors"
	"log"
	"net/http"
	"thresher/usecase/model"
	"thresher/utils/errors"

	"github.com/gin-gonic/gin"
)

type IMessagePresenter interface {
	RenderMessage(Message model.Message)
	RenderMessages(Message []model.Message)
	RenderError(err error)
	RenderSuccess()
}

type MessagePresenter struct {
	ctx *gin.Context
}

func NewMessagePresenter(ctx *gin.Context) IMessagePresenter {
	return &MessagePresenter{
		ctx: ctx,
	}
}

func (mp *MessagePresenter) RenderMessage(message model.Message) {
	mp.ctx.JSON(http.StatusOK, message)
}
func (mp *MessagePresenter) RenderMessages(message []model.Message) {
	mp.ctx.JSON(http.StatusOK, message)
}

func (mp *MessagePresenter) RenderError(err error) {
	if e, ok := err.(*errors.Error); ok {
		log.Println(err.Error())
		mp.ctx.JSON(e.StatusCode, gin.H{"error": errors.ErrorResponse{StatusText: e.StatusText, Detail: e.Detail}})
		return
	}
	log.Println(err.Error())
	mp.ctx.JSON(http.StatusInternalServerError, e.New("unknown error"))
}
func (mp *MessagePresenter) RenderSuccess() {
	mp.ctx.JSON(http.StatusOK, errors.SuccessResponse{StatusText: "ok"})
}
