package aoc2023

import (
	"fmt"
	"slices"

	"github.com/fzipp/astar"
)

type Grid interface {
	GetLoss(c Coord) int
}

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

func (g HeatGrid) Neighbours(c Coord) []Coord {
	coords := []Coord{}
	for _, dir := range AllDirections {
		newCoord := Coord{c.x + dir.dx, c.y + dir.dy}
		if g.isValidCoord(newCoord) {
			coords = append(coords, newCoord)
		}
	}
	return coords
}

type PartialGrid struct {
	g              HeatGrid
	forbiddenNodes []Coord
	spurNode       Coord
	forbiddenDirs  []Direction
}

func (g PartialGrid) GetLoss(c Coord) int {
	return g.g.GetLoss(c)
}

func (g PartialGrid) isHiddenEdge(from Coord, to Coord) bool {
	if slices.Contains(g.forbiddenNodes, to) {
		return true
	}
	if from == g.spurNode {
		dir := ToDir(from, to)
		if slices.Contains(g.forbiddenDirs, dir) {
			return true
		}
	}
	return false
}

func (g PartialGrid) Neighbours(c Coord) []Coord {
	coords := g.g.Neighbours(c)

	filtered := []Coord{}

	for _, coord := range coords {
		if !g.isHiddenEdge(c, coord) {
			filtered = append(filtered, coord)
		}
	}
	return filtered
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

func comparePaths(e HeatPath, target HeatPath) int {
	diff := e.loss - target.loss
	if diff != 0 {
		return diff
	}

	i := 0
	for ; i < len(e.coords) && i < len(target.coords); i++ {
		diff = e.coords[i].y - target.coords[i].y
		if diff != 0 {
			return diff
		}
		diff = e.coords[i].x - target.coords[i].x
		if diff != 0 {
			return diff
		}
	}
	diff = len(e.coords) - len(target.coords)
	return diff
}

func estimatedCost(c0, c1 Coord) float64 {
	return float64(Dist(c0, c1))
}

func (g HeatGrid) cost(c0, c1 Coord) float64 {
	return float64(g.GetLoss(c1))
}

func FindLeastLossPath(grid HeatGrid, from Coord, to Coord) HeatPath {
	path := astar.FindPath[Coord](grid, from, to, grid.cost, estimatedCost)

	heatPath := HeatPath{path, int(path.Cost(grid.cost))}

	if isAllowedPath(heatPath) {
		return heatPath
	}
	// fmt.Println("A", "0", "is", heatPath)

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

			spurPath := astar.FindPath(pg, spurNode, to, grid.cost, estimatedCost)
			if spurPath != nil {
				totalPath := append([]Coord{}, forbiddenNodes...)
				totalPath = append(totalPath, spurPath...)
				totalLoss := ComputeTotalLoss(totalPath, grid)
				newPath := HeatPath{totalPath, totalLoss}
				// fmt.Println("new path", newPath)
				// fmt.Println("enqueued are")
				// for _, cp := range candidatePaths {
				// 	fmt.Println(cp)
				// }
				// fmt.Println()
				index, found := slices.BinarySearchFunc(candidatePaths, newPath, comparePaths)
				// fmt.Println("index", index)
				if !found {
					candidatePaths = slices.Insert(candidatePaths, index, newPath)
				}
			}
		}

		newShortest := candidatePaths[0]
		if isAllowedPath(newShortest) {
			return newShortest
		}

		shortestPaths = append(shortestPaths, newShortest)
		candidatePaths = candidatePaths[1:]

		// fmt.Println("A", k, "is", newShortest)
		// fmt.Println("enqueued are")
		// for _, cp := range candidatePaths {
		// 	fmt.Println(cp)
		// }
		// fmt.Println()
	}
}

func PrintPath(steps []Coord, grid Grid) {
	loss := 0
	for _, step := range steps[1:] {
		fmt.Print(step)
		loss += grid.GetLoss(step)
		fmt.Println(" -> loss=", loss)
	}
}
