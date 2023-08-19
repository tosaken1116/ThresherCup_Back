package model

import (
	"thresher/infra/model"
	"time"

	"github.com/google/uuid"
)

type Location struct {
	Longitude string    `json:"longitude"`
	Latitude  string    `json:"latitude"`
	CreatedAt time.Time `json:"created_at"`
	UserID    uuid.UUID `json:"user_id"`
}

func LocationFromDomainModel(m *model.Location)*Location{
	l := &Location{
		UserID : m.UserID,
		Longitude: m.Longitude,
		Latitude: m.Latitude,
		CreatedAt: m.CreatedAt,
	}
	return l
}