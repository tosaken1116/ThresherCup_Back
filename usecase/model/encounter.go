package model

import (
	"thresher/infra/model"
	"time"
)

type Encounter struct {
	Longitude float32    `json:"longitude"`
	Latitude  float32    `json:"latitude"`
	CreatedAt time.Time `json:"created_at"`
	PassingId string `json:"passing_id"`
	PassedId  string `json:"passed_id"`
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

func EncountersFromDomainModels(m *[]model.Encounter)*[]Encounter{
	encounters := make([]Encounter,len(*m))
	for i, v := range *m {
		encounters[i] = *EncounterFromDomainModel(&v)
	}
	return &encounters
}