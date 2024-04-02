package cmd

import (
	"go-cli-skeleton/core"
	customer "go-cli-skeleton/internals/customers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/urfave/cli/v2"
)

func ServeAPICommand() *cli.Command {
	return &cli.Command{
		Name:  "serve-api",
		Usage: "Start the REST API server",
		Action: func(c *cli.Context) error {
			db := core.InitDatabase()
			customer.Seeder(db)
			router := chi.NewRouter()

			router.Mount("/customers", customer.SetupRouter())

			server := &http.Server{
				Addr:    ":8080", // Set your desired port
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
