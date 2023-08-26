package service

import (
	"thresher/domain/repository"
	"thresher/infra/model"

	"github.com/gin-gonic/gin"
)

type IMessageService interface {
	GetMessages(ctx *gin.Context, senderId string, responderId string) (*[]model.Message, error)
}

type messageService struct {
	repo repository.IMessageRepository
}

func NewMessageService(mr repository.IMessageRepository) IMessageService {
	return &messageService{
		repo: mr,
	}
}

func (ms *messageService) GetMessages(ctx *gin.Context, senderId string, responderId string) (*[]model.Message, error) {
	return ms.repo.GetMessages(ctx, senderId, responderId)
}
