package aoc2023

import (
	"strconv"
	"strings"
)

type RaceGame struct {
	time       int
	recordDist int
}

func ReadRaceGames(path string) []RaceGame {
	lines := ReadFile(path)

	times := ParseNumList(CutLabel(lines[0]))
	records := ParseNumList(CutLabel(lines[1]))

	var games []RaceGame
	for i, time := range times {
		games = append(games, RaceGame{time, records[i]})
	}
	return games
}

func ReadRaceGame(path string) RaceGame {
	lines := ReadFile(path)

	time, _ := strconv.Atoi(CutLabel(strings.ReplaceAll(lines[0], " ", "")))
	record, _ := strconv.Atoi(CutLabel(strings.ReplaceAll(lines[1], " ", "")))

	return RaceGame{time, record}
}

func WinCount(game RaceGame) int {
	wins := 0
	for i := 0; i < game.time; i++ {
		speed := i
		dist := (game.time - i) * speed

		if dist > game.recordDist {
			wins++
		}
	}
	return wins
}

func WinCounts(games []RaceGame) int {
	product := 1

	for _, game := range games {
		product *= WinCount(game)
	}
	return product
}
