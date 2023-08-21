package model

import (
	"thresher/infra/model"
	"time"

	"github.com/google/uuid"
)

type InputPost struct{
	Description string `json:"description"`
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`

	User     Users `json:"user"`
	LikedNum int64 `json:"liked_num"`
}

func PostFromDomainModel(m *model.Posts, ln int64) *Post {
	u := UserFromDomainModel(&m.User)
	p := &Post{
		ID:          m.ID,
		Description: m.Description,
		CreatedAt:   m.CreatedAt,
		User:        *u,
		LikedNum:    ln,
	}
	return p
}
