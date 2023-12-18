package aoc2023

import (
	"container/heap"
	"fmt"
	"math"
)

type Heat struct {
	loss         int
	minTotalLoss int
}

func ReadHeatLossGrid(path string) [][]Heat {
	lines := ReadFile(path)

	var grid [][]Heat
	for _, line := range lines {
		if line == "" {
			continue
		}
		var gridLine []Heat
		for _, r := range line {
			gridLine = append(gridLine, Heat{int(r - '0'), math.MaxInt})
		}
		grid = append(grid, gridLine)
	}
	return grid
}

func isUTurn(prev []Direction, dir Direction) bool {
	if len(prev) == 0 {
		return false
	}
	prevDir := prev[len(prev)-1]
	return dir.dx+prevDir.dx == 0 && dir.dy+prevDir.dy == 0
}

func isCycle(prevs []Direction, dir Direction) bool {
	dx := dir.dx
	dy := dir.dy

	for _, prev := range prevs {
		dx += prev.dx
		dy += prev.dy

		if dx == 0 && dy == 0 {
			return true
		}
	}
	return false
}

func isOutOfBounds(c Coord, grid [][]Heat) bool {
	return c.x < 0 || c.x >= len(grid[0]) || c.y < 0 || c.y >= len(grid)
}

func isFourthInOneDirection(prev []Direction, dir Direction) bool {
	count := len(prev)
	return count >= 3 && prev[count-3] == dir && prev[count-2] == dir && prev[count-1] == dir
}

// func isAnotherPathAtLeastAsEfficient(loss int, grid [][]Heat, c Coord) bool {
// 	return grid[c.y][c.x].minTotalLoss < loss-1
// }

type HeatPathState struct {
	steps []Direction
	cur   Coord
	loss  int
}

type HeatQueue []HeatPathState

func (hq HeatQueue) Len() int {
	return len(hq)
}

func (hq HeatQueue) Less(i, j int) bool {
	return hq[i].loss < hq[j].loss
}

func (hq HeatQueue) Swap(i, j int) {
	hq[i], hq[j] = hq[j], hq[i]
}

func (hq *HeatQueue) Push(x any) {
	item := x.(HeatPathState)
	*hq = append(*hq, item)
}

func (hq *HeatQueue) Pop() any {
	old := *hq
	n := len(old)
	item := old[n-1]
	*hq = old[0 : n-1]
	return item
}

func copyAndAppend(prev []Direction, dir Direction) []Direction {
	steps := make([]Direction, len(prev)+1)
	copy(steps, prev)
	steps[len(steps)-1] = dir
	return steps
}

func heatLossPathStep(prevSteps []Direction, c Coord, loss int, dir Direction, grid [][]Heat) *HeatPathState {
	// fmt.Print("step", dir, "-> ")
	// if isUTurn(prevSteps, dir) {
	// 	// fmt.Println("u-turn")
	// 	return nil
	// }
	if isCycle(prevSteps, dir) {
		return nil
	}

	if isFourthInOneDirection(prevSteps, dir) {
		// fmt.Println("4th move")
		return nil
	}

	newCoord := Coord{c.x + dir.dx, c.y + dir.dy}

	if isOutOfBounds(newCoord, grid) {
		// fmt.Println("oob", newCoord)
		return nil
	}

	loss += grid[newCoord.y][newCoord.x].loss

	// if isAnotherPathAtLeastAsEfficient(loss, grid, newCoord) {
	// 	fmt.Println("less efficient")
	// 	return nil
	// }
	// grid[newCoord.y][newCoord.x].minTotalLoss = loss

	return &HeatPathState{copyAndAppend(prevSteps, dir), newCoord, loss}
}

func FindLeastLossPath(grid [][]Heat) ([]Direction, int) {
	states := &HeatQueue{{[]Direction{}, Coord{0, 0}, 0}}
	heap.Init(states)

	for len(*states) > 0 {
		state := heap.Pop(states).(HeatPathState)
		fmt.Println("popped", state.loss, "steps", len(state.steps), "from", len(*states))
		if state.cur.x == len(grid[0])-1 && state.cur.y == len(grid)-1 {
			return state.steps, state.loss
		}

		for _, dir := range AllDirections {
			newState := heatLossPathStep(state.steps, state.cur, state.loss, dir, grid)
			if newState != nil {
				heap.Push(states, *newState)
				// fmt.Println("pushed", *newState)
			}
		}
	}
	panic("min loss in the grid should correspond with one of the paths")
}

func PrintPath(steps []Direction, grid [][]Heat) {
	cur := Coord{0, 0}
	loss := 0
	for _, step := range steps {
		fmt.Print(cur, "+", step)
		cur = Coord{cur.x + step.dx, cur.y + step.dy}
		loss += grid[cur.y][cur.x].loss
		fmt.Println(" ->", cur, "loss=", loss)
	}
}
