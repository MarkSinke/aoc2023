package aoc2023

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay11InputExample(t *testing.T) {
	coords := ReadGalaxies("day11input_example.txt")
	coordsNew := ExpandGalaxy(coords)
	dists := ComputeDistances(coordsNew)
	sum := Sum(dists)

	assert.Equal(t, 374, sum)
}

func TestDay11Input(t *testing.T) {
	coords := ReadGalaxies("day11input.txt")
	coordsNew := ExpandGalaxy(coords)
	dists := ComputeDistances(coordsNew)
	sum := Sum(dists)

	fmt.Println("Day 11 result (star 1):", sum)
}
