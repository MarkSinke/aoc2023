package aoc2023

import "fmt"

type Rock int

const (
	round Rock = iota
	square
	none
)

func (r Rock) String() string {
	switch r {
	case round:
		return "O"
	case square:
		return "#"
	case none:
		return "."
	}
	return "?"
}

func toRock(r rune) Rock {
	switch r {
	case 'O':
		return round
	case '#':
		return square
	case '.':
		return none
	}
	return none
}

func ReadRockPositions(path string) [][]Rock {
	lines := ReadFile(path)

	var rocks [][]Rock
	for _, line := range lines {
		if line == "" {
			continue
		}
		var rockLine []Rock
		for _, rune := range line {
			rockLine = append(rockLine, toRock(rune))
		}
		rocks = append(rocks, rockLine)
	}

	return rocks
}

func PrintRocks(rocks [][]Rock) {
	for _, rockLine := range rocks {
		for _, rock := range rockLine {
			fmt.Print(rock.String())
		}
		fmt.Println()
	}
	fmt.Println()
}

func findDySouth(x int, y int, rocks [][]Rock) int {
	dy := 1
	for dy < len(rocks)-y && rocks[y+dy][x] == none {
		dy++
	}
	return dy - 1
}

func findDyNorth(x int, y int, rocks [][]Rock) int {
	dy := 1
	for dy <= y && rocks[y-dy][x] == none {
		dy++
	}
	return dy - 1
}

func findDxEast(x int, y int, rocks [][]Rock) int {
	dx := 1
	for dx < len(rocks[0])-x && rocks[y][x+dx] == none {
		dx++
	}
	return dx - 1
}

func findDxWest(x int, y int, rocks [][]Rock) int {
	dx := 1
	for dx <= x && rocks[y][x-dx] == none {
		dx++
	}
	return dx - 1
}

func TiltSouth(rocks [][]Rock) {
	for y := len(rocks) - 2; y >= 0; y-- {
		for x := 0; x < len(rocks[0]); x++ {
			if rocks[y][x] == round {
				dy := findDySouth(x, y, rocks)
				if dy > 0 {
					rocks[y][x] = none
					rocks[y+dy][x] = round
				}
			}
		}
	}
}

func TiltNorth(rocks [][]Rock) {
	for y := 1; y < len(rocks); y++ {
		for x := 0; x < len(rocks[0]); x++ {
			if rocks[y][x] == round {
				dy := findDyNorth(x, y, rocks)
				if dy > 0 {
					rocks[y][x] = none
					rocks[y-dy][x] = round
				}
			}
		}
	}
}

func TiltEast(rocks [][]Rock) {
	for x := len(rocks[0]) - 2; x >= 0; x-- {
		for y := 0; y < len(rocks); y++ {
			if rocks[y][x] == round {
				dx := findDxEast(x, y, rocks)
				if dx > 0 {
					rocks[y][x] = none
					rocks[y][x+dx] = round
				}
			}
		}
	}
}

func TiltWest(rocks [][]Rock) {
	for x := 1; x < len(rocks[0]); x++ {
		for y := 0; y < len(rocks); y++ {
			if rocks[y][x] == round {
				dx := findDxWest(x, y, rocks)
				if dx > 0 {
					rocks[y][x] = none
					rocks[y][x-dx] = round
				}
			}
		}
	}
}

func SpinRocks(rocks [][]Rock) {
	TiltNorth(rocks)
	TiltWest(rocks)
	TiltSouth(rocks)
	TiltEast(rocks)
}

func isEqual(left [][]Rock, right [][]Rock) bool {
	for y, line := range left {
		for x, rock := range line {
			if rock != right[y][x] {
				return false
			}
		}
	}
	return true
}

func findIndex(prevs [][][]Rock, now [][]Rock) int {
	for i := len(prevs) - 1; i >= 0; i-- {
		if isEqual(prevs[i], now) {
			return i
		}
	}
	return -1
}

func SpinUntilStable(rocks [][]Rock) (int, int) {
	var prev [][][]Rock
	for {
		prev = append(prev, CopyRocks(rocks))
		SpinRocks(rocks)

		prevIndex := findIndex(prev, rocks)
		if prevIndex != -1 {
			// found a cycle, starting a prevIndex, and running until the last spin result
			return prevIndex, len(prev) - prevIndex
		}
	}
}

func GetNorthWeight(rocks [][]Rock) int {
	weight := 0
	for y, rockLine := range rocks {
		for _, rock := range rockLine {
			if rock == round {
				weight += len(rocks) - y
			}
		}
	}
	return weight
}

func CopyRocks(rocks [][]Rock) [][]Rock {
	var res [][]Rock
	for _, rockLine := range rocks {
		copyLine := make([]Rock, len(rockLine))
		copy(copyLine, rockLine)
		res = append(res, copyLine)
	}
	return res
}

func SpinOneBillionTimes(rocks [][]Rock) {
	start, cycleLength := SpinUntilStable(rocks)
	maxSpins := 1000000000

	remainingSpins := maxSpins - cycleLength - start
	remainingCycles := remainingSpins / cycleLength
	remainingSpins -= remainingCycles * cycleLength

	for i := 0; i < remainingSpins; i++ {
		SpinRocks(rocks)
	}
}
