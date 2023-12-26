package aoc2023

import (
	"fmt"
	"testing"
)

func TestDay3InputExample(t *testing.T) {
	engine := ReadFile("day3input_example.txt")
	parts := GetPartNumbers(engine)
	sum := Sum(parts)

	if sum != 4361 {
		t.Errorf("expected 4361, but got %d", sum)
	}
}

func TestDay3Input(t *testing.T) {
	engine := ReadFile("../aoc_inputs/2023/day3input.txt")
	parts := GetPartNumbers(engine)
	sum := Sum(parts)

	fmt.Printf("Day 3 result (star 1): %d\n", sum)
}

func TestDay3Star2InputExample(t *testing.T) {
	engine := ReadFile("day3input_example.txt")
	gears := GetSymbolCoordSet(engine).FindAll('*')
	numbers := GetNumbers(engine)
	sum := computeGearSum(gears, numbers)

	if sum != 467835 {
		t.Errorf("expcted 467835, but got %v", sum)
	}
}

func computeGearSum(gears []Coord, numbers []NumberCoord) int {
	sum := 0
	for _, gear := range gears {
		connectedNums := FindConnectedNumbers(gear, numbers)
		if len(connectedNums) == 2 {
			sum += connectedNums[0].number * connectedNums[1].number
		}
	}
	return sum
}

func TestDay3Star2Input(t *testing.T) {
	engine := ReadFile("../aoc_inputs/2023/day3input.txt")
	gears := GetSymbolCoordSet(engine).FindAll('*')
	numbers := GetNumbers(engine)
	sum := computeGearSum(gears, numbers)

	fmt.Printf("Day 3 result (star 2): %v\n", sum)
}
