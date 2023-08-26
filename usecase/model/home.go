package model

import (
	"thresher/infra/model"
)

type Home struct {
	UserID       string  `json:"user_id"`
	Longitude    float32 `json:"longitude"`
	Latitude     float32 `json:"latitude"`
	NonPassRange uint16  `json:"non_pass_range"`
	User         Users   `json:"user"`
}

type InputHome struct {
	Longitude    float32 `json:"longitude"`
	Latitude     float32 `json:"latitude"`
	NonPassRange uint16  `json:"non_pass_range"`
}
type UpdateMyHome struct {
	Longitude    *float32 `json:"longitude"`
	Latitude     *float32 `json:"latitude"`
	NonPassRange *uint16  `json:"non_pass_range"`
}

func HomeFromDomainModel(m *model.Home) *Home {
	User := UserFromDomainModel(&m.User)
	h := &Home{
		UserID:       m.UserID,
		Longitude:    m.Longitude,
		Latitude:     m.Latitude,
		NonPassRange: m.NonPassRange,
		User:         *User,
	}
	return h
}
