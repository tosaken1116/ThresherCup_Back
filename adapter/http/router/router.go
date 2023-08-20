package router

import (
	"net/http"
	"thresher/utils/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Auth())
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/health", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"status": "OK",
				})
			})
			InitUserRouter(v1)
			InitHomeRouter(v1)
			InitPostRouter(v1)
			InitLocationRouter(v1)
			InitEncounterRouter(v1)
		}
	}
	return r

}
