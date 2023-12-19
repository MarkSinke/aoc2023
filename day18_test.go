package aoc2023

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay18InputExample(t *testing.T) {
	steps := ReadDigPlan("day18input_example.txt")

	tiles := ExecuteDigPlan(steps)

	Fill(tiles)

	sum := CountInOrPipe(tiles)

	assert.Equal(t, 62, sum)
}

func TestDay18Input(t *testing.T) {
	steps := ReadDigPlan("day18input.txt")

	tiles := ExecuteDigPlan(steps)

	Fill(tiles)

	sum := CountInOrPipe(tiles)

	fmt.Println("Day 18 result (star 1):", sum)
}

func XTestDay18Star2InputExample(t *testing.T) {
	steps := ReadDigPlan2("day18input_example.txt")
	fmt.Println(steps)
	fmt.Println(getExtent(steps))

	// tiles := ExecuteDigPlan(steps)

	// Fill(tiles)

	// sum := CountInOrPipe(tiles)

	// assert.Equal(t, 952408144115, sum)
}
