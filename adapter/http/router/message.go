package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitMessageRouter(r *gin.RouterGroup) {
	// db := infra.NewPostgresConnector()
	// messageRepository := repository.NewMessageRepository(db.Conn)
	// messageService := service.NewMessageService(messageRepository)
	// messageUsecase := usecase.NewMessageUsecase(messageService)
	// messageController := controller.NewMessageController(messageUsecase)

	messageGroup := r.Group("/message")
	{
		messageGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		})
	}
}
