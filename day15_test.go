package aoc2023

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay15InputExample(t *testing.T) {
	instructions := ReadInitInstructions("day15input_example.txt")

	sum := SumHashes(instructions)

	assert.Equal(t, 1320, sum)
}

func TestDay15Input(t *testing.T) {
	instructions := ReadInitInstructions("day15input.txt")

	sum := SumHashes(instructions)

	fmt.Println("Day 15 result (star 1):", sum)
}

func TestDay15Star2InputExample(t *testing.T) {
	instructions := ReadInitInstructions("day15input_example.txt")

	boxes := RunInit(instructions)

	sum := SumFocusingPower(boxes)

	assert.Equal(t, 145, sum)
}

func TestDay15Star2Input(t *testing.T) {
	instructions := ReadInitInstructions("day15input.txt")

	boxes := RunInit(instructions)

	sum := SumFocusingPower(boxes)

	fmt.Println("Day 15 result (star 2):", sum)
}
