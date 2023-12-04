package aoc2023

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay4InputExample(t *testing.T) {
	cards := ReadGameCards(("day4input_example.txt"))
	total := sumOfScores(cards)

	assert.Equal(t, total, 13)
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
