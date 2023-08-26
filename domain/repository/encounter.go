package repository

import (
	"fmt"
	"net/http"
	"thresher/infra/model"
	"thresher/utils/errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IEncounterRepository interface {
	GetEncounter(ctx *gin.Context,userId string)(*[]model.Encounter,error)
}

type encounterRepository struct {
	Db *gorm.DB
}
func NewEncounterRepository(db *gorm.DB) IEncounterRepository {
	return &encounterRepository{
		Db: db,
	}
}

func (er *encounterRepository)GetEncounter(ctx *gin.Context,userId string)(*[]model.Encounter,error){
	e := []model.Encounter{}
	if err := er.Db.Where("passing_id = ?",userId).Or("passed_id = ?",userId).Find(&e).Error; err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound){
			return nil,nil
		}
		return nil,errors.New(http.StatusInternalServerError,"can not get encounter",fmt.Sprintf("/domain/repository/encounter\n%d", err))
	}
	return &e,nil
}