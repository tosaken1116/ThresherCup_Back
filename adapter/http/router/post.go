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

func InitPostRouter(r *gin.RouterGroup){
	db := infra.NewPostgresConnector()
	postRepository := repository.NewPostRepository(db.Conn)
	postService := service.NewPostService(postRepository)
	postUsecase := usecase.NewPostUsecase(postService)
	postController := controller.NewPostController(postUsecase)
	postGroup := r.Group("/posts")
	{
		postGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		})
		postGroup.GET("/:id",func(c *gin.Context) {postController.GetPostById(c)})
		postGroup.POST("",func(c *gin.Context) {postController.CreateNewPost(c)})
	}
}