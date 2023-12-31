package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.New()
	return
}

func (p *Posts) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Model(p).Preload("User").Find(p).Error
}
func (h *Home) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Model(h).Preload("User").Find(h).Error
}
func (l *Location) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Model(l).Preload("User").Find(l).Error
}

type Base struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;not null;primaryKey"`
}

type Users struct {
	ID          string    `json:"id" gorm:"primary_key; not null"`
	Name        string    `json:"name" gorm:"not null"`
	ImageUrl    string    `json:"image_url" gorm:"not null"`
	Email       string    `json:"email" gorm:"unique; not null"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	DeletedAt   time.Time `json:"deleted_at"`

	Likes         []Posts   `json:"likes" gorm:"many2many:likes"`
	Following     []Users   `json:"following" gorm:"many2many:following; foreignKey:ID;association_foreignKey:ID;joinForeignKey:followed_id;JoinReferences:following_id"`
	AutoResponder []Users   `json:"auto_responder" gorm:"many2many:auto_response; foreignKey:ID;association_foreignKey:ID;joinForeignKey:sender_id;JoinReferences:responder_id"`
	Followed      []Users   `json:"followed" gorm:"many2many:following; foreignKey:ID;association_foreignKey:ID;joinForeignKey:following_id;JoinReferences:followed_id"`
	Blocking      []Users   `json:"blocking" gorm:"many2many:blocking; foreignKey:ID;association_foreignKey:ID;joinForeignKey:blocking_id;JoinReferences:blocked_id"`
	Encountered   Encounter `json:"encountered" gorm:"foreignKey:PassingId;"`
}

type Posts struct {
	Base
	Description string    `json:"description" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	UserID      string    `json:"user_id" gorm:"not null"`

	User  Users   `json:"user"`
	Liked []Users `json:"liked" gorm:"many2many:likes"`
}

type Home struct {
	UserID       string  `json:"user_id" gorm:"primaryKey; not null"`
	Longitude    float32 `json:"longitude" gorm:"not null"`
	Latitude     float32 `json:"latitude" gorm:"not null"`
	NonPassRange uint16  `json:"non_pass_range" gorm:"not null"`

	User Users `json:"user"`
}

type Location struct {
	Base
	Longitude float32   `json:"longitude" gorm:"not null"`
	Latitude  float32   `json:"latitude" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UserID    string    `json:"user_id" gorm:"not null"`

	User Users `json:"user"`
}

type Encounter struct {
	Base
	Longitude float32   `json:"longitude" gorm:"not null"`
	Latitude  float32   `json:"latitude" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	PassingId string    `json:"passing_id" gorm:"not null"`
	PassedId  string    `json:"passed_id" gorm:"not null"`
}

type Message struct {
	Base
	SenderID    string    `json:"sender_id" gorm:"not null"`
	ResponderID string    `json:"responder_id" gorm:"not null"`
	Content     string    `json:"content" gorm:"not null"`
	IsRead      bool      `json:"is_read" gorm:"not null; default:false"`
	CreatedAt   time.Time `json:"created_at"`

	Sender    Users `json:"sender" gorm:"foreignKey:SenderID;"`
	Responder Users `json:"responder" gorm:"foreignKey:ResponderID;"`
}
