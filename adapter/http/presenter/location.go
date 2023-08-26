package presenter

import (
	e "errors"
	"log"
	"net/http"
	"thresher/usecase/model"
	"thresher/utils/errors"

	"github.com/gin-gonic/gin"
)


type ILocationPresenter interface{
	RenderLocation(Location model.Location)
	RenderLocations(Location []model.Location)
	RenderError(err error)
	RenderSuccess()
}

type LocationPresenter struct{
	ctx *gin.Context
}

func NewLocationPresenter(ctx *gin.Context)ILocationPresenter{
	return &LocationPresenter{
		ctx: ctx,
	}
}

func (ep *LocationPresenter) RenderLocation(location model.Location){
	ep.ctx.JSON(http.StatusOK,location)
}
func (ep *LocationPresenter) RenderLocations(location []model.Location){
	ep.ctx.JSON(http.StatusOK,location)
}

func (ep *LocationPresenter) RenderSuccess(){
	ep.ctx.JSON(http.StatusOK,errors.SuccessResponse{StatusText: "Ok"})
}
func (ep *LocationPresenter) RenderError(err error){
	if e, ok := err.(*errors.Error); ok {
		log.Println(err.Error())
		ep.ctx.JSON(e.StatusCode,gin.H{"error": errors.ErrorResponse{StatusText: e.StatusText, Detail: e.Detail}})
		return
	}
	log.Println(err.Error())
	ep.ctx.JSON(http.StatusInternalServerError, e.New("unknown error"))
}