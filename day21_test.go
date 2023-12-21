package aoc2023

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay21InputExample(t *testing.T) {
	garden, start := ReadGarden("day21input_example.txt")

	WalkGarden(garden, start, 6)

	count := CountVisited(garden)

	assert.Equal(t, 16, count)
}

func TestDay21Input(t *testing.T) {
	garden, start := ReadGarden("day21input.txt")

	WalkGarden(garden, start, 64)

	count := CountVisited(garden)

	fmt.Println("Day 21 result (star 1):", count)
}
