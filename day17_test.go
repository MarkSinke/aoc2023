package aoc2023

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay17InputExample(t *testing.T) {
	grid := ReadHeatLossGrid("day17input_example.txt")

	_, cost := FindLeastLossPath(grid, 1, 3)

	assert.Equal(t, 102, cost)
}

func TestDay17InputExample2(t *testing.T) {
	grid := ReadHeatLossGrid("day17input_example2.txt")

	_, cost := FindLeastLossPath(grid, 1, 3)

	assert.Equal(t, 21, cost)
}

func TestDay17Input(t *testing.T) {
	grid := ReadHeatLossGrid("../aoc_inputs/2023/day17input.txt")

	_, cost := FindLeastLossPath(grid, 1, 3)

	fmt.Println("Day 17 result (star 1):", cost)
}

func TestDay17Star2InputExample(t *testing.T) {
	grid := ReadHeatLossGrid("day17input_example.txt")

	_, cost := FindLeastLossPath(grid, 4, 10)

	assert.Equal(t, 94, cost)
}

func TestDay17Star2Input(t *testing.T) {
	grid := ReadHeatLossGrid("../aoc_inputs/2023/day17input.txt")

	_, cost := FindLeastLossPath(grid, 4, 10)

	fmt.Println("Day 17 result (star 2):", cost)
}
