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

func Reverse(str string) string {
	chars := []rune(str)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

var regexString = "0|1|2|3|4|5|6|7|8|9|one|two|three|four|five|six|seven|eight|nine"
var digitRegex = regexp.MustCompile(regexString)
var digitRegexRev = regexp.MustCompile(Reverse(regexString))
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
	"eno":   1,
	"two":   2,
	"owt":   2,
	"three": 3,
	"eerht": 3,
	"four":  4,
	"ruof":  4,
	"five":  5,
	"evif":  5,
	"six":   6,
	"xis":   6,
	"seven": 7,
	"neves": 7,
	"eight": 8,
	"thgie": 8,
	"nine":  9,
	"enin":  9,
}

func FirstAndLastDigitWithWords(line string) (int, int) {
	first := digitRegex.FindString(line)
	last := digitRegexRev.FindString(Reverse(line))
	return textToValue[first], textToValue[last]
}

func CalibrationValue(first string, last string) int {
	res, _ := strconv.Atoi(first + last)
	return res
}
