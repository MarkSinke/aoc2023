package aoc2023

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay7InputExample(t *testing.T) {
	hands := ReadHands("day7input_example.txt")
	win := TotalWinnings(hands)

	assert.Equal(t, 6440, win)
}

func TestDay7Input(t *testing.T) {
	hands := ReadHands("../aoc_inputs/2023/day7input.txt")
	win := TotalWinnings(hands)

	fmt.Println("Day 7 result (star 1):", win)
}

func TestDay7Star2InputExample(t *testing.T) {
	hands := ReadHandsJoker("day7input_example.txt")
	win := TotalWinnings(hands)

	assert.Equal(t, 5905, win)
}

func TestDay7Star2Input(t *testing.T) {
	hands := ReadHandsJoker("../aoc_inputs/2023/day7input.txt")
	win := TotalWinnings(hands)

	fmt.Println("Day 7 result (star 2):", win)
}
