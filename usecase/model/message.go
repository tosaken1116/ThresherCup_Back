package model

import (
	"thresher/infra/model"
	"time"

	"github.com/google/uuid"
)

type InputMessage struct {
	Content string `json:"content"`
}
type UnreadMessageOne struct {
	ID        uuid.UUID `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
type UnreadMessage struct {
	User     Users              `json:"user"`
	Messages []UnreadMessageOne `json:"messages"`
}

type Message struct {
	ID        uuid.UUID `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`

	Sender    Users `json:"sender"`
	Responder Users `json:"responder"`
}

func MessageFromDomainModel(m *model.Message) *Message {
	sender := UserFromDomainModel(&m.Sender)
	responder := UserFromDomainModel(&m.Responder)
	ms := &Message{
		ID:        m.ID,
		Content:   m.Content,
		CreatedAt: m.CreatedAt,
		Sender:    *sender,
		Responder: *responder,
	}
	return ms
}

func MessagesFromDomainModels(m *[]model.Message) *[]Message {
	messages := make([]Message, len(*m))
	for i, v := range *m {
		messages[i] = *MessageFromDomainModel(&v)
	}
	return &messages
}

func UnreadMessageFromDomainModels(m *[]model.Message) *map[string]UnreadMessage {
	unreadMessagesMap := make(map[string]UnreadMessage)
	for _, message := range *m {
		unreadMsgOne := UnreadMessageOne{
			ID:        message.ID,
			Content:   message.Content,
			CreatedAt: message.CreatedAt,
		}

		if _, exists := unreadMessagesMap[message.SenderID]; !exists {
			sender := UserFromDomainModel(&message.Sender)
			unreadMsg := UnreadMessage{User: *sender, Messages: []UnreadMessageOne{unreadMsgOne}}
			unreadMessagesMap[message.SenderID] = unreadMsg
		} else {
			unreadMsg := unreadMessagesMap[message.SenderID]
			unreadMsg.Messages = append(unreadMsg.Messages, unreadMsgOne)
			unreadMessagesMap[message.SenderID] = unreadMsg
		}
	}

	return &unreadMessagesMap
}
