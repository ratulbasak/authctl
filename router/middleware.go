package router

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs the details of the incoming HTTP request and its response status.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Custom ResponseWriter to capture status code
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		log.Printf("%s %s %s %d %s", r.Method, r.RequestURI, r.RemoteAddr, rw.statusCode, time.Since(start))
	})
}

// responseWriter wraps the standard http.ResponseWriter to capture the status code.
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
