package main

import (
	"fmt"
	"net/http"

	"simply-go/config"
	"simply-go/router"
	"simply-go/service"
)

func main() {
	cfg := config.LoadConfig()

	helloService := &service.HelloService{}
	mux := router.SetupRouter(helloService)
	fmt.Println("Server started on ", cfg.Port)
	err := http.ListenAndServe(cfg.Port, mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
