package main

import (
	"fmt"
	"thresher/adapter/http"
	"thresher/utils"
)

func main() {
	con := utils.LoadConfig()
	addr := fmt.Sprintf(":%s", con.Srv.Port)
	router := http.InitRouter()

	router.Run(addr)
}
