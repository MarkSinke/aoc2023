package aoc2023

import (
	"fmt"
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
	for _, str := range lines {
		first, last := FirstAndLastDigitWithWords(str)
		sum += first*10 + last
	}
	return sum
}

func TestDay1Example2(t *testing.T) {
	sum := ComputeCalibrationSumWordAware(t, "day1input_example2.txt")
	if sum != 281 {
		t.Errorf("expected 281, but got %d", sum)
	}
}

func TestDay1InputStar2(t *testing.T) {
	fmt.Printf("Day 1 result (star 2): %d\n", ComputeCalibrationSumWordAware(t, "day1input.txt"))
}

func TestReverse(t *testing.T) {
	str := "This is a test"
	reversed := Reverse(str)
	expected := "tset a si sihT"
	if reversed != expected {
		t.Errorf("expected %s, but got %s", expected, reversed)
	}
}
