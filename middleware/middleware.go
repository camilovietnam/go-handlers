package middleware

import (
	"log"
	"net/http"

	"github.com/google/uuid"
)

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-ID")
		if id == "" {
			id = uuid.NewString()
		}
		log.Printf("Request ID: %s %s %s", id, r.Method, r.URL.Path)
		r.Header.Set("X-Request-ID", id) // optional: propagate internally
		next.ServeHTTP(w, r)
	})
}
