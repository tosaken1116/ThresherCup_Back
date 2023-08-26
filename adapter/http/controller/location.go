package controller

import (
	"net/http"
	"thresher/adapter/http/presenter"
	"thresher/usecase"
	"thresher/usecase/model"
	"thresher/utils/errors"

	"github.com/gin-gonic/gin"
)

type ILocationController interface {
	CreateNewLocation(ctx *gin.Context)
	GetMyLocations(ctx *gin.Context)
}
type locationController struct {
	usc usecase.ILocationUsecase
}

func NewLocationController(lu usecase.ILocationUsecase) ILocationController {
	return &locationController{
		usc: lu,
	}
}

// @Summary 現在地の作成
// @Tags location
// @Accept  json
// @Produce  json
// @Security Bearer
// @Param       Location   body   model.InputLocation   true  "Location"
// @Success 200 {object} model.Location
// @Failure 400 {object} errors.ErrorResponse
// @Failure 403 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /location [post]
func (lc *locationController) CreateNewLocation(ctx *gin.Context) {
	presenter := presenter.NewLocationPresenter(ctx)
	input := model.InputLocation{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		presenter.RenderError(errors.New(http.StatusBadRequest, "Invalid Location param", "adapter/http/controller/location.go"))
		return
	}
	err := lc.usc.CreateNewLocation(ctx, input)
	if err != nil {
		presenter.RenderError(err)
		return
	}
	presenter.RenderSuccess()
}

// @Summary 現在地の取得
// @Tags location
// @Accept  json
// @Produce  json
// @Security Bearer
// @Success 200 {object} []model.Location
// @Failure 400 {object} errors.ErrorResponse
// @Failure 403 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /location [get]
func (lc *locationController) GetMyLocations(ctx *gin.Context) {
	presenter := presenter.NewLocationPresenter(ctx)
	l, err := lc.usc.GetMyLocations(ctx)
	if err != nil {
		presenter.RenderError(err)
		return
	}
	presenter.RenderLocations(*l)
}
