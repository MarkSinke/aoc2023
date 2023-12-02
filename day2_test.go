package aoc2023

import (
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

	if sum != 8 {
		t.Errorf("expected 8, but got %v", sum)
	}
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
