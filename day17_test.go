package aoc2023

import (
	"fmt"
	"testing"
)

func XTestDay17InputExample(t *testing.T) {
	grid := ReadHeatLossGrid("day17input_example.txt")

	mx, my := grid.GetBounds()
	result := FindLeastLossPath(grid, Coord{0, 0}, Coord{mx - 1, my - 1})

	fmt.Println("path", result.coords)
	fmt.Println(result.loss)

	PrintPath(result.coords, grid)
}

func XTestDay17InputExample2(t *testing.T) {
	grid := ReadHeatLossGrid("day17input_example2.txt")

	mx, my := grid.GetBounds()
	result := FindLeastLossPath(grid, Coord{0, 0}, Coord{mx - 1, my - 1})

	fmt.Println("path", result.coords)
	fmt.Println(result.loss)

	PrintPath(result.coords, grid)
}
