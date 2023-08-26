package controller

import (
	"net/http"
	"thresher/adapter/http/presenter"
	"thresher/usecase"
	"thresher/usecase/model"
	"thresher/utils/errors"

	"github.com/gin-gonic/gin"
)

type IHomeController interface {
	CreateNewHome(ctx *gin.Context)
	GetMyHome(ctx *gin.Context)
}
type homeController struct {
	usc usecase.IHomeUsecase
}

func NewHomeController(hu usecase.IHomeUsecase) IHomeController {
	return &homeController{
		usc: hu,
	}
}

// @Summary 家の作成
// @Tags home
// @Accept  json
// @Produce  json
// @Security Bearer
// @Param       Home   body   model.InputHome   true  "Home"
// @Success 200 {object} model.Home
// @Failure 400 {object} errors.ErrorResponse
// @Failure 403 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 409 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /home [post]
func (hc *homeController) CreateNewHome(ctx *gin.Context) {
	presenter := presenter.NewHomePresenter(ctx)
	input := model.InputHome{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		presenter.RenderError(errors.New(http.StatusBadRequest, "Invalid Home param", "adapter/http/controller/home.go"))
		return
	}
	h, err := hc.usc.CreateNewHome(ctx, input)
	if err != nil {
		presenter.RenderError(err)
		return
	}
	presenter.RenderHome(*h)
}

// @Summary 家の取得
// @Tags home
// @Accept  json
// @Produce  json
// @Security Bearer
// @Success 200 {object} model.Home
// @Failure 400 {object} errors.ErrorResponse
// @Failure 403 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 409 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /home [get]
func (hc *homeController) GetMyHome(ctx *gin.Context) {
	presenter := presenter.NewHomePresenter(ctx)
	h, err := hc.usc.GetMyHome(ctx)
	if err != nil {
		presenter.RenderError(err)
		return
	}
	presenter.RenderHome(*h)
}
