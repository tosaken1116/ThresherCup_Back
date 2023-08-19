package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitHomeRouter(r *gin.RouterGroup){
	// db := infra.NewPostgresConnector()
	// homeRepository := repository.NewHomeRepository(db.Conn)
	// homeService := service.NewHomeService(homeRepository)
	// homeUsecase := usecase.NewHomeUsecase(homeService)
	homeGroup := r.Group("/home")
	{
		homeGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		})
	}
}