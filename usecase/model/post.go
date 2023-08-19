package model

import (
	"thresher/infra/model"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UserID      uuid.UUID `json:"user_id"`
}

func PostFromDomainModel(m *model.Posts)*Post{
	p := &Post{
		ID : m.ID,
		Description: m.Description,
		CreatedAt: m.CreatedAt,
		UserID: m.UserID,
	}
	return p
}