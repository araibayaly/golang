package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/context"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		fmt.Fprintf(w, "Request took %v\n", duration)
		context.Set(r, "request-time", duration)
	})
}