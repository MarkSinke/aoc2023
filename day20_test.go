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
