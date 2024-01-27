package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 10
	height = 10
)

// Cell represents a single cell in the Game of Life.
type Cell struct {
	alive bool
}

// Grid represents the game board.
type Grid struct {
	cells [][]Cell
}

// NewGrid creates a new empty grid.
func NewGrid(width, height int) *Grid {
	grid := &Grid{
		cells: make([][]Cell, height),
	}

	for i := range grid.cells {
		grid.cells[i] = make([]Cell, width)
	}

	return grid
}

// InitializeRandom fills the grid with random alive or dead cells.
func (g *Grid) InitializeRandom() {
	rand.Seed(time.Now().UnixNano())

	for i := range g.cells {
		for j := range g.cells[i] {
			g.cells[i][j].alive = rand.Intn(2) == 1
		}
	}
}

// Print displays the current state of the grid.
func (g *Grid) Print() {
	for i := range g.cells {
		for j := range g.cells[i] {
			if g.cells[i][j].alive {
				fmt.Print("■ ")
			} else {
				fmt.Print("□ ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// Update evolves the grid to the next generation based on the rules of the Game of Life.
func (g *Grid) Update() {
	newGrid := NewGrid(len(g.cells[0]), len(g.cells))

	for i := range g.cells {
		for j := range g.cells[i] {
			neighbors := g.countAliveNeighbors(i, j)
			newGrid.cells[i][j].alive = g.shouldCellLive(g.cells[i][j].alive, neighbors)
		}
	}

	g.cells = newGrid.cells
}

// countAliveNeighbors returns the number of alive neighbors for a given cell.
func (g *Grid) countAliveNeighbors(row, col int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			r, c := row+i, col+j
			if r >= 0 && r < len(g.cells) && c >= 0 && c < len(g.cells[0]) {
				if g.cells[r][c].alive {
					count++
				}
			}
		}
	}
	return count
}

// shouldCellLive determines whether a cell should be alive in the next generation.
func (g *Grid) shouldCellLive(currentState bool, neighbors int) bool {
	if currentState {
		return neighbors == 2 || neighbors == 3
	}
	return neighbors == 3
}

func main() {
	grid := NewGrid(width, height)
	grid.InitializeRandom()

	for generation := 0; generation < 10; generation++ {
		fmt.Printf("Generation %d:\n", generation)
		grid.Print()
		time.Sleep(time.Second)
		grid.Update()
	}
}
