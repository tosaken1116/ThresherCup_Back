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

func InitEncounterRouter(r *gin.RouterGroup){
	db := infra.NewPostgresConnector()
	encounterRepository := repository.NewEncounterRepository(db.Conn)
	encounterService := service.NewEncounterService(encounterRepository)
	encounterUsecase := usecase.NewEncounterUsecase(encounterService)
	encounterController := controller.NewEncounterController(encounterUsecase)

	encounterGroup := r.Group("/encounter")
	{
		encounterGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		})
		encounterGroup.GET("",func(c *gin.Context){
			encounterController.GetEncounter(c)
		})
	}
}