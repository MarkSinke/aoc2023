package aoc2023

import (
	"regexp"
	"slices"
	"strconv"
)

type Coord3 struct {
	x int
	y int
	z int
}

type Brick struct {
	c0 Coord3
	c1 Coord3
}

func (b Brick) MinX() int {
	return min(b.c0.x, b.c1.x)
}

func (b Brick) MinY() int {
	return min(b.c0.y, b.c1.y)
}

func (b Brick) MinZ() int {
	return min(b.c0.z, b.c1.z)
}

func (b Brick) MaxX() int {
	return max(b.c0.x, b.c1.x)
}

func (b Brick) MaxY() int {
	return max(b.c0.y, b.c1.y)
}

func (b Brick) MaxZ() int {
	return max(b.c0.z, b.c1.z)
}

func (b *Brick) DropTo(z int) {
	minZ := b.MinZ()

	b.c0.z = b.c0.z - minZ + z
	b.c1.z = b.c1.z - minZ + z
}

func (b Brick) XYCoords() []Coord {
	coords := []Coord{}
	for x := b.MinX(); x <= b.MaxX(); x++ {
		for y := b.MinY(); y <= b.MaxY(); y++ {
			coords = append(coords, Coord{x, y})
		}
	}
	return coords
}

func (b Brick) Supports(above Brick) bool {
	if b.MaxZ()+1 != above.MinZ() {
		return false
	}

	for _, c0 := range b.XYCoords() {
		for _, c1 := range above.XYCoords() {
			if c0 == c1 {
				return true
			}
		}
	}

	return false
}

var brickRegex = regexp.MustCompile(`(\d+),(\d+),(\d+)~(\d+),(\d+),(\d+)`)

func toInt(str string) int {
	res, _ := strconv.Atoi(str)
	return res
}

func parseBrick(str string) Brick {
	matches := brickRegex.FindStringSubmatch(str)

	c0 := Coord3{toInt(matches[1]), toInt(matches[2]), toInt(matches[3])}
	c1 := Coord3{toInt(matches[4]), toInt(matches[5]), toInt(matches[6])}

	return Brick{c0, c1}
}

func ReadBricks(path string) []Brick {
	lines := ReadFile(path)

	bricks := []Brick{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		brick := parseBrick(line)
		bricks = append(bricks, brick)
	}

	return bricks
}

func compareLowestZ(b0 Brick, b1 Brick) int {
	return b0.MinZ() - b1.MinZ()
}

func findFreeZ(zPlane map[Coord]int, coords []Coord) int {
	freeZ := 1

	for _, coord := range coords {
		z, found := zPlane[coord]

		if found {
			freeZ = max(z+1, freeZ)
		}
	}

	return freeZ
}

func DropBricks(bricks []Brick) {
	slices.SortFunc(bricks, compareLowestZ)

	zPlane := map[Coord]int{}
	for i := range bricks {
		coords := bricks[i].XYCoords()
		freeZ := findFreeZ(zPlane, coords)
		bricks[i].DropTo(freeZ)
		for _, coord := range coords {
			zPlane[coord] = bricks[i].MaxZ()
		}
	}
}

func CountDisintegratable(bricks []Brick) int {
	// a brick can disintegrate if it's not supporting any bricks by itself only
	// i.e. supportedBy[b] > 1
	supportedBy := map[Brick]int{}
	supports := map[Brick][]Brick{}

	for _, b0 := range bricks {
		for _, b1 := range bricks {
			if b0.Supports(b1) {
				count, found := supportedBy[b1]
				if !found {
					count = 1
				} else {
					count++
				}
				supportedBy[b1] = count

				list, found := supports[b0]
				if !found {
					list = []Brick{b1}
				} else {
					list = append(list, b1)
				}
				supports[b0] = list
			}
		}
	}

	res := 0
	for _, b := range bricks {
		if isDisintegratable(b, supports, supportedBy) {
			res++
		}
	}

	return res
}

func isDisintegratable(b Brick, supports map[Brick][]Brick, supportedBy map[Brick]int) bool {
	aboves := supports[b]
	for _, above := range aboves {
		if supportedBy[above] == 1 {
			return false
		}
	}

	return true
}
