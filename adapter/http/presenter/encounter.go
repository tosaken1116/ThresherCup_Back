package presenter


import (
	e "errors"
	"log"
	"net/http"
	"thresher/usecase/model"
	"thresher/utils/errors"

	"github.com/gin-gonic/gin"
)


type IEncounterPresenter interface{
	RenderEncounter(Encounter model.Encounter)
	RenderEncounters(Encounter []model.Encounter)
	RenderError(err error)
}

type EncounterPresenter struct{
	ctx *gin.Context
}

func NewEncounterPresenter(ctx *gin.Context)IEncounterPresenter{
	return &EncounterPresenter{
		ctx: ctx,
	}
}

func (ep *EncounterPresenter) RenderEncounter(encounter model.Encounter){
	ep.ctx.JSON(http.StatusOK,encounter)
}
func (ep *EncounterPresenter) RenderEncounters(encounter []model.Encounter){
	ep.ctx.JSON(http.StatusOK,encounter)
}


func (ep *EncounterPresenter) RenderError(err error){
	if e, ok := err.(*errors.Error); ok {
		log.Println(err.Error())
		ep.ctx.JSON(e.StatusCode,gin.H{"error": errors.ErrorResponse{StatusText: e.StatusText, Detail: e.Detail}})
		return
	}
	log.Println(err.Error())
	ep.ctx.JSON(http.StatusInternalServerError, e.New("unknown error"))
}