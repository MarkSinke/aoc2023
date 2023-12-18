package aoc2023

import (
	"slices"
)

func ReadGalaxies(path string) []Coord {
	lines := ReadFile(path)

	coords := []Coord{}

	for y, line := range lines {
		if line == "" {
			continue
		}
		for x, r := range line {
			if r == '#' {
				coords = append(coords, Coord{x, y})
			}
		}
	}
	return coords
}

func updateCoord(xs []int, ys []int, coord Coord, expand int) Coord {
	i, _ := slices.BinarySearch(xs, coord.x)
	dx := coord.x - i
	j, _ := slices.BinarySearch(ys, coord.y)
	dy := coord.y - j
	return Coord{coord.x + dx*expand, coord.y + dy*expand}
}

func ExpandGalaxy(coords []Coord, d int) []Coord {
	xs := make([]int, 0, len(coords))
	ys := make([]int, 0, len(coords))

	for _, coord := range coords {
		xs = append(xs, coord.x)
		ys = append(ys, coord.y)
	}

	slices.Sort(xs)
	slices.Sort(ys)
	xs = slices.Compact(xs)
	ys = slices.Compact(ys)

	res := make([]Coord, 0, len(coords))
	for _, coord := range coords {
		res = append(res, updateCoord(xs, ys, coord, d))
	}

	return res
}

func ComputeDistances(coords []Coord) []int {
	res := []int{}
	for i, coord0 := range coords {
		for _, coord1 := range coords[0:i] {
			res = append(res, Dist(coord0, coord1))
		}
	}
	return res
}
