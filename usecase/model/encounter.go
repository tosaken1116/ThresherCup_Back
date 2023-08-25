package model

import (
	"thresher/infra/model"
	"time"

	"github.com/google/uuid"
)

type Encounter struct {
	Longitude float32    `json:"longitude"`
	Latitude  float32    `json:"latitude"`
	CreatedAt time.Time `json:"created_at"`
	PassingId uuid.UUID `json:"passing_id"`
	PassedId  uuid.UUID `json:"passed_id"`
}

func EncounterFromDomainModel(m *model.Encounter)*Encounter{
	e := &Encounter{
		Longitude : m.Longitude,
		Latitude: m.Latitude,
		CreatedAt: m.CreatedAt,
		PassingId: m.PassingId,
		PassedId: m.PassedId,
	}
	return e
}