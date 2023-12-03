package aoc2023

import (
	"regexp"
	"strconv"
	"unicode"
)

type Coord struct {
	x int
	y int
}

type NumberCoord struct {
	number int
	xMin   int
	xMax   int
	y      int
}

// coordinate set, implemted as a map to 0-sized structs
type CoordSet map[Coord]struct{}

func (s CoordSet) add(xy Coord) {
	s[xy] = struct{}{}
}

func (s CoordSet) remove(xy Coord) {
	delete(s, xy)
}

func (s CoordSet) has(xy Coord) bool {
	_, ok := s[xy]
	return ok
}

func GetSymbolCoordSet(engine []string) CoordSet {
	coords := CoordSet{}
	for y, line := range engine {
		for x, rune := range line {
			if !unicode.IsDigit(rune) && rune != '.' {
				coords.add(Coord{x, y})
			}
		}
	}
	return coords
}

var numRegexp = regexp.MustCompile("\\d+")

func GetNumbers(engine []string) []NumberCoord {
	var coords []NumberCoord

	for y, line := range engine {
		matches := numRegexp.FindAllStringIndex(line, -1)
		for _, match := range matches {
			xMin := match[0]
			xMax := match[1]
			number, _ := strconv.Atoi(line[xMin:xMax])
			coords = append(coords, NumberCoord{number, xMin, xMax, y})
		}
	}
	return coords
}

func GetPartNumbers(engine []string) []int {
	symbols := GetSymbolCoordSet(engine)
	numbers := GetNumbers(engine)

	var result []int
	for _, num := range numbers {
		if isSymbolClose(num, symbols) {
			result = append(result, num.number)
		}
	}
	return result
}

func isSymbolClose(num NumberCoord, symbols CoordSet) bool {
	for x := num.xMin - 1; x < num.xMax+1; x++ {
		if symbols.has(Coord{x, num.y - 1}) || symbols.has(Coord{x, num.y}) || symbols.has(Coord{x, num.y + 1}) {
			return true
		}
	}
	return false
}

func sum(ints []int) int {
	sum := 0
	for _, val := range ints {
		sum += val
	}
	return sum
}
