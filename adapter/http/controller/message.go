package controller

import (
	"thresher/usecase"
)

type IMessageController interface {
}
type messageController struct {
	usc usecase.IMessageUsecase
}

func NewMessageController(mu usecase.IMessageUsecase) IMessageController {
	return &messageController{
		usc: mu,
	}
}
