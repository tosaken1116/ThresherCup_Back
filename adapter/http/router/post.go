package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitPostRouter(r *gin.RouterGroup){
	// db := infra.NewPostgresConnector()
	// postRepository := repository.NewPostRepository(db.Conn)
	// postService := service.NewPostService(postRepository)
	// postUsecase := usecase.NewPostUsecase(postService)
	postGroup := r.Group("/posts")
	{
		postGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		})
	}
}