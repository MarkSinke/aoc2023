package aoc2023

import "fmt"

type Tile struct {
	n      bool
	e      bool
	s      bool
	w      bool
	status Status
}

type Direction struct {
	dx int
	dy int
}

func (d Direction) Length() int {
	if d.dx != 0 {
		return absDiff(d.dx, 0)
	}
	return absDiff(d.dy, 0)
}

type Status int

const (
	undefined Status = iota
	out
	pipe
)

var North, East, South, West = Direction{0, -1}, Direction{1, 0}, Direction{0, 1}, Direction{-1, 0}
var AllDirections = []Direction{North, East, South, West}

func (t Tile) isStart() bool {
	return t.n && t.e && t.s && t.w
}

func (t Tile) isGround() bool {
	return !t.n && !t.e && !t.s && !t.w
}

func (t *Tile) setDirs(one Direction, two Direction) {
	t.n = false
	t.e = false
	t.s = false
	t.w = false

	t.setDir(one)
	t.setDir(two)
}

func (t *Tile) setDir(d Direction) {
	switch d {
	case North:
		t.n = true
	case East:
		t.e = true
	case South:
		t.s = true
	case West:
		t.w = true
	}
}

func (t Tile) findDir(canN bool, canE bool, canS bool, canW bool) Direction {
	if t.n && canN {
		return North
	} else if t.e && canE {
		return East
	} else if t.s && canS {
		return South
	} else if t.w && canW {
		return West
	}
	return Direction{0, 0}
}

func (t Tile) follow(d Direction) Direction {
	if d == North {
		return t.findDir(true, true, false, true)
	} else if d == East {
		return t.findDir(true, true, true, false)
	} else if d == South {
		return t.findDir(false, true, true, true)
	} else if d == West {
		return t.findDir(true, false, true, true)
	} else {
		return Direction{0, 0}
	}
}

func (t Tile) asRune() rune {
	if t.n && t.s {
		if t.w && t.e {
			return 'S'
		}
		return '\u2502'
	} else if t.w && t.e {
		return '\u2500'
	} else if t.n && t.e {
		return '\u2514'
	} else if t.e && t.s {
		return '\u250c'
	} else if t.s && t.w {
		return '\u2510'
	} else if t.w && t.n {
		return '\u2518'
	} else {
		return '#'
	}
}

