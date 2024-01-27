package main

import (
	"fmt"
	"time"

	"github.com/zinderic/game-of-life/gol"
)

func main() {
	grid := gol.NewGrid(gol.Width, gol.Height)
	grid.InitializeRandom()

	for generation := 0; generation < 100000; generation++ {
		// Set 5 adjacent random cells to be alive
		grid.SetRandomAlive()
		fmt.Printf("Generation %d:\n", generation)
		grid.Print()
		time.Sleep(time.Second / 30)
		grid.Update()
	}
}
