package aoc2023

import (
	"fmt"
	"strings"
	"testing"
)

func TestDay1Example(t *testing.T) {
	sum := ComputeCalibrationSum(t, "day1input_example.txt")
	if sum != 142 {
		t.Errorf("expected 142, but got %d", sum)
	}
}

func ComputeCalibrationSum(t *testing.T, filename string) int {
	lines, err := ReadFile(filename)
	if err != nil {
		t.Fatal("input file cannot be read")
	}
	sum := 0
	for _, str := range lines {
		sum += CalibrationValue(FirstAndLastDigit(str))
	}
	return sum
}

func TestDay1Input(t *testing.T) {
	fmt.Printf("Day 1 result (star 1): %d\n", ComputeCalibrationSum(t, "day1input.txt"))
}

func ComputeCalibrationSumWordAware(t *testing.T, filename string) int {
	lines, err := ReadFile(filename)
	if err != nil {
		t.Fatal("input file cannot be read")
	}
	sum := 0
	replacer := strings.NewReplacer("one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9")
	for _, str := range lines {
		sum += CalibrationValue(FirstAndLastDigit(replacer.Replace(str)))
	}
	return sum
}

func TestDay1Example2(t *testing.T) {
	sum := ComputeCalibrationSumWordAware(t, "day1input_example2.txt")
	if sum != 281 {
		t.Errorf("expected 142, but got %d", sum)
	}
}

func TestDay1InputStar2(t *testing.T) {
	fmt.Printf("Day 1 result (star 2): %d\n", ComputeCalibrationSumWordAware(t, "day1input.txt"))
}
