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
	tiles := ReadTiles("../aoc_inputs/2023/day10input.txt")
	loopLen := FindLoopLen(tiles)
	steps := loopLen / 2

	fmt.Println("Day 10 result (star 1):", steps)
}

func TestDay10Star2InputExample2(t *testing.T) {
	tiles := ReadTiles("day10input_example3.txt")
	main := CopyMainLoop(tiles)
	exploded := ExplodeLoop(main)
	Fill(exploded)
	imploded := ImplodeLoop(exploded)

	in := CountIn(imploded)
	assert.Equal(t, 8, in)
}

func TestDay10Star2Input(t *testing.T) {
	tiles := ReadTiles("../aoc_inputs/2023/day10input.txt")
	main := CopyMainLoop(tiles)
	// we need to explode and implode to make sure every point outside can be reached via a path
	// exploding makes sure such paths exist
	exploded := ExplodeLoop(main)
	Fill(exploded)
	imploded := ImplodeLoop(exploded)

	in := CountIn(imploded)
	fmt.Println("Day 10 result (star 2):", in)
}
