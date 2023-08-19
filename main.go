package main

import (
	"fmt"
	"thresher/adapter/http"
	"thresher/infra"
	"thresher/utils"
)

func main() {
	infra.InitDatabase()
	con := utils.LoadConfig()
	addr := fmt.Sprintf(":%s", con.Srv.Port)
	router := http.InitRouter()

	router.Run(addr)
}
