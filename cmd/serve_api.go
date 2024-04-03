package cmd

import (
	"go-cli-skeleton/core"
	customer "go-cli-skeleton/internals/customers"
	"go-cli-skeleton/internals/customers/repositories"
	"go-cli-skeleton/internals/customers/usecase"
	"log"
	"net/http"

	"github.com/urfave/cli/v2"
)

func ServeCustomerAPICommand() *cli.Command {
	return &cli.Command{
		Name:  "serve-customer",
		Usage: "Start the Customers REST API server",
		Action: func(c *cli.Context) error {
			db := core.InitDatabase()
			customer.Seeder(db)
			customerRepo := repositories.NewCustomerRepository(db)
			customerUseCase := usecase.NewCustomerUsecase(customerRepo)
			customerHandler := customer.NewCustomerHander(customerUseCase)

			router := customer.SetupRouter(customerHandler)

			server := &http.Server{
				Addr:    ":8080",
				Handler: router,
			}

			log.Println("Starting REST API server on port 8080...")
			if err := server.ListenAndServe(); err != nil {
				return err
			}

			return nil
		},
	}
}

func ServeAPICommand() *cli.Command {
	return &cli.Command{
		Name:  "serve-api",
		Usage: "Start the REST API server",
		Action: func(c *cli.Context) error {
			db := core.InitDatabase()
			customer.Seeder(db)
			customerRepo := repositories.NewCustomerRepository(db)
			customerUseCase := usecase.NewCustomerUsecase(customerRepo)
			customerHandler := customer.NewCustomerHander(customerUseCase)

			router := customer.SetupRouter(customerHandler)

			server := &http.Server{
				Addr:    ":8080",
				Handler: router,
			}

			log.Println("Starting REST API server on port 8080...")
			if err := server.ListenAndServe(); err != nil {
				return err
			}

			return nil
		},
	}
}
