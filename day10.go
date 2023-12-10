package aoc2023

type Tile struct {
	n bool
	e bool
	s bool
	w bool
}

type Direction struct {
	dx int
	dy int
}

var north, east, south, west = Direction{0, -1}, Direction{1, 0}, Direction{0, 1}, Direction{-1, 0}

func (t Tile) isStart() bool {
	return t.n && t.e && t.s && t.w
}

func (t Tile) findDir(canN bool, canE bool, canS bool, canW bool) Direction {
	if t.n && canN {
		return north
	} else if t.e && canE {
		return east
	} else if t.s && canS {
		return south
	} else if t.w && canW {
		return west
	}
	return Direction{0, 0}
}

func (t Tile) follow(d Direction) Direction {
	if d == north {
		return t.findDir(true, true, false, true)
	} else if d == east {
		return t.findDir(true, true, true, false)
	} else if d == south {
		return t.findDir(false, true, true, true)
	} else if d == west {
		return t.findDir(true, false, true, true)
	} else {
		return Direction{0, 0}
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
	'|': {true, false, true, false},
	'-': {false, true, false, true},
	'L': {true, true, false, false},
	'J': {true, false, false, true},
	'7': {false, false, true, true},
	'F': {false, true, true, false},
	'.': {false, false, false, false},
	'S': {true, true, true, true},
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

	return followLoop(tiles, x, y, d)
}

func findDirFromStart(tiles [][]Tile, x int, y int) Direction {
	if y > 0 && tiles[y-1][x].s {
		return north
	} else if y < len(tiles)-1 && tiles[y+1][x].n {
		return south
	} else if x > 0 && tiles[y][x-1].e {
		return west
	} else if x < len(tiles[y])-1 && tiles[y][x+1].w {
		return east
	}
	return Direction{0, 0}
}

func followLoop(tiles [][]Tile, xS int, yS int, d Direction) int {
	steps := 0
	x := xS
	y := yS
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
