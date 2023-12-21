package aoc2023

import (
	"regexp"
	"strings"

	"golang.org/x/exp/maps"
)

type Node struct {
	left  string
	right string
}

var nodeRegexp = regexp.MustCompile(`([0-9A-Z]+) = \(([0-9A-Z]+), ([0-9A-Z]+)\)`)

func ReadLRMap(path string) (string, map[string]Node) {
	lines := ReadFile(path)

	lr := lines[0]
	nodes := map[string]Node{}

	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		parts := nodeRegexp.FindStringSubmatch(line)
		nodes[parts[1]] = Node{parts[2], parts[3]}
	}

	return lr, nodes
}

func StepsToZzz(lr string, m map[string]Node) int {
	location := "AAA"
	steps := 0
	lrIndex := 0

	for {
		if location == "ZZZ" {
			return steps
		}

		if lr[lrIndex] == 'L' {
			location = m[location].left
		} else {
			location = m[location].right
		}
		steps++

		lrIndex = (lrIndex + 1) % len(lr)
	}
}

func StepsToZ(location string, lr string, m map[string]Node) int {
	steps := 0
	lrIndex := 0

	for {
		if strings.HasSuffix(location, "Z") {
			return steps
		}

		if lr[lrIndex] == 'L' {
			location = m[location].left
		} else {
			location = m[location].right
		}
		steps++

		lrIndex = (lrIndex + 1) % len(lr)
	}
}

func filterAllEndingInA(locations []string) []string {
	result := []string{}

	for _, loc := range locations {
		if strings.HasSuffix(loc, "A") {
			result = append(result, loc)
		}
	}
	return result
}

// greatest common divisor
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func lcm2(a int, b int) int {
	return a * b / gcd(a, b)
}

func LeastCommonMultiple(ints ...int) int {
	switch len(ints) {
	case 0:
		return 0
	case 1:
		return ints[0]
	}
	result := lcm2(ints[0], ints[1])

	for _, val := range ints[2:] {
		result = lcm2(result, val)
	}

	return result
}

func StepsToZzzParallel(lr string, m map[string]Node) int {
	locations := filterAllEndingInA(maps.Keys(m))
	stepsOne := []int{}
	for _, loc := range locations {
		stepsOne = append(stepsOne, StepsToZ(loc, lr, m))
	}
	return LeastCommonMultiple(stepsOne...)
}
