package aoc2023

import (
	"regexp"
	"strconv"
)

type DigStep struct {
	dir   Direction
	count int
	// color RGBA
}

var digSpec = regexp.MustCompile(`([URDL]) ([0-9]+) \(#([0-9a-f]{2})([0-9a-f]{2})([0-9a-f]{2})\)`)

func parseStep(str string) DigStep {
	matches := digSpec.FindStringSubmatch(str)

	var dir Direction
	switch matches[1] {
	case "U":
		dir = North
	case "R":
		dir = East
	case "D":
		dir = South
	case "L":
		dir = West
	default:
		panic("unknown direction " + matches[1])
	}

	count, _ := strconv.Atoi(matches[2])

	return DigStep{dir, count}
}

func ReadDigPlan(path string) []DigStep {
	lines := ReadFile(path)

	var steps []DigStep
	for _, line := range lines {
		if line == "" {
			continue
		}

		steps = append(steps, parseStep(line))
	}

	return steps
}

var digSpec2 = regexp.MustCompile(`([URDL]) ([0-9]+) \(#([0-9a-f]{5})([0-9a-f])\)`)

func parseStep2(str string) DigStep {
	matches := digSpec2.FindStringSubmatch(str)

	var dir Direction
	switch matches[4] {
	case "0":
		dir = East
	case "1":
		dir = South
	case "2":
		dir = West
	case "3":
		dir = North
	default:
		panic("unknown direction " + matches[1])
	}

	count, _ := strconv.ParseInt(matches[3], 16, 0)

	return DigStep{dir, int(count)}
}

func ReadDigPlan2(path string) []DigStep {
	lines := ReadFile(path)

	var steps []DigStep
	for _, line := range lines {
		if line == "" {
			continue
		}

		steps = append(steps, parseStep2(line))
	}

	return steps
}

func getExtent(steps []DigStep) (Coord, Coord) {
	c := Coord{0, 0}
	minX := 0
	minY := 0
	maxX := 0
	maxY := 0
	for _, step := range steps {
		c.x += step.dir.dx * step.count
		c.y += step.dir.dy * step.count

		maxX = max(maxX, c.x)
		maxY = max(maxY, c.y)

		minX = min(minX, c.x)
		minY = min(minY, c.y)
	}

	return Coord{minX, minY}, Coord{maxX + 1, maxY + 1}
}

func createSurface(extent Coord) [][]Tile {
	tiles := make([][]Tile, extent.y)

	for y := range tiles {
		tiles[y] = make([]Tile, extent.x)
	}

	return tiles
}

func ExecuteDigPlan(steps []DigStep) [][]Tile {
	minCoord, maxCoord := getExtent(steps)
	tiles := createSurface(Coord{maxCoord.x - minCoord.x, maxCoord.y - minCoord.y})

	c := Coord{0, 0}

	for _, step := range steps {
		for i := 0; i < step.count; i++ {
			tiles[c.y-minCoord.y][c.x-minCoord.x].status = pipe
			c.x += step.dir.dx
			c.y += step.dir.dy
		}
	}
	return tiles
}
