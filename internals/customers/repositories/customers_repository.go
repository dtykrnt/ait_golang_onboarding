package repositories

import (
	"go-cli-skeleton/internals/customers/models"
	"log"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

type ICustomerRepository interface {
	GetAll() ([]*models.Customers, error)
	GetByID(id uint) (*models.Customers, error)
}

func NewCustomerRepository(db *gorm.DB) ICustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func (r *CustomerRepository) GetAll() ([]*models.Customers, error) {
	var customers []*models.Customers
	result := r.db.Find(&customers)
	log.Println("get all from repository")
	return customers, result.Error
}

func (r *CustomerRepository) GetByID(id uint) (*models.Customers, error) {
	var customer *models.Customers
	result := r.db.Find(&customer, id)
	return customer, result.Error
}
