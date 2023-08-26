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
}

type MessagePresenter struct {
	ctx *gin.Context
}

func NewMessagePresenter(ctx *gin.Context) IMessagePresenter {
	return &MessagePresenter{
		ctx: ctx,
	}
}

func (ep *MessagePresenter) RenderMessage(message model.Message) {
	ep.ctx.JSON(http.StatusOK, message)
}
func (ep *MessagePresenter) RenderMessages(message []model.Message) {
	ep.ctx.JSON(http.StatusOK, message)
}

func (ep *MessagePresenter) RenderError(err error) {
	if e, ok := err.(*errors.Error); ok {
		log.Println(err.Error())
		ep.ctx.JSON(e.StatusCode, gin.H{"error": errors.ErrorResponse{StatusText: e.StatusText, Detail: e.Detail}})
		return
	}
	log.Println(err.Error())
	ep.ctx.JSON(http.StatusInternalServerError, e.New("unknown error"))
}
