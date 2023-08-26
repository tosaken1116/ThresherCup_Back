package usecase

import (
	"thresher/domain/service"
)

type IMessageUsecase interface {
}

type messageUsecase struct {
	svc service.IMessageService
}

func NewMessageUsecase(ms service.IMessageService) IMessageUsecase {
	return &messageUsecase{
		svc: ms,
	}
}
