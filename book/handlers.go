package book

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func AddHandlers(r chi.Router) {
	r.Get("/book", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "/book to be implemented")
	})

	r.Get("/book/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "/book/{id} to be implemented")
	})
}
