package aoc2023

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay10InputExample1(t *testing.T) {
	tiles := ReadTiles("day10input_example.txt")
	loopLen := FindLoopLen(tiles)
	steps := loopLen / 2
	assert.Equal(t, 4, steps)
}

func TestDay10InputExample2(t *testing.T) {
	tiles := ReadTiles("day10input_example2.txt")
	loopLen := FindLoopLen(tiles)
	steps := loopLen / 2
	assert.Equal(t, 8, steps)
}

func TestDay10Input(t *testing.T) {
	tiles := ReadTiles("day10input.txt")
	loopLen := FindLoopLen(tiles)
	steps := loopLen / 2

	fmt.Println("Day 10 result (star 1):", steps)
}
