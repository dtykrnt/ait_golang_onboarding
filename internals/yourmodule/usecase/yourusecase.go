package usecase

import (
	"go-cli-skeleton/internals/yourmodule/models"
)

type ItemUsecase struct {
	// Define dependencies as needed
}

func NewItemUsecase() *ItemUsecase {
	// Initialize and return usecase with dependencies injected
	return &ItemUsecase{}
}

func (uc *ItemUsecase) GetAllItems() ([]models.Item, error) {
	// Implement business logic to fetch all items
	return []models.Item{}, nil
}

// Implement other usecase methods for CRUD operations
