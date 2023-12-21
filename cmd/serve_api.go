package cmd

import (
	"go-cli-skeleton/internals/yourmodule"
	"log"
	"net/http"

	"github.com/urfave/cli/v2"
)

func ServeAPICommand() *cli.Command {
	return &cli.Command{
		Name:  "serve-api",
		Usage: "Start the REST API server",
		Action: func(c *cli.Context) error {
			router := yourmodule.SetupRouter() // Set up your Chi router in handler.SetupRouter()

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
