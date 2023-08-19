package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/health", func(c *gin.Context) {
				// 必要に応じて追加のチェックを行うことができます
				// ここでは単純に200 OKを返します
				c.JSON(http.StatusOK, gin.H{
					"status": "OK",
				})
			})

		}
	}
	return r

}
