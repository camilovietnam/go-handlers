package user

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func AddHandlers(r chi.Router) {
	r.Get("/user", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "/user to be implemented")
	})

	r.Get("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "/user/{id} to be implemented")
	})
}
