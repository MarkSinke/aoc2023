package aoc2023

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay24InputExample(t *testing.T) {
	hails := ReadHail("day24input_example.txt")

	count := CountIntersectionPairs(hails, 7, 27)

	assert.Equal(t, 2, count)
}

func TestDay24Input(t *testing.T) {
	hails := ReadHail("day24input.txt")

	count := CountIntersectionPairs(hails, 2e14, 4e14)

	fmt.Println("Day 24 result (star 1):", count)
}

func TestDay24Star2InputExample(t *testing.T) {
	hails := ReadHail("day24input_example.txt")

	h := ComputeIntersectingHailStone(hails, 7, 27)

	assert.Equal(t, Coord3F{24, 13, 10}, h.pos)
	assert.Equal(t, Dir3F{-3, 1, 2}, h.dir)
}

func TestDay24Star2Input(t *testing.T) {
	hails := ReadHail("day24input.txt")

	h := ComputeIntersectingHailStone(hails, 2e14, 4e14)

	fmt.Println("h", h)
	fmt.Println("Day 24 result (star 2):", int(math.Round(h.pos.x+h.pos.y+h.pos.z)))
	assert.Equal(t, Coord3F{24, 13, 10}, h.pos)
	assert.Equal(t, Dir3F{-3, 1, 2}, h.dir)
}
