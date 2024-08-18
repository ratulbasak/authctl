package router

import (
	"net/http"

	"simply-go/service"
)

func SetupRouter(helloService *service.HelloService) http.Handler {
	mux := http.NewServeMux()

	InitializeRoutes(mux, helloService)
	loggedMux := LoggingMiddleware(mux)

	return loggedMux
}
