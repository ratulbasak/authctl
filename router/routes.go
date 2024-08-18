package router

import (
	"net/http"

	"simply-go/service"
)

var helloService *service.HelloService

func InitializeRoutes(mux *http.ServeMux, helloSvc *service.HelloService) {
	helloService = helloSvc

	mux.HandleFunc("/", hello)
}
