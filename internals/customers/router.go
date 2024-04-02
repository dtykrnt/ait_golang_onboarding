package customer

import (
	"go-cli-skeleton/internals/customers/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	usecase := usecase.NewCustomerUsecase()
	handler := NewCustomerHander(&usecase)

	r.Get("/", handler.HandleGetByID)

	r.Get("/{id}", handler.HandleGetAll)

	return r
}
