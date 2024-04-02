package customer

import (
	"go-cli-skeleton/internals/customers/usecase"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type CustomerHandler struct {
	usecase usecase.CustomerUseCase
}

func NewCustomerHander(usecase usecase.NewCustomerUsecase) CustomerHandler {
	return CustomerHandler{
		usecase: usecase,
	}
}

func (h *CustomerHandler) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	w.Write([]byte("Received ID: " + id))
}

func (h *CustomerHandler) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Receive:"))
}
