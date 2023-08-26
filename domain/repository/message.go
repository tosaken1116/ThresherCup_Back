package repository

import (
	"thresher/infra/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IMessageRepository interface {
	GetMessages(ctx *gin.Context, senderId string, responderId string) (*[]model.Message, error)
}

type messageRepository struct {
	Db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) IMessageRepository {
	return &messageRepository{
		Db: db,
	}
}

func (mr *messageRepository) GetMessages(ctx *gin.Context, senderId string, responderId string) (*[]model.Message, error) {
	var messages []model.Message
	if err := mr.Db.Preload("Sender").Preload("Responder").Where("sender_id = ? AND responder_id = ?", senderId, responderId).Find(&messages).Error; err != nil {
		return nil, err
	}
	return &messages, nil
}
