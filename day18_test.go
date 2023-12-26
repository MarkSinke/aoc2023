package aoc2023

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay18InputExample(t *testing.T) {
	steps := ReadDigPlan("day18input_example.txt")

	coords := ComputeCoords(steps)
	area := ComputePolygonArea(coords)

	assert.Equal(t, 62, area)
}

func TestDay18Input(t *testing.T) {
	steps := ReadDigPlan("../aoc_inputs/2023/day18input.txt")

	coords := ComputeCoords(steps)
	area := ComputePolygonArea(coords)

	fmt.Println("Day 18 result (star 1):", area)

}

func TestDay18Star2InputExample(t *testing.T) {
	steps := ReadDigPlan2("day18input_example.txt")

	coords := ComputeCoords(steps)
	area := ComputePolygonArea(coords)

	assert.Equal(t, 952408144115, area)
}

func TestDay18Star2Input(t *testing.T) {
	steps := ReadDigPlan2("../aoc_inputs/2023/day18input.txt")

	coords := ComputeCoords(steps)
	area := ComputePolygonArea(coords)

	fmt.Println("Day 18 result (star 2):", area)

}
