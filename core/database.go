package core

import (
	"fmt"
	"go-cli-skeleton/config"
	"go-cli-skeleton/internals/customers/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func provideDSN() string {
	cfg := config.LoadConfig()
	dbUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword)
	return dbUrl
}

func InitDatabase() *gorm.DB {

	dsn := provideDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&models.Customers{})
	return db
}
