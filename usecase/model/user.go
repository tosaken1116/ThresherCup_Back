package model

import (
	"thresher/infra/model"

	"github.com/google/uuid"
)

type Users struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	ImageUrl string `json:"image_url"`
}

func UserFromDomainModel(m *model.Users)*Users{
	u := &Users{
		ID : m.ID,
		Name: m.Name,
		Email: m.Email,
		ImageUrl: m.ImageUrl,
	}
	return u
}