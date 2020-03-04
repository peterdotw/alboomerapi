package utils

import (
	"log"
	"net/http"
	"time"
)

//RequestLoggerMiddleware - Simple logger middleware
func RequestLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		clientIP := r.RemoteAddr

		next.ServeHTTP(w, r)

		log.Printf(
			"%s\t\t%s\t\t%s\t\t%v",
			r.Method,
			r.RequestURI,
			clientIP,
			time.Since(start),
		)
	})
}
