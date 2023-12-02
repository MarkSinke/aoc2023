package aoc2023

import (
	"os"
	"regexp"
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

var digitRegex = regexp.MustCompile("0|1|2|3|4|5|6|7|8|9|one|two|three|four|five|six|seven|eight|nine")
var textToValue = map[string]int{
	"0":     0,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func FirstAndLastDigitWithWords(line string) (int, int) {
	println(line)
	matches := digitRegex.FindAllString(line, -1)
	first, last := matches[0], matches[len(matches)-1]
	println(first, last)
	return textToValue[first], textToValue[last]
}

func CalibrationValue(first string, last string) int {
	res, _ := strconv.Atoi(first + last)
	return res
}
