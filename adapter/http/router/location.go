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

func InitLocationRouter(r *gin.RouterGroup){
	db := infra.NewPostgresConnector()
	locationRepository := repository.NewLocationRepository(db.Conn)
	locationService := service.NewLocationService(locationRepository)
	locationUsecase := usecase.NewLocationUsecase(locationService)
	homeController := controller.NewLocationController(locationUsecase)

	locationGroup := r.Group("/location")
	{
		locationGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		})
		locationGroup.POST("", func(c *gin.Context) {homeController.CreateNewLocation(c)})
	}
}