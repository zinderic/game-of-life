package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/zinderic/game-of-life/gol"
)

func main() {

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "xSize",
				Aliases:     []string{"x"},
				Value:       30,
				Usage:       "set xSize of the grid",
				DefaultText: "30",
				Destination: &gol.XSize,
			},
			&cli.IntFlag{
				Name:        "ySize",
				Aliases:     []string{"y"},
				Value:       30,
				Usage:       "set ySize of the grid",
				DefaultText: "30",
				Destination: &gol.YSize,
			},
			&cli.BoolFlag{
				Name:        "term-size",
				Aliases:     []string{"t"},
				Value:       false,
				Usage:       "set grid size to terminal size",
				DefaultText: "false",
				Destination: &gol.TermSize,
			},
		},
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
