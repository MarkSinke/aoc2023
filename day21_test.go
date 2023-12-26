package aoc2023

import (
	"fmt"
	"math"
	"testing"

	"github.com/cryptowilliam/goutil/container/gpoly/polyfit"
	"github.com/stretchr/testify/assert"
)

func TestDay21InputExample(t *testing.T) {
	garden, start := ReadGarden("day21input_example.txt")

	WalkGarden(garden, start, 6)

	count := CountVisited(garden)

	assert.Equal(t, 16, count)
}

func TestDay21Input(t *testing.T) {
	garden, start := ReadGarden("../aoc_inputs/2023/day21input.txt")

	WalkGarden(garden, start, 64)

	count := CountVisited(garden)

	fmt.Println("Day 21 result (star 1):", count)
}

func TestDay21Star2Input(t *testing.T) {
	garden, _ := ReadGarden("../aoc_inputs/2023/day21input.txt")

	// for 65 steps, we exacty reach the edge of the original 131x131 grid
	// (the center is excluded, as it's an odd number), every 131 steps yields a next tile
	// there is a a clear diagonal of . in that position
	// so we compute a few data points and then extrapolate (quadratically, as it's area growth)
	steps := 26501365
	largeGarden := ExplodeGarden(garden, 13)

	xs := []float64{}
	ys := []float64{}
	// increasing to 5 did not change coefficients
	for i := 0; i < 4; i++ {
		d := i*131 + 65
		WalkGarden(largeGarden, Coord{d, d}, d)
		count := CountVisited(largeGarden)
		ResetGarden(largeGarden)

		xs = append(xs, float64(i))
		ys = append(ys, float64(count))
	}

	// fit a polynomial through these numbers (quadratic), and compute f((26501365-65)/131)
	// we ask for cubic (degree = 3), and then check that the cubic coefficient is zero
	f := polyfit.NewFitting(xs, ys, 3)
	fCoeff := f.Solve(false)

	// // cubic exponent should be very small
	assert.LessOrEqual(t, math.Abs(fCoeff[3]), 1e-9)

	var iCoeff [3]int
	iCoeff[0] = int(math.Round(fCoeff[0]))
	iCoeff[1] = int(math.Round(fCoeff[1]))
	iCoeff[2] = int(math.Round(fCoeff[2]))

	// coefficients should be integers (barring a few floating point errors)
	assert.LessOrEqual(t, math.Abs(float64(iCoeff[0])-fCoeff[0]), 1e-6)
	assert.LessOrEqual(t, math.Abs(float64(iCoeff[1])-fCoeff[1]), 1e-6)
	assert.LessOrEqual(t, math.Abs(float64(iCoeff[2])-fCoeff[2]), 1e-6)

	giantSteps := (steps - 65) / 131
	result := iCoeff[0] + iCoeff[1]*giantSteps + iCoeff[2]*giantSteps*giantSteps

	fmt.Println("Day 21 result (star 2):", result)
}
