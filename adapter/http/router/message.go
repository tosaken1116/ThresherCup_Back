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

func InitMessageRouter(r *gin.RouterGroup) {
	db := infra.NewPostgresConnector()
	messageRepository := repository.NewMessageRepository(db.Conn)
	messageService := service.NewMessageService(messageRepository)
	messageUsecase := usecase.NewMessageUsecase(messageService)
	messageController := controller.NewMessageController(messageUsecase)

	messageGroup := r.Group("/message")
	{
		messageGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		})
		messageGroup.GET("/:id", func(c *gin.Context) { messageController.GetMessages(c) })
		messageGroup.POST("/:id", func(c *gin.Context) { messageController.CreateNewMessage(c) })
		messageGroup.GET("/unread", func(c *gin.Context) { messageController.GetUnreadMessages(c) })
	}
}
