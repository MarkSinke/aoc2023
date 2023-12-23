package aoc2023

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay22InputExample(t *testing.T) {
	bricks := ReadBricks("day22input_example.txt")

	DropBricks(bricks)
	res := CountDisintegratable(bricks)

	assert.Equal(t, 5, res)
}

func TestDay22Input(t *testing.T) {
	bricks := ReadBricks("day22input.txt")

	DropBricks(bricks)
	res := CountDisintegratable(bricks)

	fmt.Println("Day 22 result (star 1):", res)
}
func TestDay22Star2InputExample(t *testing.T) {
	bricks := ReadBricks("day22input_example.txt")

	DropBricks(bricks)
	res := CountDisintegrationFalls(bricks)

	assert.Equal(t, 7, res)
}

func TestDay22Star2Input(t *testing.T) {
	bricks := ReadBricks("day22input.txt")

	DropBricks(bricks)
	res := CountDisintegrationFalls(bricks)

	fmt.Println("Day 22 result (star 2):", res)
}
