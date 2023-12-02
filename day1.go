package aoc2023

import (
	"os"
	"strconv"
	"strings"
	"unicode"
)

func ReadFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}

func FirstAndLastDigit(line string) (string, string) {
	firstDigit, lastDigit := "", ""
	for _, rune := range line {
		if unicode.IsDigit(rune) {
			if firstDigit == "" {
				firstDigit = string(rune)
			}
			lastDigit = string(rune)
		}
	}
	return firstDigit, lastDigit
}

func CalibrationValue(first string, last string) int {
	res, _ := strconv.Atoi(first + last)
	return res
}
