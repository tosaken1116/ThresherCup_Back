package repository

import (
	"fmt"
	"net/http"
	"thresher/infra/model"
	"thresher/utils/errors"
	"thresher/utils/openai"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IMessageRepository interface {
	GetMessages(ctx *gin.Context, senderId string, responderId string) (*[]model.Message, error)
	CreateMessage(ctx *gin.Context, senderId string, responderId string, content string) error
	GetUnreadMessages(ctx *gin.Context, userId string) (*[]model.Message, error)
	ChangeAutoResponse(ctx *gin.Context, senderId string, responderId string) error
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
		return nil, errors.New(http.StatusInternalServerError, "cannot get message", fmt.Sprintf("/domain/repository/message.go/GetMessages\n%d", err))
	}
	return &messages, nil
}
func (mr *messageRepository) CreateMessage(ctx *gin.Context, senderId string, responderId string, content string) error {
	// isFollowing := mr.Db.Table("following").Where("followed_id = ? AND following_id = ?", senderId, responderId).RowsAffected
	// if isFollowing == 0 {
	// 	return errors.New(http.StatusForbidden, "You are not following this user", "/domain/repository/message.go/GetMessages")
	// }
	m := &model.Message{
		SenderID:    senderId,
		ResponderID: responderId,
		Content:     content,
	}
	if err := mr.Db.Create(&m).Error; err != nil {
		return err
	}
	IsAutoResponse := mr.Db.Table("auto_response").Where("sender_id = ? AND responder_id = ?", responderId, senderId).Find(&model.Users{}).RowsAffected
	if IsAutoResponse != 0 {
		messages := []model.Message{}
		if err := mr.Db.Table("messages").Preload("Sender").Preload("Responder").Where("sender_id = ? AND responder_id = ?", responderId, senderId).Or("sender_id = ? AND responder_id = ?", senderId, responderId).Order("created_at desc").Find(&messages).Error; err != nil {
			return nil
		}
		fmt.Println(messages)
		me := &model.Users{}
		if err := mr.Db.Table("users").Where("id = ?", senderId).First(&me).Error; err != nil {
			return nil
		}
		fmt.Println(me)
		content, err := openai.GetOpenAiChat(ctx, messages, *me)
		if err != nil {
			return err
		}
		m := &model.Message{
			SenderID:    responderId,
			ResponderID: senderId,
			Content:     *content,
		}
		if err := mr.Db.Create(&m).Error; err != nil {
			return err
		}
	}
	return nil
}

func (mr *messageRepository) GetUnreadMessages(ctx *gin.Context, userId string) (*[]model.Message, error) {
	var messages []model.Message
	if err := mr.Db.Preload("Sender").Preload("Responder").Where("responder_id = ? AND is_read = ?", userId, false).Find(&messages).Error; err != nil {
		return nil, err
	}
	return &messages, nil
}

type AutoResponder struct {
	SenderID    string
	ResponderID string
}

func (mr *messageRepository) ChangeAutoResponse(ctx *gin.Context, senderId string, responderId string) error {
	autoResponder := []AutoResponder{}
	if IsAutoResponse := mr.Db.Raw("SELECT FROM auto_response WHERE sender_id = ? AND responder_id = ?", senderId, responderId).Scan(&autoResponder).RowsAffected; IsAutoResponse != 0 {
		fmt.Println("===")
		if err := mr.Db.Raw("DELETE FROM auto_response WHERE sender_id = ? AND responder_id = ?", senderId, responderId).Scan(&autoResponder).Error; err != nil {
			return err
		}
		fmt.Println("=====")
		fmt.Println(IsAutoResponse)
		return nil
	}
	fmt.Println(autoResponder)
	fmt.Println("=====================")
	if err := mr.Db.Raw("INSERT INTO auto_response (sender_id, responder_id) VALUES (?, ?)", senderId, responderId).Scan(&autoResponder).Error; err != nil {
		return err
	}
	return nil
}
