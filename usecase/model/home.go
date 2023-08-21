package model

import (
	"thresher/infra/model"
)

type Home struct {
	UserID       string `json:"user_id"`
	Longitude    string    `json:"longitude"`
	Latitude     string    `json:"latitude"`
	NonPassRange uint16    `json:"non_pass_range"`
}

func HomeFromDomainModel(m *model.Home)*Home{
	h := &Home{
		UserID : m.UserID,
		Longitude: m.Longitude,
		Latitude: m.Latitude,
		NonPassRange: m.NonPassRange,
	}
	return h
}