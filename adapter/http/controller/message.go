package controller

import (
	"thresher/adapter/http/presenter"
	"thresher/usecase"
	_ "thresher/usecase/model"

	"github.com/gin-gonic/gin"
)

type IMessageController interface {
	GetMessages(ctx *gin.Context)
}
type messageController struct {
	usc usecase.IMessageUsecase
}

func NewMessageController(mu usecase.IMessageUsecase) IMessageController {
	return &messageController{
		usc: mu,
	}
}

// @Summary メッセージの取得
// @Tags message
// @Accept  json
// @Produce  json
// @Security Bearer
// @Param       id   path   string   true  "ID"
// @Success 200 {object} []model.Message
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /message/{id} [get]
func (mc *messageController) GetMessages(ctx *gin.Context) {
	presenter := presenter.NewMessagePresenter(ctx)
	m, err := mc.usc.GetMessages(ctx, ctx.Param("id"))
	if err != nil {
		presenter.RenderError(err)
		return
	}
	presenter.RenderMessages(*m)
}
