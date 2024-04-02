package customer

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HandleGetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	w.Write([]byte("Received ID: " + id))
}

// Add more handlers for POST, PUT, DELETE, etc., as needed
