package aoc2023

import (
	"fmt"
	"slices"
)

type MazeTile struct {
	path       bool
	slope      Direction
	visited    bool
	validSteps []Coord
}

type Maze [][]MazeTile

func (m Maze) isValidStep(c Coord) bool {
	inMaze := c.x >= 0 && c.x < len(m[0]) && c.y >= 0 && c.y < len(m)
	return inMaze && m[c.y][c.x].path
}

func ReadMaze(path string) Maze {
	lines := ReadFile(path)

	maze := Maze{}
	for _, line := range lines {
		if line == "" {
			continue
		}

		mazeLine := parseMazeLine(line)
		maze = append(maze, mazeLine)
	}

	precomputeValidSteps(maze)

	return maze
}

func parseMazeLine(str string) []MazeTile {
	mazeLine := []MazeTile{}
	for _, r := range str {
		mazeLine = append(mazeLine, parseMazeTile(r))
	}
	return mazeLine
}

func parseMazeTile(r rune) MazeTile {
	switch r {
	case '.':
		return MazeTile{path: true}
	case '#':
		return MazeTile{path: false}
	case '^':
		return MazeTile{path: true, slope: North}
	case '>':
		return MazeTile{path: true, slope: East}
	case 'v':
		return MazeTile{path: true, slope: South}
	case '<':
		return MazeTile{path: true, slope: West}
	default:
		panic("unknwon tile " + string(r))
	}
}

func FindLongestPath(m Maze) int {
	start := Coord{1, 0}
	return findLongestPath(&m[0][1], start, m)
}

func findLongestPath(ptile *MazeTile, c Coord, m Maze) int {
	maxTailPath := 0
	ptile.visited = true
	for _, coord := range ptile.validSteps {
		ptile := &m[coord.y][coord.x]
		if !ptile.visited {
			maxNew := findLongestPath(ptile, coord, m) + 1
			maxTailPath = max(maxTailPath, maxNew)
		}
	}
	ptile.visited = false

	// must end on bottom right exit
	if maxTailPath == 0 && (c.x != len(m[0])-2 || c.y != len(m)-1) {
		return -1000
	}
	return maxTailPath
}

func findNextSteps(c Coord, m Maze) []Coord {
	tile := m[c.y][c.x]
	if tile.slope.dx != 0 || tile.slope.dy != 0 {
		newCoord := AddDir(c, tile.slope)
		if !m.isValidStep(newCoord) {
			return []Coord{}
		}
		return []Coord{newCoord}
	}

	coords := make([]Coord, 0, 4)
	for _, dir := range AllDirections {
		newCoord := AddDir(c, dir)
		if m.isValidStep(newCoord) {
			coords = append(coords, newCoord)
		}
	}
	return coords
}

func FlattenSlopes(m Maze) {
	for y := range m {
		for x := range m[0] {
			m[y][x].slope = Direction{0, 0}
		}
	}

	precomputeValidSteps(m)
}

func precomputeValidSteps(m Maze) {
	for y := range m {
		for x := range m[0] {
			m[y][x].validSteps = findNextSteps(Coord{x, y}, m)
		}
	}
}

func PrintMazePath(m Maze, coords []Coord) {
	for y := range m {
		for x := range m[0] {
			if slices.Index(coords, Coord{x, y}) != -1 {
				fmt.Print("O")
			} else if m[y][x].path {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}
