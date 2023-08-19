package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.RouterGroup){
	// db := infra.NewPostgresConnector()
	// userRepository := repository.NewUserRepository(db.Conn)
	// userService := service.NewUserService(userRepository)
	// userUsecase := usecase.NewUserUsecase(userService)
	userGroup := r.Group("/users")
	{
		userGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		})
	}
}