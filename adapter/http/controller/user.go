package controller

import (
	"net/http"
	"thresher/adapter/http/presenter"
	"thresher/usecase"
	"thresher/usecase/model"
	"thresher/utils/errors"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	UpdateUser(ctx *gin.Context)
	GetFollowing(ctx *gin.Context)
	GetFollowed(ctx *gin.Context)
}

type userController struct {
	usc usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{
		usc: uu,
	}
}

// @Summary ユーザーの更新
// @Tags user
// @Accept  json
// @Produce  json
// @Security Bearer
// @Param       User   body   model.UpdateUser   true  "User"
// @Success 200 {object} model.Users
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /users [put]
func (pc *userController) UpdateUser(ctx *gin.Context) {
	presenter := presenter.NewUserPresenter(ctx)
	input := model.UpdateUser{}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		presenter.RenderError(errors.New(http.StatusBadRequest,"Invalid User param","/adapter/http/controller/user/UpdateUser"))
	}
	err := pc.usc.UpdateUser(ctx, input)
	if err != nil {
		presenter.RenderError(err)
		return
	}
	presenter.RenderUpdateSuccess()
}

// @Summary フォロー中のユーザーの取得
// @Tags user
// @Accept  json
// @Produce  json
// @Security Bearer
// @Success 200 {object} []model.Users
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /users/following [get]
func (pc *userController) GetFollowing(ctx *gin.Context) {
	presenter := presenter.NewUserPresenter(ctx)

	u,err := pc.usc.GetFollowing(ctx)
	if err != nil {
		presenter.RenderError(err)
		return
	}
	presenter.RenderUsers(u)
}

// @Summary フォローされてるユーザーの取得
// @Tags user
// @Accept  json
// @Produce  json
// @Security Bearer
// @Success 200 {object} []model.Users
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /users/followed [get]
func (pc *userController) GetFollowed(ctx *gin.Context) {
	presenter := presenter.NewUserPresenter(ctx)

	u,err := pc.usc.GetFollowed(ctx)
	if err != nil {
		presenter.RenderError(err)
		return
	}
	presenter.RenderUsers(u)
}