package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/zinderic/game-of-life/gol"
)

func main() {

	app := &cli.App{
		Name:  "game-of-life",
		Usage: "start the game of life",
		Action: func(*cli.Context) error {
			gol.Start()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
