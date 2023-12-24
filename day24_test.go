package aoc2023

import (
	"fmt"
	"math/big"
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

	assert.Equal(t, Coord3F{big.NewInt(24), big.NewInt(13), big.NewInt(10)}, h.pos)
	assert.Equal(t, Dir3F{big.NewInt(-3), big.NewInt(1), big.NewInt(2)}, h.dir)
}

func TestDay24Star2Input(t *testing.T) {
	hails := ReadHail("day24input.txt")

	h := ComputeIntersectingHailStone(hails, 2e14, 4e14)

	var sum big.Int
	sum.Add(h.pos.x, h.pos.y)
	sum.Add(&sum, h.pos.z)

	fmt.Println("Day 24 result (star 2):", sum.String())
	assert.Equal(t, big.NewInt(1004774995964534), &sum)
}
