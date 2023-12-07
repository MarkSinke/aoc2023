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
	hands := ReadHands("day7input.txt")
	win := TotalWinnings(hands)

	fmt.Println("Day 7 result (star 1):", win)
}
