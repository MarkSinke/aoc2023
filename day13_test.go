package aoc2023

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay13InputExample(t *testing.T) {
	notes := ReadMirrorNotes("day13input_example.txt")

	sum := SumMirrors(notes)

	assert.Equal(t, 405, sum)
}

func TestDay13Input(t *testing.T) {
	notes := ReadMirrorNotes("day13input.txt")

	sum := SumMirrors(notes)

	fmt.Println("Day 13 result (star 1):", sum)
}
