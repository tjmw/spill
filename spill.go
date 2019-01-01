package main

import (
	"log"
	"os"

	"github.com/tjmw/spill/cli"
)

func main() {
	app := cli.Cli()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
