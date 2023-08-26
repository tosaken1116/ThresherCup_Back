package presenter

import (
	e "errors"
	"log"
	"net/http"
	"thresher/usecase/model"
	"thresher/utils/errors"

	"github.com/gin-gonic/gin"
)

type ILocationPresenter interface {
	RenderLocation(Location model.Location)
	RenderLocations(Location []model.Location)
	RenderError(err error)
	RenderSuccess()
}

type LocationPresenter struct {
	ctx *gin.Context
}

func NewLocationPresenter(ctx *gin.Context) ILocationPresenter {
	return &LocationPresenter{
		ctx: ctx,
	}
}

func (lp *LocationPresenter) RenderLocation(location model.Location) {
	lp.ctx.JSON(http.StatusOK, location)
}
func (lp *LocationPresenter) RenderLocations(location []model.Location) {
	lp.ctx.JSON(http.StatusOK, location)
}

func (lp *LocationPresenter) RenderSuccess() {
	lp.ctx.JSON(http.StatusOK, errors.SuccessResponse{StatusText: "Ok"})
}
func (lp *LocationPresenter) RenderError(err error) {
	if e, ok := err.(*errors.Error); ok {
		log.Println(err.Error())
		lp.ctx.JSON(e.StatusCode, gin.H{"error": errors.ErrorResponse{StatusText: e.StatusText, Detail: e.Detail}})
		return
	}
	log.Println(err.Error())
	lp.ctx.JSON(http.StatusInternalServerError, e.New("unknown error"))
}
