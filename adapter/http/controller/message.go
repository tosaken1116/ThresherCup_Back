package controller

import (
	"thresher/usecase"
)

type IMessageController interface {
}
type messageController struct {
	usc usecase.IMessageUsecase
}

func NewMessageController(hu usecase.IMessageUsecase) IMessageController {
	return &messageController{
		usc: hu,
	}
}
