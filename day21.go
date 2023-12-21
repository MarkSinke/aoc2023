package aoc2023

import (
	"fmt"
)

type GardenPlot struct {
	rock             bool
	visitedStepsLeft int
}

type Garden [][]GardenPlot

func ReadGarden(path string) (Garden, Coord) {
	lines := ReadFile(path)
	start := Coord{-1, -1}

	garden := Garden{}
	for y, line := range lines {
		if line == "" {
			continue
		}
		gardenLine := []GardenPlot{}
		for x, rune := range line {
			switch rune {
			case 'S':
				gardenLine = append(gardenLine, GardenPlot{false, -1})
				start = Coord{x, y}
			case '.':
				gardenLine = append(gardenLine, GardenPlot{false, -1})
			case '#':
				gardenLine = append(gardenLine, GardenPlot{true, -1})
			default:
				panic("unknown type")
			}
		}
		garden = append(garden, gardenLine)
	}
	return garden, start
}

var twoSteps = [][]Direction{
	{North, North},
	{North, East},
	{North, West},
	{East, North},
	{East, East},
	{East, South},
	{South, East},
	{South, South},
	{South, West},
	{West, North},
	{West, South},
	{West, West},
}

func addDir(c Coord, d Direction) Coord {
	return Coord{c.x + d.dx, c.y + d.dy}
}

func canStep(garden Garden, c Coord) bool {
	return c.x >= 0 && c.x < len(garden[0]) && c.y >= 0 && c.y < len(garden) &&
		!garden[c.y][c.x].rock
}

func WalkGarden(garden Garden, c Coord, maxSteps int) {
	if garden[c.y][c.x].visitedStepsLeft >= maxSteps {
		return
	}
	garden[c.y][c.x].visitedStepsLeft = maxSteps

	if maxSteps <= 0 {
		return
	}

	for _, twoStep := range twoSteps {
		c1 := addDir(c, twoStep[0])
		c2 := addDir(c1, twoStep[1])
		if canStep(garden, c1) && canStep(garden, c2) {
			WalkGarden(garden, c2, maxSteps-2)
		}
	}
}

func CountVisited(garden Garden) int {
	count := 0

	for _, line := range garden {
		for _, plot := range line {
			if plot.visitedStepsLeft >= 0 {
				count++
			}
		}
	}
	return count
}

func PrintGarden(garden Garden, start Coord) {
	for y, line := range garden {
		for x, plot := range line {
			if plot.visitedStepsLeft >= 0 {
				fmt.Print("O")
			} else if plot.rock {
				fmt.Print("#")
			} else if start.x == x && start.y == y {
				fmt.Print("S")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
