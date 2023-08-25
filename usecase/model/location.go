package model

import (
	"thresher/infra/model"
	"time"
)

type Location struct {
	Longitude float32    `json:"longitude"`
	Latitude  float32    `json:"latitude"`
	CreatedAt time.Time `json:"created_at"`
	UserID    string `json:"user_id"`
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