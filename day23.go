package aoc2023

type MazeTile struct {
	path    bool
	slope   Direction
	visited bool
}

type Maze [][]MazeTile

func (m Maze) isValidStep(c Coord) bool {
	inMaze := c.x >= 0 && c.x < len(m[0]) && c.y >= 0 && c.y < len(m)
	return inMaze && m[c.y][c.x].path && !m[c.y][c.x].visited
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
	return findLongestPath(start, m)
}

func findLongestPath(c Coord, m Maze) int {
	maxTail := 0
	m[c.y][c.x].visited = true
	for _, coord := range findNextSteps(c, m) {
		maxNew := findLongestPath(coord, m) + 1
		maxTail = max(maxTail, maxNew)
	}
	m[c.y][c.x].visited = false
	return maxTail
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

	coords := []Coord{}
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
}
