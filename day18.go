package aoc2023

import (
	"regexp"
	"strconv"
)

type DigStep struct {
	dir   Direction
	count int
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

func ComputeCoords(steps []DigStep) []Coord {
	c := Coord{0, 0}

	coords := []Coord{}
	for _, step := range steps {
		coords = append(coords, c)
		c.x += step.dir.dx * step.count
		c.y += step.dir.dy * step.count
	}
	return coords
}

func inProduct(c0 Coord, c1 Coord) int {
	return c0.x*c1.y - c1.x*c0.y
}

// compute polygon coords using Green's theorem
func ComputePolygonArea(coords []Coord) int {
	area := 0      // Green: area = 1/2 * SUM inProducts
	perimeter := 0 // SUM of lengths of step vectors

	for i := range coords {
		c0 := coords[i]
		c1 := coords[(i+1)%len(coords)]

		area += inProduct(c0, c1)
		perimeter += Dist(c0, c1)
	}

	// now we have 2x the area, using the midpoints of the coords
	// we need to add the "outer" half pixel to the area
	// i.e., we need to add the perimeter / 2 + 1  and area / 2
	return (area+perimeter)/2 + 1
}
