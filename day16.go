package aoc2023

import (
	"fmt"
	"slices"
)

type CellType int

const (
	empty CellType = iota
	bltr
	tlbr
	lr
	tb
)

type Cell struct {
	cellType CellType
	hitFrom  []Direction
}

func toCellType(r rune) CellType {
	switch r {
	case '.':
		return empty
	case '/':
		return bltr
	case '\\':
		return tlbr
	case '-':
		return lr
	case '|':
		return tb
	default:
		return empty
	}
}

func PrintEnergized(cells [][]Cell) {
	for _, cellLine := range cells {
		for _, cell := range cellLine {
			if len(cell.hitFrom) > 0 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func CountEnergized(cells [][]Cell) int {
	sum := 0
	for _, cellLine := range cells {
		for _, cell := range cellLine {
			if len(cell.hitFrom) > 0 {
				sum++
			}
		}
	}
	return sum
}

func ResetEnergized(cells [][]Cell) {
	for y, cellLine := range cells {
		for x := range cellLine {
			cells[y][x].hitFrom = []Direction{}
		}
	}
}

func (c *Cell) HitFrom(d Direction) []Direction {
	if slices.Contains(c.hitFrom, d) {
		return []Direction{}
	}
	c.hitFrom = append(c.hitFrom, d)

	switch c.cellType {
	default:
		return []Direction{d}
	case lr:
		if d.dx != 0 {
			return []Direction{d}
		} else {
			return []Direction{{-1, 0}, {1, 0}}
		}
	case tb:
		if d.dx != 0 {
			return []Direction{{0, -1}, {0, 1}}
		} else {
			return []Direction{d}
		}
	case bltr:
		if d.dx != 0 {
			return []Direction{{0, -d.dx}}
		} else {
			return []Direction{{-d.dy, 0}}
		}
	case tlbr:
		if d.dx != 0 {
			return []Direction{{0, d.dx}}
		} else {
			return []Direction{{d.dy, 0}}
		}
	}
}

type Beam struct {
	coord Coord
	dir   Direction
}

func FollowBeam(cells [][]Cell, initial Beam) {
	beams := []Beam{initial}

	for len(beams) > 0 {
		beam := beams[0]
		beams = append(beams[1:], FollowBeamStep(cells, beam)...)
	}
}

func FollowBeamStep(cells [][]Cell, beam Beam) []Beam {
	coord := beam.coord
	if coord.x < 0 || coord.x >= len(cells[0]) || coord.y < 0 || coord.y >= len(cells) {
		// out of bounds, end of beam
		return []Beam{}
	}

	newDirs := cells[coord.y][coord.x].HitFrom(beam.dir)

	var newBeams []Beam
	for _, Direction := range newDirs {
		newBeams = append(newBeams, Beam{Coord{coord.x + Direction.dx, coord.y + Direction.dy}, Direction})
	}

	return newBeams
}

func ReadContraption(path string) [][]Cell {
	lines := ReadFile(path)

	var res [][]Cell
	for _, line := range lines {
		if line == "" {
			continue
		}
		var resLine []Cell
		for _, r := range line {
			resLine = append(resLine, Cell{toCellType(r), []Direction{}})
		}
		res = append(res, resLine)
	}
	return res
}

func FindMaxBeam(cells [][]Cell) int {
	var beams []Beam
	for y := range cells {
		beams = append(beams, Beam{Coord{0, y}, Direction{1, 0}}, Beam{Coord{len(cells[0]) - 1, y}, Direction{-1, 0}})
	}
	for x := range cells[0] {
		beams = append(beams, Beam{Coord{x, 0}, Direction{0, 1}}, Beam{Coord{x, len(cells) - 1}, Direction{0, -1}})
	}

	maxEnergized := 0
	for _, beam := range beams {
		FollowBeam(cells, beam)
		energized := CountEnergized(cells)
		maxEnergized = max(maxEnergized, energized)
		ResetEnergized(cells)
	}

	return maxEnergized
}
