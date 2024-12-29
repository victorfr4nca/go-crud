package http

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler

func MiddlewareChain(m ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for _, middleware := range m {
			next = middleware(next)
		}
		return next
	}
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Println(r.Method, r.URL.Path, time.Since(start))
	})
}
