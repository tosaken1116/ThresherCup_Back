package model

import (
	"thresher/infra/model"
)

type Users struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
}

type UpdateUser struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

func UserFromDomainModel(m *model.Users) *Users {
	u := &Users{
		ID:          m.ID,
		Name:        m.Name,
		Email:       m.Email,
		ImageUrl:    m.ImageUrl,
		Description: m.Description,
	}
	return u
}

func UsersFromDomainModels(m *[]model.Users) *[]Users {
	users := make([]Users, len(*m))
	for i,v := range *m {
		users[i] = *UserFromDomainModel(&v)
	}
	return &users
}

