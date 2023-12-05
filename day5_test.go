package aoc2023

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Min(ints []int) int {
	res := math.MaxInt
	for _, i := range ints {
		res = min(res, i)
	}
	return res
}

func MinOfRanges(ranges []Range) int {
	res := math.MaxInt
	for _, r := range ranges {
		res = min(res, r.start)
	}
	return res
}

func TestDay5InputExammple(t *testing.T) {
	seeds, maps := ReadSeedMaps("day5input_example.txt")
	res := MapSeedsThroughMaps(seeds, maps)
	m := Min(res)

	assert.Equal(t, 35, m)
}

func TestDay5Input(t *testing.T) {
	seeds, maps := ReadSeedMaps("day5input.txt")
	res := MapSeedsThroughMaps(seeds, maps)
	m := Min(res)

	fmt.Println("Day 5 result (star 1):", m)
}

func TestDay5Star2InputExample(t *testing.T) {
	seeds, maps := ReadSeedMaps("day5input_example.txt")
	ranges := seedsToRanges(seeds)
	res := MapRangesThroughMaps(ranges, maps)
	m := MinOfRanges(res)

	assert.Equal(t, 46, m)
}

func seedsToRanges(seeds []int) []Range {
	var res []Range
	for i := 0; i < len(seeds); i += 2 {
		res = append(res, Range{seeds[i], seeds[i] + seeds[i+1]})
	}
	return res
}

func TestDay5Star2Input(t *testing.T) {
	seeds, maps := ReadSeedMaps("day5input.txt")
	ranges := seedsToRanges(seeds)
	res := MapRangesThroughMaps(ranges, maps)
	m := MinOfRanges(res)

	fmt.Println("Day 5 result (star 2):", m)
}
