package core

import (
	"net/http"
)

func ExampleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Implement your middleware logic here
		// For example, you might perform actions before passing to the next handler

		// Pass on to the next middleware/handler
		next.ServeHTTP(w, r)
	})
}
