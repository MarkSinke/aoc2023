package aoc2023

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func press1000(graph GateGraph) (int, int) {
	lowSum, highSum := 0, 0

	for i := 0; i < 1000; i++ {
		low, high := ExecNetwork(graph)
		lowSum += low
		highSum += high
	}

	return lowSum, highSum
}

func TestDay20InputExample1(t *testing.T) {
	graph := ReadNetwork("day20input_example.txt")

	lowSum, highSum := press1000(graph)

	assert.Equal(t, 32000000, lowSum*highSum)
}

func TestDay20InputExample2(t *testing.T) {
	graph := ReadNetwork("day20input_example2.txt")

	lowSum, highSum := press1000(graph)

	assert.Equal(t, 11687500, lowSum*highSum)
}

func TestDay20Input(t *testing.T) {
	graph := ReadNetwork("day20input.txt")

	lowSum, highSum := press1000(graph)

	fmt.Println("Day 20 result (star 1):", lowSum*highSum)
}

func TestDay20Star2Input(t *testing.T) {
	graph := ReadNetwork("day20input.txt")

	// manual analysis of the particular network shows that tx, dd, nz, and ph
	// are inverters that feed into the final conjunction ls. Hence these need
	// all to have low inputs before ls output (=rx input) turns low.
	countTx := PressUntilLowRx(graph, "tx", false)
	graph.Reset()
	countDd := PressUntilLowRx(graph, "dd", false)
	graph.Reset()
	countNz := PressUntilLowRx(graph, "nz", false)
	graph.Reset()
	countPh := PressUntilLowRx(graph, "ph", false)
	graph.Reset()

	count := LeastCommonMultiple(countTx, countDd, countNz, countPh)

	fmt.Println("Day 20 result (star 2):", count)
}
