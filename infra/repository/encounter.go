package repository

import (
	"thresher/domain/repository"

	"gorm.io/gorm"
)

type encounterRepository struct {
	Db *gorm.DB
}
func NewEncounterRepository(db *gorm.DB) repository.IEncounterRepository {
	return &encounterRepository{
		Db: db,
	}
}