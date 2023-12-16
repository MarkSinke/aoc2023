package aoc2023

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay16InputExample(t *testing.T) {
	cells := ReadContraption("day16input_example.txt")

	FollowBeam(cells, Beam{Coord{0, 0}, Dir{1, 0}})

	sum := CountEnergized(cells)

	assert.Equal(t, 46, sum)
}

func TestDay16Input(t *testing.T) {
	cells := ReadContraption("day16input.txt")

	FollowBeam(cells, Beam{Coord{0, 0}, Dir{1, 0}})

	sum := CountEnergized(cells)

	fmt.Println("Day 16 result (star 1):", sum)
}

func TestDay16Star2InputExample(t *testing.T) {
	cells := ReadContraption("day16input_example.txt")

	FollowBeam(cells, Beam{Coord{3, 0}, Dir{0, 1}})
	sum := CountEnergized(cells)

	assert.Equal(t, 51, sum)
}

func TestDay16Star2InputExampleMax(t *testing.T) {
	cells := ReadContraption("day16input_example.txt")

	sum := FindMaxBeam(cells)

	assert.Equal(t, 51, sum)
}

func TestDay16Star2Input(t *testing.T) {
	cells := ReadContraption("day16input.txt")

	sum := FindMaxBeam(cells)

	fmt.Println("Day 16 result (star 2):", sum)
}
