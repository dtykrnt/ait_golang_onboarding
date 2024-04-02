package customer

import (
	"encoding/csv"
	"go-cli-skeleton/internals/customers/models"
	"log"
	"os"

	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) {
	log.Println("Seeding Customers...")

	file, err := os.Open("external/seeder-files/customers.csv")
	if err != nil {
		log.Fatal("Open Seeder Error:", file)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading CSV:", err)
		return
	}

	var customers []models.Customers
	for _, record := range records {
		name := record[0]
		email := record[1]

		// Check if the email already exists in the customers table
		var existingCustomer models.Customers
		result := db.Where("email = ?", email).First(&existingCustomer)

		if result.Error == nil {
			continue
		}

		customer := models.Customers{
			Email: email,
			Name:  name,
		}
		customers = append(customers, customer)
	}

	if len(customers) > 0 {
		err = db.Create(&customers).Error
		if err != nil {
			log.Fatal("Error Create Customer:", err)
			return
		}
	}

	log.Print("Seeding Customers Completed.")
}
