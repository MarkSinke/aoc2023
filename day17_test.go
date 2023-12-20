package aoc2023

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay17InputExample(t *testing.T) {
	grid := ReadHeatLossGrid("day17input_example.txt")

	mx, my := grid.GetBounds()
	_, cost := FindLeastLossPath(grid, Coord{0, 0}, Coord{mx - 1, my - 1}, 1, 3)

	assert.Equal(t, 102, cost)
}

func TestDay17InputExample2(t *testing.T) {
	grid := ReadHeatLossGrid("day17input_example2.txt")

	mx, my := grid.GetBounds()
	_, cost := FindLeastLossPath(grid, Coord{0, 0}, Coord{mx - 1, my - 1}, 1, 3)

	assert.Equal(t, 21, cost)
}

func TestDay17Input(t *testing.T) {
	grid := ReadHeatLossGrid("day17input.txt")

	mx, my := grid.GetBounds()
	_, cost := FindLeastLossPath(grid, Coord{0, 0}, Coord{mx - 1, my - 1}, 1, 3)

	fmt.Println("Day 17 result (star 1):", cost)
}

func XTestDay17Star2InputExample(t *testing.T) {
	grid := ReadHeatLossGrid("day17input_example.txt")

	mx, my := grid.GetBounds()
	path, cost := FindLeastLossPath(grid, Coord{0, 0}, Coord{mx - 1, my - 1}, 4, 10)

	fmt.Println(path)

	assert.Equal(t, 94, cost)
}

func XTestDay17Star2Input(t *testing.T) {
	grid := ReadHeatLossGrid("day17input.txt")

	mx, my := grid.GetBounds()
	_, cost := FindLeastLossPath(grid, Coord{0, 0}, Coord{mx - 1, my - 1}, 4, 10)

	fmt.Println("Day 17 result (star 2):", cost)
}
