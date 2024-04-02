package usecase

import "go-cli-skeleton/internals/customers/models"

type CustomerUseCase struct {
	// Define dependencies as needed
}

func NewCustomerUsecase() *CustomerUseCase {
	return &CustomerUseCase{}
}

func (uc *CustomerUseCase) GetAllItems() ([]models.Customers, error) {
	// Implement business logic to fetch all items
	return []models.Customers{}, nil
}

// Implement other usecase methods for CRUD operations
