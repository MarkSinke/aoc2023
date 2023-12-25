package aoc2023

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay25InputExample(t *testing.T) {
	nodes := ReadComponentGraph("day25input_example.txt")

	a, b := PartitionGraph(nodes)

	sizeA := CountComponentSize(nodes, a)
	sizeB := CountComponentSize(nodes, b)
	res := sizeA * sizeB

	assert.Equal(t, 54, res)
}

func TestDay25Input(t *testing.T) {
	nodes := ReadComponentGraph("day25input.txt")

	a, b := PartitionGraph(nodes)

	sizeA := CountComponentSize(nodes, a)
	sizeB := CountComponentSize(nodes, b)
	res := sizeA * sizeB

	fmt.Println("Day 25 result: (star 1):", res)
	assert.Equal(t, 556467, res)
}
