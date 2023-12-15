package aoc2023

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay14InputExample(t *testing.T) {
	rocks := ReadRockPositions("day14input_example.txt")

	TiltNorth(rocks)
	weight := GetNorthWeight(rocks)

	assert.Equal(t, 136, weight)
}

func TestDay14Input(t *testing.T) {
	rocks := ReadRockPositions("day14input.txt")

	TiltNorth(rocks)
	weight := GetNorthWeight(rocks)

	fmt.Println("Day 14 result (star 1):", weight)
}

func TestDay14Star2InputExample(t *testing.T) {
	rocks := ReadRockPositions("day14input_example.txt")

	SpinOneBillionTimes(rocks)
	weight := GetNorthWeight(rocks)

	assert.Equal(t, 64, weight)
}

func TestDay14Star2Input(t *testing.T) {
	rocks := ReadRockPositions("day14input.txt")

	SpinOneBillionTimes(rocks)
	weight := GetNorthWeight(rocks)

	fmt.Println("Day 14 result (star 2):", weight)
}
