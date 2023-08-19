package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitLocationRouter(r *gin.RouterGroup){
	// db := infra.NewPostgresConnector()
	// locationRepository := repository.NewLocationRepository(db.Conn)
	// locationService := service.NewLocationService(locationRepository)
	// locationUsecase := usecase.NewLocationUsecase(locationService)
	locationGroup := r.Group("/locations")
	{
		locationGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		})
	}
}