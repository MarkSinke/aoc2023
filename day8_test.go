package aoc2023

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay8InputExample(t *testing.T) {
	lr, nodes := ReadLRMap("day8input_example.txt")

	steps := StepsToZzz(lr, nodes)

	assert.Equal(t, 6, steps)
}

func TestDay8Input(t *testing.T) {
	lr, nodes := ReadLRMap("day8input.txt")

	steps := StepsToZzz(lr, nodes)

	fmt.Println("Day 8 result (star 1):", steps)
}

func TestDay8Star2InputExample(t *testing.T) {
	lr, nodes := ReadLRMap("day8input_example2.txt")

	steps := StepsToZzzParallel(lr, nodes)

	assert.Equal(t, 6, steps)
}

func TestDay8Star2Input(t *testing.T) {
	lr, nodes := ReadLRMap("day8input.txt")

	steps := StepsToZzzParallel(lr, nodes)

	fmt.Println("Day 8 result (star 2):", steps)
}
