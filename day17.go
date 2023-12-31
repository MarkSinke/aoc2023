package aoc2023

import (
	"github.com/fzipp/astar"
)

type HeatGrid [][]int

func (g HeatGrid) GetBounds() (int, int) {
	return len(g[0]), len(g)
}

func (g HeatGrid) GetLoss(c Coord) int {
	return g[c.y][c.x]
}

func (g HeatGrid) isValidCoord(to Coord) bool {
	mx, my := g.GetBounds()
	return to.x >= 0 && to.x < mx && to.y >= 0 && to.y < my
}

type SearchNode struct {
	coord      Coord
	horizontal bool
}

type SearchGrid struct {
	grid    HeatGrid
	minStep int
	maxStep int
}

func (g SearchGrid) Neighbours(c SearchNode) []SearchNode {
	coords := []SearchNode{}
	var dirs []Direction
	if c.coord.x == 0 && c.coord.y == 0 {
		dirs = []Direction{East, South}
	} else if c.horizontal {
		dirs = []Direction{West, East}
	} else {
		dirs = []Direction{North, South}
	}
	for _, dir := range dirs {
		for i := g.minStep; i <= g.maxStep; i++ {
			newCoord := Coord{c.coord.x + dir.dx*i, c.coord.y + dir.dy*i}
			if g.grid.isValidCoord(newCoord) {
				// make sure the target node is always the same (arbitrarily: horizontal)
				isTarget := newCoord.x == len(g.grid[0])-1 && newCoord.y == len(g.grid)-1
				coords = append(coords, SearchNode{newCoord, dir.dx == 0 && !isTarget})
			}
		}
	}
	return coords
}

func ReadHeatLossGrid(path string) HeatGrid {
	lines := ReadFile(path)

	var grid [][]int
	for _, line := range lines {
		if line == "" {
			continue
		}
		var gridLine []int
		for _, r := range line {
			gridLine = append(gridLine, int(r-'0'))
		}
		grid = append(grid, gridLine)
	}
	return grid
}

func estimatedCost(c0, c1 SearchNode) float64 {
	return float64(Dist(c0.coord, c1.coord))
}

func (g HeatGrid) cost(c0, c1 SearchNode) float64 {
	dir := ToDir(c0.coord, c1.coord)
	length := dir.Length()
	unitDir := Direction{dir.dx / length, dir.dy / length}

	cost := 0
	for i := 1; i <= length; i++ {
		c := Coord{c0.coord.x + unitDir.dx*i, c0.coord.y + unitDir.dy*i}
		cost += g.GetLoss(c)
	}
	return float64(cost)
}

func toCoords(ns []SearchNode) []Coord {
	coords := []Coord{}
	for _, n := range ns {
		coords = append(coords, n.coord)
	}
	return coords
}

func FindLeastLossPath(grid HeatGrid, minStep int, maxStep int) ([]Coord, int) {
	mx, my := grid.GetBounds()
	from := SearchNode{Coord{0, 0}, false}
	to := SearchNode{Coord{mx - 1, my - 1}, false}

	path := astar.FindPath(SearchGrid{grid, minStep, maxStep}, from, to, grid.cost, estimatedCost)

	return toCoords(path), int(path.Cost(grid.cost))
}
