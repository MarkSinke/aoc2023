package aoc2023

import (
	"fmt"
	"testing"
)

func TestDay2Example(t *testing.T) {
	gameResults := ReadGameResults("day2input_example.txt")

	sum := sumIdsPlayable(gameResults)

	if sum != 8 {
		t.Errorf("expected 8, but got %v", sum)
	}
}

func TestDay2Input(t *testing.T) {
	gameResults := ReadGameResults("day2input.txt")

	sum := sumIdsPlayable(gameResults)

	fmt.Printf("Day 2 result (star 1): %d\n", sum)
}

func sumIdsPlayable(gameResults []GameResult) int {
	sum := 0
	for _, res := range gameResults {
		if CanPlayWith(12, 13, 14, res) {
			sum += res.id
		}
	}
	return sum
}