func PrintConcisely(tiles [][]Tile) {
	for _, tileLine := range tiles {
		for _, tile := range tileLine {
			switch tile.status {
			case out:
				fmt.Print("O")
			case pipe:
				fmt.Print(string(tile.asRune()))
			case undefined:
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

// | is a vertical pipe connecting north and south.
// - is a horizontal pipe connecting east and west.
// L is a 90-degree bend connecting north and east.
// J is a 90-degree bend connecting north and west.
// 7 is a 90-degree bend connecting south and west.
// F is a 90-degree bend connecting south and east.
// . is ground; there is no pipe in this tile.
// S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.

var letterToTile = map[rune]Tile{
	'|': {true, false, true, false, undefined},
	'-': {false, true, false, true, undefined},
	'L': {true, true, false, false, undefined},
	'J': {true, false, false, true, undefined},
	'7': {false, false, true, true, undefined},
	'F': {false, true, true, false, undefined},
	'.': {false, false, false, false, undefined},
	'S': {true, true, true, true, undefined},
}

func parseTileLine(line string) []Tile {
	tileLine := make([]Tile, 0, len(line))
	for _, r := range line {
		tileLine = append(tileLine, letterToTile[r])
	}
	return tileLine
}

func ReadTiles(path string) [][]Tile {
	lines := ReadFile(path)

	var tiles [][]Tile
	for _, line := range lines {
		if line == "" {
			continue
		}
		tileLine := parseTileLine(line)
		tiles = append(tiles, tileLine)
	}

	return tiles
}

func findStartTile(tiles [][]Tile) (int, int) {
	for y, tileLine := range tiles {
		for x, tile := range tileLine {
			if tile.isStart() {
				return x, y
			}
		}
	}
	return -1, -1
}

func FindLoopLen(tiles [][]Tile) int {
	x, y := findStartTile(tiles)
	d := findDirFromStart(tiles, x, y)
	steps := 0
	for {
		x += d.dx
		y += d.dy
		steps += 1

		tile := tiles[y][x]
		if tile.isStart() {
			return steps
		}

		d = tile.follow(d)
	}
}

func findDirFromStart(tiles [][]Tile, x int, y int) Direction {
	if y > 0 && tiles[y-1][x].s {
		return North
	} else if y < len(tiles)-1 && tiles[y+1][x].n {
		return South
	} else if x > 0 && tiles[y][x-1].e {
		return West
	} else if x < len(tiles[y])-1 && tiles[y][x+1].w {
		return East
	}
	return Direction{0, 0}
}

func CopyMainLoop(tiles [][]Tile) [][]Tile {
	res := make([][]Tile, 0, len(tiles))

	for _, tileLine := range tiles {
		res = append(res, make([]Tile, len(tileLine)))
	}

	x, y := findStartTile(tiles)
	dStart := findDirFromStart(tiles, x, y)
	d := dStart
	steps := 0
	for {
		x += d.dx
		y += d.dy
		steps += 1

		tile := tiles[y][x]
		tile.status = pipe
		res[y][x] = tile
		if tile.isStart() {
			res[y][x].setDirs(dStart, Direction{-d.dx, -d.dy})
			return res
		}

		d = tile.follow(d)
	}
}

func appendExploded(line0 []Tile, line1 []Tile, tile Tile) ([]Tile, []Tile) {
	if tile.n && tile.s {
		return append(line0, Tile{}, tile), append(line1, Tile{}, tile)
	} else if tile.w && tile.e {
		return append(line0, Tile{}, Tile{}), append(line1, tile, tile)
	} else if tile.n && tile.e {
		return append(line0, Tile{}, tile), append(line1, Tile{}, Tile{})
	} else if tile.e && tile.s {
		return append(line0, Tile{}, Tile{}), append(line1, Tile{}, tile)
	} else if tile.s && tile.w {
		return append(line0, Tile{}, Tile{}), append(line1, tile, Tile{})
	} else if tile.w && tile.n {
		return append(line0, tile, Tile{}), append(line1, Tile{}, Tile{})
	} else {
		return append(line0, Tile{}, Tile{}), append(line1, Tile{}, Tile{})
	}
}

func getImploded(tl Tile, tr Tile, bl Tile, br Tile) Tile {
	if !tl.isGround() {
		return tl
	}
	if !tr.isGround() {
		return tr
	}
	if !bl.isGround() {
		return bl
	}
	// either br is ground, or the return value is abitrary - hence return bottom-right
	return br
}

// make a 2x2 for every tile, such that we always have a 2 or more empty tiles in an exploded tile
// this makes sure we always have a path to outer cells
func ExplodeLoop(tiles [][]Tile) [][]Tile {
	res := make([][]Tile, 0, len(tiles)*2)
	for _, tileLine := range tiles {
		resLine0 := make([]Tile, 0, len(tileLine)*2)
		resLine1 := make([]Tile, 0, len(tileLine)*2)
		for _, tile := range tileLine {
			resLine0, resLine1 = appendExploded(resLine0, resLine1, tile)
		}
		res = append(res, resLine0, resLine1)
	}
	return res
}

func ImplodeLoop(tiles [][]Tile) [][]Tile {
	res := make([][]Tile, 0, len(tiles)/2)
	for y := 0; y < len(tiles)/2; y++ {
		resLine := make([]Tile, 0, len(tiles[0])/2)
		for x := 0; x < len(tiles[0])/2; x++ {
			resLine = append(resLine, getImploded(tiles[2*y][2*x], tiles[2*y][2*x+1], tiles[2*y+1][2*x], tiles[2*y+1][2*x+1]))
		}
		res = append(res, resLine)
	}
	return res

}

func appendIfUndefined(tiles [][]Tile, found []Coord, c Coord) []Coord {
	if c.x >= 0 && c.x < len(tiles[0]) &&
		c.y >= 0 && c.y < len(tiles) &&
		tiles[c.y][c.x].status == undefined {
		tiles[c.y][c.x].status = out
		return append(found, c)
	}
	return found
}

func findAndMarkEligibleTiles(tiles [][]Tile, coord Coord) []Coord {
	var eligible = []Coord{}
	eligible = appendIfUndefined(tiles, eligible, Coord{coord.x + 1, coord.y})
	eligible = appendIfUndefined(tiles, eligible, Coord{coord.x - 1, coord.y})
	eligible = appendIfUndefined(tiles, eligible, Coord{coord.x, coord.y + 1})
	eligible = appendIfUndefined(tiles, eligible, Coord{coord.x, coord.y - 1})
	return eligible
}

func Fill(tiles [][]Tile) {
	toCheck := []Coord{}

	// check borders first
	for x := range tiles[0] {
		toCheck = appendIfUndefined(tiles, toCheck, Coord{x, 0})
		toCheck = appendIfUndefined(tiles, toCheck, Coord{x, len(tiles) - 1})
	}
	for y := range tiles {
		toCheck = appendIfUndefined(tiles, toCheck, Coord{0, y})
		toCheck = appendIfUndefined(tiles, toCheck, Coord{len(tiles[0]) - 1, y})
	}

	for len(toCheck) > 0 {
		c := toCheck[0]
		toCheck = toCheck[1:]
		newTiles := findAndMarkEligibleTiles(tiles, c)
		toCheck = append(toCheck, newTiles...)
	}
}

func CountIn(tiles [][]Tile) int {
	in := 0
	for _, tileLine := range tiles {
		for _, tile := range tileLine {
			if tile.status == undefined {
				in++
			}
		}
	}
	return in
}

func CountInOrPipe(tiles [][]Tile) int {
	in := 0
	for _, tileLine := range tiles {
		for _, tile := range tileLine {
			if tile.status == undefined || tile.status == pipe {
				in++
			}
		}
	}
	return in
}
