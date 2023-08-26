package repository

import (
	"gorm.io/gorm"
)

type IMessageRepository interface {
}

type messageRepository struct {
	Db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) IMessageRepository {
	return &messageRepository{
		Db: db,
	}
}
