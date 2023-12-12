package aoc2023

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay12Matching(t *testing.T) {
	records := ReadSpringRecords("day12input_example_no_unknowns.txt")

	for _, record := range records {
		assert.True(t, record.Matches(), "record %v is supposed to match", record)
	}
}

func TestDay12InputExample(t *testing.T) {
	records := ReadSpringRecords("day12input_example.txt")

	sum := 0
	for _, record := range records {
		matches := record.PossibleMatches()
		sum = sum + matches
	}

	assert.Equal(t, 21, sum)
}

func TestDay12Input(t *testing.T) {
	records := ReadSpringRecords("day12input.txt")

	sum := 0
	for _, record := range records {
		sum = sum + record.PossibleMatches()
	}

	fmt.Println("Day 12 result (star 1):", sum)
}
