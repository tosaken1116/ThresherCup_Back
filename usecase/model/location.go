package model

import (
	"thresher/infra/model"
	"time"
)

type Location struct {
	Longitude float32   `json:"longitude"`
	Latitude  float32   `json:"latitude"`
	CreatedAt time.Time `json:"created_at"`
	UserID    string    `json:"user_id"`
}

type InputLocation struct {
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

func LocationFromDomainModel(m *model.Location) *Location {
	l := &Location{
		UserID:    m.UserID,
		Longitude: m.Longitude,
		Latitude:  m.Latitude,
		CreatedAt: m.CreatedAt,
	}
	return l
}

func LocationsFromDomainModels(m *[]model.Location) *[]Location {
	locations := make([]Location, len(*m))
	for i, v := range *m {
		locations[i] = *LocationFromDomainModel(&v)
	}
	return &locations
}
