package main

import (
	"fmt"
	"thresher/adapter/http/router"
	_ "thresher/docs"
	"thresher/infra"
	"thresher/utils/config"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title   Thresher
// @version  1.0
// @license.name Tosaken
// @description This is Thresher hackathon backend.

// @BasePath /api/v1
// @securityDefinitions.basic BearerAuth
// @in header
// @name Authorization
func main() {
	infra.InitDatabase()
	con := config.LoadConfig()
	addr := fmt.Sprintf(":%s", con.Srv.Port)
	r := router.InitRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(addr)
}
