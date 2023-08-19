package repository

import "gorm.io/gorm"

type IEncounterRepository interface {
}

type encounterRepository struct {
	Db *gorm.DB
}
func NewEncounterRepository(db *gorm.DB) IEncounterRepository {
	return &encounterRepository{
		Db: db,
	}
}