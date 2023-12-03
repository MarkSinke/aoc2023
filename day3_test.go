package aoc2023

import (
	"fmt"
	"testing"
)

func TestDay3InputExample(t *testing.T) {
	engine := ReadFile("day3input_example.txt")
	parts := GetPartNumbers(engine)
	sum := sum(parts)

	if sum != 4361 {
		t.Errorf("expected 4361, but got %d", sum)
	}
}

func TestDay3Input(t *testing.T) {
	engine := ReadFile("day3input.txt")
	parts := GetPartNumbers(engine)
	sum := sum(parts)

	fmt.Printf("Day 3 result (star 1): %d", sum)
}
