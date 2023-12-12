package aoc2023

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay12InputExample(t *testing.T) {
	records := ReadSpringRecords("day12input_example.txt")

	sum := 0
	for _, record := range records {
		// fmt.Print("matching", record)
		matches := record.PossibleMatches()
		// fmt.Println(" ->", matches)
		sum = sum + matches
	}

	assert.Equal(t, 21, sum)
}

func TestDay12Input(t *testing.T) {
	records := ReadSpringRecords("day12input.txt")

	sum := 0
	for _, record := range records {
		matches := record.PossibleMatches()
		sum = sum + matches
	}

	fmt.Println("Day 12 result (star 1):", sum)
}

func TestDay12Unfold(t *testing.T) {
	record := SpringRecord{".#", []int{1}}

	unfolded := record.Unfold()

	assert.Equal(t, ".#?.#?.#?.#?.#", unfolded.record)
	assert.Equal(t, []int{1, 1, 1, 1, 1}, unfolded.counts)
}

func TestDay12Star2InputExample(t *testing.T) {
	records := ReadSpringRecords("day12input_example.txt")

	sum := 0
	for _, record := range records {
		matches := record.Unfold().PossibleMatches()
		sum = sum + matches
	}

	assert.Equal(t, 525152, sum)
}

func TestDay12Star2Input(t *testing.T) {
	records := ReadSpringRecords("day12input.txt")

	sum := 0
	for _, record := range records {
		fmt.Print("matching", record)
		fmt.Printf(" (min %v) ", len(record.record)-minLengthForTail(record.counts))
		matches := 0
		// matches := record.Unfold().PossibleMatches2()
		fmt.Println(" ->", matches)
		// sum = sum + matches
	}

	fmt.Println("Day 12 result (star 2):", sum)
}

func TestDay12Star2HardRecord(t *testing.T) {
	record := SpringRecord{"..?????.????", []int{1, 1, 1, 2}}

	assert.Equal(t, 9, record.PossibleMatches())
	assert.Equal(t, 7811529, record.Unfold().PossibleMatches())
}

func TestDay12Star2HardRecord2(t *testing.T) {
	record := SpringRecord{"??????.????", []int{1, 1}}

	assert.Equal(t, 37, record.PossibleMatches())
	assert.Equal(t, 3985514935, record.Unfold().PossibleMatches())
}

func TestDay12Star2HardRecord3(t *testing.T) {
	record := SpringRecord{"?????????????##", []int{3, 4, 3}}

	assert.Equal(t, 10, record.PossibleMatches())
	// assert.Equal(t, 1846252, record.Unfold().PossibleMatches2())
}
