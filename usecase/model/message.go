package model

import (
	"thresher/infra/model"
	"time"

	"github.com/google/uuid"
)

type InputMessage struct {
	Description string `json:"description"`
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
