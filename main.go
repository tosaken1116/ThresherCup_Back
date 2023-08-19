package main

import (
	"fmt"
	"thresher/adapter/http/router"
	"thresher/infra"
	"thresher/utils"
)

func main() {
	infra.InitDatabase()
	con := utils.LoadConfig()
	addr := fmt.Sprintf(":%s", con.Srv.Port)
	router := router.InitRouter()

	router.Run(addr)
}
