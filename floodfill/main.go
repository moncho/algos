package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var colors = []func(a ...interface{}) string{
	color.New(color.FgYellow).SprintFunc(),
	color.New(color.FgGreen).SprintFunc(),
	color.New(color.FgRed).SprintFunc(),
	color.New(color.FgBlue).SprintFunc(),
}

const block = string('\u2589')

func main() {
	rand.Seed(time.Now().UnixNano())
	rows := 1 + rand.Intn(10)
	cols := 1 + rand.Intn(20)
	grid := randomGrid(rows, cols, len(colors))

	printGrid(grid)
	color, count := maxConnectedValue(grid)

	fmt.Printf("Block %s has %d consecutive cells\n", colors[color](block), count)

}

func randomGrid(rows, cols, options int) [][]int {

	grid := make([][]int, rows)

	for x := 0; x < rows; x++ {
		grid[x] = make([]int, cols)
		for y := 0; y < cols; y++ {
			grid[x][y] = rand.Intn(options)
		}
	}
	return grid
}

func printGrid(grid [][]int) {
	fmt.Printf("grid dimensions: %d rows, %d cols\n", len(grid), len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%s ", colors[grid[i][j]](block))
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")

}
