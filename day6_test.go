package aoc2023

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay6InputExample(t *testing.T) {
	games := ReadRaceGames("day6input_example.txt")
	result := WinCounts(games)

	assert.Equal(t, 288, result)
}

func TestDay6Inpu(t *testing.T) {
	games := ReadRaceGames("../aoc_inputs/2023/day6input.txt")
	result := WinCounts(games)

	fmt.Println("Day 6 result (star 1):", result)
}

func TestDay6Star2InputExample(t *testing.T) {
	game := ReadRaceGame("day6input_example.txt")
	result := WinCount(game)

	assert.Equal(t, 71503, result)
}

func TestDay6Star2Inpu(t *testing.T) {
	game := ReadRaceGame("../aoc_inputs/2023/day6input.txt")
	result := WinCount(game)

	fmt.Println("Day 6 result (star 2):", result)
}
