package aoc2023

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay19InputExample(t *testing.T) {
	program, parts := ReadWorkflowsAndParts("day19input_example.txt")

	accepted := ExecuteProgram(program, parts)

	sum := SumRatings(accepted)

	assert.Equal(t, 19114, sum)
}

func TestDay19Input(t *testing.T) {
	program, parts := ReadWorkflowsAndParts("day19input.txt")

	accepted := ExecuteProgram(program, parts)

	sum := SumRatings(accepted)

	fmt.Println("Day 19 result (star 1):", sum)
}

func TestDay19Star2InputExample(t *testing.T) {
	program, _ := ReadWorkflowsAndParts("day19input_example.txt")

	accepted := InspectProgram(program)

	sum := SumCombinations(accepted)

	assert.Equal(t, 167409079868000, sum)
}

func TestDay19Star2Input(t *testing.T) {
	program, _ := ReadWorkflowsAndParts("day19input.txt")

	accepted := InspectProgram(program)

	sum := SumCombinations(accepted)

	fmt.Println("Day 19 result (star 2):", sum)
}
