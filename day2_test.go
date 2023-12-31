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
	gameResults := ReadGameResults("../aoc_inputs/2023/day2input.txt")

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

func TestDay2ExampleStar2(t *testing.T) {
	gameResults := ReadGameResults("day2input_example.txt")

	sum := sumMinPowers(gameResults)

	if sum != 2286 {
		t.Errorf("expected 8, but got %v", sum)
	}
}

func TestDay2Star2(t *testing.T) {
	gameResults := ReadGameResults("../aoc_inputs/2023/day2input.txt")

	sum := sumMinPowers(gameResults)

	fmt.Printf("Day 2 result (star 2): %d\n", sum)
}

func sumMinPowers(gameResults []GameResult) int {
	sum := 0
	for _, res := range gameResults {
		sum += MinPower(res)
	}
	return sum
}
