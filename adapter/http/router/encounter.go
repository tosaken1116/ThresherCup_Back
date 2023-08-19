package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitEncounterRouter(r *gin.RouterGroup){
	// db := infra.NewPostgresConnector()
	// encounterRepository := repository.NewEncounterRepository(db.Conn)
	// encounterService := service.NewEncounterService(encounterRepository)
	// encounterUsecase := usecase.NewEncounterUsecase(encounterService)
	encounterGroup := r.Group("/encounters")
	{
		encounterGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		})
	}
}