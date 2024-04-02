package customer

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Define your API routes here using Chi
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the API!"))
	})

	// Register your RESTful handlers
	r.Get("/{id}", HandleGetByID) // For handling GET requests with an ID parameter

	return r
}
