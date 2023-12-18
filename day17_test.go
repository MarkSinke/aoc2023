package aoc2023

import (
	"fmt"
	"testing"
)

func TestDay17InputExample(t *testing.T) {
	grid := ReadHeatLossGrid("day17input_example.txt")

	steps, loss := FindLeastLossPath(grid)

	fmt.Println("path", steps)
	fmt.Println(loss)

	PrintPath(steps, grid)
}
