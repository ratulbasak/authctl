package main

import (
	"fmt"
	"net/http"

	"simply-go/config"
	"simply-go/service"
)

// type hello struct {
// 	Message string `json:"message"`
// }

func main() {
	fmt.Println("started!")
	cfg := config.LoadConfig()
	// message := hello{
	// 	Message: "hello world",
	// }
	helloService := &service.HelloService{}
	data, error := helloService.GetHello()
	if error != nil {
		fmt.Println("Internal Server Error")
		return
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})

	err := http.ListenAndServe(cfg.Port, mux)
	if err != nil {
		fmt.Println(err.Error())
	}
}
