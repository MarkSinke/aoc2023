package aoc2023

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay4InputExample(t *testing.T) {
	cards := ReadGameCards(("day4input_example.txt"))
	total := sumOfScores(cards)

	assert.Equal(t, 13, total)
}

func TestDay4Input(t *testing.T) {
	cards := ReadGameCards(("day4input.txt"))
	total := sumOfScores(cards)

	fmt.Printf("Day 4 result (star 1): %v\n", total)
}

func sumOfScores(cards []GameCard) int {
	sum := 0
	for _, card := range cards {
		sum += card.Score()
	}
	return sum
}

func TestDay4Star2InputExample(t *testing.T) {
	cards := ReadGameCards(("day4input_example.txt"))
	counts := ComputeCardCounts(cards)
	total := Sum(counts)

	assert.Equal(t, 30, total)
}

func TestDay4Star2Input(t *testing.T) {
	cards := ReadGameCards(("day4input.txt"))
	counts := ComputeCardCounts(cards)
	total := Sum(counts)

	fmt.Printf("Day 4 result (star 2): %v\n", total)
}
