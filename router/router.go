package router

import (
	"net/http"

	"simply-go/service"
)

func SetupRouter(helloService *service.HelloService) *http.ServeMux {
	mux := http.NewServeMux()

	InitializeRoutes(mux, helloService)
	return mux
}
