package gol

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	XSize         = 30
	YSize         = 30
	numAliveCells = 5
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
func NewGrid(XSize, YSize int) *Grid {
	grid := &Grid{
		cells: make([][]Cell, YSize),
	}

	for i := range grid.cells {
		grid.cells[i] = make([]Cell, XSize)
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

// SetRandomAlive generates 5 alive cells adjacent to each other.
func (g *Grid) SetRandomAlive() {
	rand.Seed(time.Now().UnixNano())

	startRow, startCol := rand.Intn(len(g.cells)-2), rand.Intn(len(g.cells[0])-2)
	if startCol+numAliveCells > len(g.cells[0]) {
		startCol = len(g.cells[0]) - numAliveCells
	}

	for i := 0; i < numAliveCells; i++ {
		row, col := startRow+i, startCol+i
		if row >= 0 && row < len(g.cells) && col >= 0 && col < len(g.cells[0]) {
			g.cells[row][col].alive = true
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
				fmt.Print("  ")
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

func Start() {
	grid := NewGrid(XSize, YSize)
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
