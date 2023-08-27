package service

import (
	"thresher/domain/repository"
	"thresher/infra/model"

	"github.com/gin-gonic/gin"
)

type IMessageService interface {
	GetMessages(ctx *gin.Context, senderId string, responderId string) (*[]model.Message, error)
	CreateMessage(ctx *gin.Context, senderId string, responderId string, content string) error
	GetUnreadMessages(ctx *gin.Context, userId string) (*[]model.Message, error)
	ChangeAutoResponse(ctx *gin.Context, senderId string, responderId string) error
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
func (ms *messageService) CreateMessage(ctx *gin.Context, senderId string, responderId string, content string) error {
	return ms.repo.CreateMessage(ctx, senderId, responderId, content)
}
func (ms *messageService) GetUnreadMessages(ctx *gin.Context, userId string) (*[]model.Message, error) {
	return ms.repo.GetUnreadMessages(ctx, userId)
}
func (ms *messageService) ChangeAutoResponse(ctx *gin.Context, senderId string, responderId string) error {
	return ms.repo.ChangeAutoResponse(ctx, senderId, responderId)
}
