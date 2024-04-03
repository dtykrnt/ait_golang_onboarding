package usecase

import (
	"go-cli-skeleton/internals/customers/models"
	"go-cli-skeleton/internals/customers/repositories"
)

type CustomerUseCase struct {
	repository repositories.ICustomerRepository
}

type ICustomerUseCase interface {
	GetAll() ([]*models.Customers, error)
	GetByID(id uint) (*models.Customers, error)
}

func NewCustomerUsecase(repository repositories.ICustomerRepository) ICustomerUseCase {
	return &CustomerUseCase{
		repository: repository,
	}
}

func (u *CustomerUseCase) GetAll() ([]*models.Customers, error) {
	customers, err := u.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (u *CustomerUseCase) GetByID(id uint) (*models.Customers, error) {
	customer, err := u.repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
