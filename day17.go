package aoc2023

import (
	"fmt"
	"slices"
)

type Grid interface {
	GetLoss(c Coord) int
	IsValidStep(from Coord, to Coord) bool
}

type HeatGrid [][]int

func (g HeatGrid) GetBounds() (int, int) {
	return len(g[0]), len(g)
}

func (g HeatGrid) GetLoss(c Coord) int {
	return g[c.y][c.x]
}

func (g HeatGrid) IsValidStep(from Coord, to Coord) bool {
	mx, my := g.GetBounds()
	return to.x >= 0 && to.x < mx && to.y >= 0 && to.y < my
}

type PartialGrid struct {
	g              Grid
	forbiddenNodes []Coord
	spurNode       Coord
	forbiddenDirs  []Direction
}

func (g PartialGrid) GetLoss(c Coord) int {
	return g.g.GetLoss(c)
}

func (g PartialGrid) IsValidStep(from Coord, to Coord) bool {
	if !g.g.IsValidStep(from, to) {
		return false
	}
	if slices.Contains(g.forbiddenNodes, to) {
		return false
	}
	if from == g.spurNode {
		dir := ToDir(from, to)
		if slices.Contains(g.forbiddenDirs, dir) {
			return false
		}
	}
	return true
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

// HeatTile is a struct which implements Pather
type HeatTile struct {
	c    Coord
	grid Grid
}

func (t HeatTile) PathEstimatedCost(to Pather) int {
	return Dist(t.c, to.(HeatTile).c)
}

func (t HeatTile) PathNeighborCost(p Pather) int {
	to := p.(HeatTile)
	return to.grid.GetLoss(to.c)
}

func (t HeatTile) Coord() Coord {
	return t.c
}

func (t HeatTile) PathNeighbors() []Pather {
	neighbors := []Pather{}
	for _, dir := range AllDirections {
		newCoord := Coord{t.c.x + dir.dx, t.c.y + dir.dy}
		if t.grid.IsValidStep(t.c, newCoord) {
			neighbors = append(neighbors, HeatTile{newCoord, t.grid})
		}
	}
	return neighbors
}

type HeatPath struct {
	coords []Coord
	loss   int
}

func isAllowedPath(path HeatPath) bool {
	coords := path.coords
	for i := 4; i < len(coords); i++ {
		c := coords[i]
		d := ToDir(coords[i-1], c)
		if ToDir(coords[i-4], coords[i-3]) == d && ToDir(coords[i-3], coords[i-2]) == d && ToDir(coords[i-2], coords[i-1]) == d {
			return false
		}
	}
	return true
}

func hasSameRoot(path HeatPath, coords []Coord) bool {
	for i := 0; i < len(coords); i++ {
		if path.coords[i] != coords[i] {
			return false
		}
	}
	return true
}

func ComputeTotalLoss(coords []Coord, g Grid) int {
	loss := 0
	for _, c := range coords[1:] {
		loss += g.GetLoss(c)
	}
	return loss
}

func comparePaths(left HeatPath, right HeatPath) int {
	diff := right.loss - left.loss
	if diff != 0 {
		return diff
	}

	i := 0
	for ; i < len(left.coords) && i < len(right.coords); i++ {
		diff = right.coords[i].y - left.coords[i].y
		if diff != 0 {
			return diff
		}
		diff = right.coords[i].x - left.coords[i].x
		if diff != 0 {
			return diff
		}
	}
	diff = len(right.coords) - len(left.coords)
	return diff
}

func FindLeastLossPath(grid Grid, from Coord, to Coord) HeatPath {
	path, distance, _ := Path(HeatTile{from, grid}, HeatTile{to, grid})
	heatPath := HeatPath{path, distance}

	if isAllowedPath(heatPath) {
		return heatPath
	}

	shortestPaths := []HeatPath{heatPath}
	var candidatePaths []HeatPath

	// use Yen's algorithm to find the k shortest paths (in this case, next shortest) - from Wikipedia
	for k := 1; ; k++ {
		fmt.Println("k=", k)
		for i := 0; i < len(shortestPaths[k-1].coords)-2; i++ {
			rootPath := shortestPaths[k-1].coords[0 : i+1] // includes spurNode
			spurNode := shortestPaths[k-1].coords[i]

			// all directions from the spur nodes that have already been tried in previous shortest paths
			forbiddenNodes := shortestPaths[k-1].coords[0:i]
			forbiddenDirs := []Direction{}
			for _, path := range shortestPaths {
				if hasSameRoot(path, rootPath) {
					forbiddenDirs = append(forbiddenDirs, ToDir(path.coords[i], path.coords[i+1]))
				}
			}

			pg := PartialGrid{grid, forbiddenNodes, spurNode, forbiddenDirs}

			spurPath, _, found := Path(HeatTile{spurNode, pg}, HeatTile{to, pg})
			if found {
				totalPath := append([]Coord{}, forbiddenNodes...)
				totalPath = append(totalPath, spurPath...)
				totalLoss := ComputeTotalLoss(totalPath, grid)
				newPath := HeatPath{totalPath, totalLoss}
				// fmt.Println("new path", newPath)
				index, found := slices.BinarySearchFunc(candidatePaths, newPath, comparePaths)
				if !found {
					candidatePaths = slices.Insert(candidatePaths, index, newPath)
				}
			}
		}

		newShortest := candidatePaths[0]
		fmt.Println("new shortest", newShortest.loss, newShortest)
		if isAllowedPath(newShortest) {
			return newShortest
		}

		shortestPaths = append(shortestPaths, newShortest)
		candidatePaths = candidatePaths[1:]
	}
}

func PrintPath(steps []Coord, grid Grid) {
	loss := 0
	for _, step := range steps {
		fmt.Print(step)
		loss += grid.GetLoss(step)
		fmt.Println(" -> loss=", loss)
	}
}
