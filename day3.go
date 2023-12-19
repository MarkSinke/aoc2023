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

func absDiff(x int, y int) int {
	if x < y {
		return y - x
	} else {
		return x - y
	}
}

func Dist(coord0 Coord, coord1 Coord) int {
	return absDiff(coord0.x, coord1.x) + absDiff(coord0.y, coord1.y)
}

func ToDir(from Coord, to Coord) Direction {
	return Direction{to.x - from.x, to.y - from.y}
}

type NumberCoord struct {
	number int
	xMin   int
	xMax   int
	y      int
}

func (n NumberCoord) touching(xy Coord) bool {
	if xy.y < n.y-1 {
		return false // above
	}
	if xy.y > n.y+1 {
		return false // below
	}
	if xy.x < n.xMin-1 {
		return false // left
	}
	if xy.x > n.xMax {
		return false // right (xMax is first index after the number, so no need to +1 even more)
	}
	return true
}

// coordinate set, implemted as a map to 0-sized structs
type CoordSet map[Coord]rune

func (s CoordSet) add(xy Coord, r rune) {
	s[xy] = r
}

func (s CoordSet) has(xy Coord) bool {
	_, ok := s[xy]
	return ok
}

func (s CoordSet) findAll(r rune) []Coord {
	var coords []Coord
	for coord, rune := range s {
		if rune == r {
			coords = append(coords, coord)
		}
	}
	return coords
}

func GetSymbolCoordSet(engine []string) CoordSet {
	coords := CoordSet{}
	for y, line := range engine {
		for x, rune := range line {
			if !unicode.IsDigit(rune) && rune != '.' {
				coords.add(Coord{x, y}, rune)
			}
		}
	}
	return coords
}

var numRegexp = regexp.MustCompile(`-?\d+`)

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

func Sum(ints []int) int {
	sum := 0
	for _, val := range ints {
		sum += val
	}
	return sum
}

func FindConnectedNumbers(coord Coord, numbers []NumberCoord) []NumberCoord {
	var result []NumberCoord
	for _, num := range numbers {
		if num.touching(coord) {
			result = append(result, num)
		}
	}
	return result
}
