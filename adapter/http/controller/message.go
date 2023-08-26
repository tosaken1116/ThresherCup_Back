package controller

import (
	"net/http"
	"thresher/adapter/http/presenter"
	"thresher/usecase"
	"thresher/usecase/model"
	"thresher/utils/errors"

	"github.com/gin-gonic/gin"
)

type IMessageController interface {
	GetMessages(ctx *gin.Context)
	CreateNewMessage(ctx *gin.Context)
	GetUnreadMessages(ctx *gin.Context)
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

// @Summary メッセージの作成
// @Tags message
// @Accept  json
// @Produce  json
// @Security Bearer
// @Param       id   path   string   true  "ID"
// @Param       message   body   model.InputMessage   true  "Message"
// @Success 200 {object} []model.Message
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /message/{id} [post]
func (mc *messageController) CreateNewMessage(ctx *gin.Context) {
	presenter := presenter.NewMessagePresenter(ctx)
	input := model.InputMessage{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		presenter.RenderError(errors.New(http.StatusBadRequest, "Invalid Post param", "/adapter/http/controller/post/CreateNewPost"))
	}
	err := mc.usc.CreateNewMessage(ctx, ctx.Param("id"), input)
	if err != nil {
		presenter.RenderError(err)
		return
	}
	presenter.RenderSuccess()
}

// @Summary 未読メッセージの取得
// @Tags message
// @Accept  json
// @Produce  json
// @Security Bearer
// @Success 200 {object} map[string][]model.UnreadMessage
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /message/unread [get]
func (mc *messageController) GetUnreadMessages(ctx *gin.Context) {
	presenter := presenter.NewMessagePresenter(ctx)
	m, err := mc.usc.GetUnreadMessages(ctx)
	if err != nil {
		presenter.RenderError(err)
		return
	}
	presenter.RenderUnreadMessages(*m)
}
