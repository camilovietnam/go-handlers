package main

import (
	"fmt"
	"main.go/book"
	"main.go/config"
	"main.go/user"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
)

func main() {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Get("/info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	// add HTTP handlers
	book.AddHandlers(r)
	user.AddHandlers(r)

	// Signal handling
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP)

	go func() {
		for sig := range sigs {
			if sig == syscall.SIGHUP {
				fmt.Println("SIGHUP intercepted")
			}
		}
	}()

	addr := ":" + cfg.Port
	fmt.Printf("Server listening on port %s...\n", cfg.Port)
	if err := http.ListenAndServe(addr, r); err != nil {
		panic(err)
	}
}
