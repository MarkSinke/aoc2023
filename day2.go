package aoc2023

import (
	"regexp"
	"strconv"
	"strings"
)

type DiceResult struct {
	red   int
	green int
	blue  int
}

type GameResult struct {
	id      int
	results []DiceResult
}

var gameResultRegex = regexp.MustCompile(`Game (\d+): (.*)`)

func ReadGameResults(path string) []GameResult {
	lines := ReadFile(path)

	gameResults := []GameResult{}
	for _, line := range lines {
		res := gameResultRegex.FindStringSubmatch(line)
		id, _ := strconv.Atoi(res[1])
		content := res[2]

		diceResults := ParseDiceResults(content)

		gameResults = append(gameResults, GameResult{id, diceResults})
	}
	return gameResults
}

func ParseDiceResults(content string) []DiceResult {
	diceResultStrs := strings.Split(content, ";")
	diceResults := []DiceResult{}
	for _, diceResultStr := range diceResultStrs {
		diceResults = append(diceResults, ParseDiceResult(diceResultStr))
	}
	return diceResults
}

var colorRegex = regexp.MustCompile(`(\d+) (red|green|blue)`)

func ParseDiceResult(str string) DiceResult {
	resultStrs := strings.Split(str, ",")
	red, green, blue := 0, 0, 0
	for _, resultStr := range resultStrs {
		res := colorRegex.FindStringSubmatch(resultStr)
		count, _ := strconv.Atoi(res[1])
		color := res[2]
		switch color {
		case "red":
			red = count
		case "green":
			green = count
		case "blue":
			blue = count
		}
	}
	return DiceResult{red: red, green: green, blue: blue}
}

func CanPlayWith(red int, green int, blue int, gameResult GameResult) bool {
	for _, res := range gameResult.results {
		if res.red > red || res.green > green || res.blue > blue {
			return false
		}
	}
	return true
}

func MinPower(gameResult GameResult) int {
	red, green, blue := 0, 0, 0
	for _, res := range gameResult.results {
		red = max(red, res.red)
		green = max(green, res.green)
		blue = max(blue, res.blue)
	}
	return red * blue * green
}
