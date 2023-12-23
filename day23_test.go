package aoc2023

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay23InputExample(t *testing.T) {
	maze := ReadMaze("day23input_example.txt")

	res := FindLongestPath(maze)

	assert.Equal(t, 94, res)
}

func TestDay23Input(t *testing.T) {
	maze := ReadMaze("day23input.txt")

	res := FindLongestPath(maze)

	fmt.Println("Day 23 result (star 1):", res)
}

func TestDay23Star2InputExample(t *testing.T) {
	maze := ReadMaze("day23input_example.txt")

	FlattenSlopes(maze)

	res := FindLongestPath(maze)

	assert.Equal(t, 154, res)
}

func TestDay23Star2Input(t *testing.T) {
	maze := ReadMaze("day23input.txt")

	FlattenSlopes(maze)

	res := FindLongestPath(maze)

	fmt.Println("Day 23 result (star 2):", res)
}
