package aoc2023

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay9InputExample(t *testing.T) {
	readings := ReadOasis("day9input_example.txt")

	res := SumOfPredictions(readings)

	assert.Equal(t, 114, res)
}

func TestDay9Input(t *testing.T) {
	readings := ReadOasis("day9input.txt")

	res := SumOfPredictions(readings)

	fmt.Println("Day 9 result (star 1):", res)
}
