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

func InitHomeRouter(r *gin.RouterGroup) {
	db := infra.NewPostgresConnector()
	homeRepository := repository.NewHomeRepository(db.Conn)
	homeService := service.NewHomeService(homeRepository)
	homeUsecase := usecase.NewHomeUsecase(homeService)
	homeController := controller.NewHomeController(homeUsecase)

	homeGroup := r.Group("/home")
	{
		homeGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		})
		homeGroup.POST("", func(c *gin.Context) { homeController.CreateNewHome(c) })
		homeGroup.GET("", func(c *gin.Context) { homeController.GetMyHome(c) })
	}
}
