package service

import (
	"thresher/domain/repository"
)

type IMessageService interface {
}

type messageService struct {
	repo repository.IMessageRepository
}

func NewMessageService(mr repository.IMessageRepository) IMessageService {
	return &messageService{
		repo: mr,
	}
}
