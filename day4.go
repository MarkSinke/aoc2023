package aoc2023

import (
	"regexp"
	"strconv"
)

type GameCard struct {
	id             int
	winningNumbers []int
	numbers        []int
}

func (g GameCard) Score() int {
	score := 0
	for _, num := range g.numbers {
		for _, win := range g.winningNumbers {
			if win == num {
				if score == 0 {
					score = 1
				} else {
					score = score * 2
				}
			}
		}
	}
	return score
}

var gameCardRegex = regexp.MustCompile("Card +(\\d+): ([\\d ]+) \\| ([\\d ]+)")

func ReadGameCards(path string) []GameCard {
	lines := ReadFile(path)

	var gameCards []GameCard
	for _, line := range lines {
		res := gameCardRegex.FindStringSubmatch(line)
		gameCards = append(gameCards, parseCard(res[1], res[2], res[3]))
	}
	return gameCards
}

func parseCard(idStr string, winningNumStr string, numStr string) GameCard {
	id, _ := strconv.Atoi(idStr)
	winningNumbers := parseNumList(winningNumStr)
	numbers := parseNumList(numStr)
	return GameCard{id, winningNumbers, numbers}
}

var numRegex = regexp.MustCompile("\\d+")

func parseNumList(str string) []int {
	nums := numRegexp.FindAllString(str, -1)

	var res []int
	for _, numStr := range nums {
		num, _ := strconv.Atoi(numStr)
		res = append(res, num)
	}
	return res
}
