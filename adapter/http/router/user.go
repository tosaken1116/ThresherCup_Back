package router

import (
	"net/http"
	"thresher/adapter/http/controller"
	"thresher/domain/repository"
	"thresher/domain/service"
	"thresher/infra"
	"thresher/usecase"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.RouterGroup) {
	db := infra.NewPostgresConnector()
	userRepository := repository.NewUserRepository(db.Conn)
	userService := service.NewUserService(userRepository)
	userUsecase := usecase.NewUserUsecase(userService)
	userController := controller.NewUserController(userUsecase)
	userGroup := r.Group("/users")
	{
		userGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		})
		userGroup.PUT("", func(c *gin.Context) {
			userController.UpdateUser(c)
		})
		userGroup.GET("/following", func(c *gin.Context) { userController.GetFollowing(c) })
		userGroup.GET("/followed", func(c *gin.Context) { userController.GetFollowed(c) })
		userGroup.POST("/follow/:id", func(c *gin.Context) { userController.NewFollow(c) })
		userGroup.DELETE("/follow/:id", func(c *gin.Context) { userController.DeleteFollow(c) })
	}
}
