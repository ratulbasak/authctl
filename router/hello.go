package router

import "net/http"

func hello(w http.ResponseWriter, r *http.Request) {
	data, error := helloService.GetHello()
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
