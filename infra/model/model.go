package model

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID        uuid.UUID `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	IconPath  string    `json:"icon_path"`
	Email     string    `json:"email" gorm:"unique"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`

	Likes       []Posts   `json:"likes" gorm:"many2many:likes"`
	Following   []Users   `json:"following" gorm:"many2many:following foreignKey:ID;association_foreignKey:ID;joinForeignKey:followed_id;JoinReferences:following_id"`
	Followed    []Users   `json:"followed" gorm:"many2many:following foreignKey:ID;association_foreignKey:ID;joinForeignKey:following_id;JoinReferences:followed_id"`
	Blocking    []Users   `json:"blocking" gorm:"many2many:blocking foreignKey:ID;association_foreignKey:ID;joinForeignKey:blocking_id;JoinReferences:blocked_id"`
	Encountered Encounter `json:"encountered" gorm:"foreignKey:PassingId;"`
}

type Posts struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	UserID      uuid.UUID `json:"user_id"`

	User Users `json:"user"`
}

type Home struct {
	ID           uuid.UUID `json:"id" gorm:"primary_key"`
	UserID       uuid.UUID `json:"user_id"`
	Longitude    string    `json:"longitude"`
	Latitude     string    `json:"latitude"`
	NonPassRange uint16    `json:"non_pass_range"`

	User Users `json:"user"`
}

type Location struct {
	Longitude string    `json:"longitude"`
	Latitude  string    `json:"latitude"`
	CreatedAt time.Time `json:"created_at"`
	UserId    uuid.UUID `json:"user_id"`

	User Users `json:"user"`
}

type Encounter struct {
	Longitude string    `json:"longitude"`
	Latitude  string    `json:"latitude"`
	CreatedAt time.Time `json:"created_at"`
	PassingId uuid.UUID `json:"passing_id"`
	PassedId  uuid.UUID `json:"passed_id"`
}
