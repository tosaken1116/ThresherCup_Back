package controller

import (
	"thresher/adapter/http/presenter"
	"thresher/usecase"

	_ "thresher/usecase/model"

	"github.com/gin-gonic/gin"
)

type IEncounterController interface {
	GetEncounter(ctx *gin.Context)
}
type encounterController struct {
	usc usecase.IEncounterUsecase
}

func NewEncounterController(eu usecase.IEncounterUsecase) IEncounterController {
	return &encounterController{
		usc: eu,
	}
}

// @Summary すれ違いの取得
// @Tags encounter
// @Accept  json
// @Produce  json
// @Security Bearer
// @Success 200 {object} model.Encounter
// @Failure 401 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /encounter [get]
func (ec *encounterController) GetEncounter(ctx *gin.Context) {
	presenter := presenter.NewEncounterPresenter(ctx)
	e, err := ec.usc.GetEncounter(ctx)
	if err != nil {
		presenter.RenderError(err)
		return
	}
	presenter.RenderEncounters(*e)
}
