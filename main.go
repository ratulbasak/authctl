package main

import (
	"fmt"
	"net/http"

	"simply-go/config"
	"simply-go/router"
	"simply-go/service"
)

// type hello struct {
// 	Message string `json:"message"`
// }

func main() {
	cfg := config.LoadConfig()

	helloService := &service.HelloService{}
	mux := router.SetupRouter(helloService)
	fmt.Println("Server started on localhost:8080")
	err := http.ListenAndServe(cfg.Port, mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
