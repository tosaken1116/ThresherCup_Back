package usecase

import (
	"net/http"
	"thresher/domain/service"
	"thresher/usecase/model"
	"thresher/utils/errors"

	"github.com/gin-gonic/gin"
)

type IMessageUsecase interface {
	GetMessages(ctx *gin.Context, responderId string) (*[]model.Message, error)
	CreateNewMessage(ctx *gin.Context, responderId string, input model.InputMessage) error
}

type messageUsecase struct {
	svc service.IMessageService
}

func NewMessageUsecase(ms service.IMessageService) IMessageUsecase {
	return &messageUsecase{
		svc: ms,
	}
}

func (mu *messageUsecase) GetMessages(ctx *gin.Context, responderId string) (*[]model.Message, error) {
	userId, gErr := ctx.Get("userId")
	if !gErr {
		return nil, errors.New(http.StatusInternalServerError, "cannot get user_id", "/usecase/location_usecase/CreateNewLocation")
	}
	m, err := mu.svc.GetMessages(ctx, userId.(string), responderId)
	if err != nil {
		return nil, err
	}
	return model.MessagesFromDomainModels(m), nil
}

func (mu *messageUsecase) CreateNewMessage(ctx *gin.Context, responderId string, input model.InputMessage) error {
	userId, gErr := ctx.Get("userId")
	if !gErr {
		return errors.New(http.StatusInternalServerError, "cannot get user_id", "/usecase/location_usecase/CreateNewLocation")
	}
	err := mu.svc.CreateMessage(ctx, userId.(string), responderId, input.Content)
	if err != nil {
		return err
	}
	return nil
}
