package customer

import (
	"encoding/json"
	"go-cli-skeleton/internals/customers/usecase"
	"net/http"
)

type CustomerHandler struct {
	usecase usecase.ICustomerUseCase
}

func NewCustomerHander(usecase usecase.ICustomerUseCase) CustomerHandler {
	return CustomerHandler{
		usecase: usecase,
	}
}

func (h *CustomerHandler) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	// id := chi.URLParam(r, "id")
	// customerId, err := strconv.ParseUint(id, 10, 64)
	// if err != nil {
	// 	return
	// }
	// result, err := h.usecase.GetByID(uint(customerId))
	// if err != nil {
	// 	return
	// }
	// resultJson, err := json.Marshal(result)
	// if err != nil {
	// 	return
	// }
	// w.Write(resultJson)
	return
}

func (h *CustomerHandler) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	result, err := h.usecase.GetAll()
	if err != nil {
		return
	}
	resultJson, err := json.Marshal(result)
	if err != nil {
		return
	}
	w.Write(resultJson)
}
