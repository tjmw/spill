package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Spill"
	app.Usage = "Sync secure notes with 1Password"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "name",
			Value: ".env.staging",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "pull",
			Usage: "Fetch secure note from 1Password",
			Flags: flags,
			Action: func(c *cli.Context) error {
				fmt.Println("Pull from 1Password")
				return nil
			},
		},
		{
			Name:  "push",
			Usage: "Push secure note to 1Password",
			Flags: flags,
			Action: func(c *cli.Context) error {
				fmt.Println("Push to 1Password")
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
