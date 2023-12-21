package main

import (
	"go-cli-skeleton/cmd"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "YourApp",
		Usage: "Your application description",
		Commands: []*cli.Command{
			cmd.ServeAPICommand(),
			{
				Name:  "start-worker",
				Usage: "Start a worker",
				Action: func(c *cli.Context) error {
					// Your code to start a worker
					log.Println("Starting a worker...")
					return nil
				},
			},
			{
				Name:  "process-data",
				Usage: "Process data",
				Action: func(c *cli.Context) error {
					// Your code to process data
					log.Println("Processing data...")
					return nil
				},
			},
			// Add other commands if needed
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
