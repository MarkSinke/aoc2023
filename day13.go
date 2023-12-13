package aoc2023

import (
	"fmt"
	"strings"
)

type Note []string

func ReadMirrorNotes(path string) []Note {
	lines := ReadFile(path)

	var notes []Note

	var note Note
	for _, line := range lines {
		if line == "" {
			if len(note) > 0 {
				notes = append(notes, note)
				note = []string{}
			}
		} else {
			note = append(note, line)
		}
	}

	if len(note) > 0 {
		notes = append(notes, note)
	}

	return notes
}

func findHorizontalMirror(n Note) int {
	for i := 0; i < len(n)-1; i++ {
		if n[i] == n[i+1] && hasMirror(n, i) {
			return i
		}
	}

	return -1
}

func hasMirror(n Note, mirrorLine int) bool {
	for i, j := mirrorLine, mirrorLine+1; i >= 0 && j < len(n); i, j = i-1, j+1 {
		if n[i] != n[j] {
			return false
		}
	}
	return true
}

func transpose(n Note) Note {
	res := Note(make([]string, len(n[0])))

	for _, line := range n {
		for x, rune := range line {
			res[x] += string(rune)
		}
	}
	return res
}

func (n Note) GetMirrorLine() int {
	mirror := findHorizontalMirror(n)

	if mirror != -1 {
		return 100 * (mirror + 1)
	}

	mirror = findHorizontalMirror(transpose(n))
	if mirror == -1 {
		return -1
	}
	return mirror + 1
}

func (n Note) Print() {
	for _, line := range n {
		fmt.Println(line)
	}
}

func (n Note) PrintMirror(mirror int) {
	if mirror >= 100 {
		yM := (mirror - 1) / 100
		for y, line := range n {
			fmt.Println(line)
			if y == yM {
				fmt.Println(strings.Repeat("-", len(line)))
			}
		}
	} else {
		for _, line := range n {
			fmt.Print(line[:mirror])
			fmt.Print("|")
			fmt.Println(line[mirror:])
		}
	}
}

func SumMirrors(notes []Note) int {
	sum := 0
	for _, n := range notes {
		mirror := n.GetMirrorLine()
		sum += mirror
	}
	return sum
}

func SumSmudgeMirrors(notes []Note) int {
	sum := 0
	for _, n := range notes {
		mirror := n.GetSmudgeMirrorLine()
		sum += mirror
	}
	return sum
}

func flip(r byte) byte {
	if r == '#' {
		return '.'
	}
	return '#'
}

func (n Note) Unsmudge(y int, x int) Note {
	var res []string
	for yRes, line := range n {
		if yRes == y {
			line = line[:x] + string(flip(line[x])) + line[x+1:]
		}
		res = append(res, line)
	}
	return res
}

func findDiffIndex(str0 string, str1 string) int {
	index := -1
	for i := range str0 {
		if str0[i] != str1[i] {
			if index == -1 {
				index = i
			} else {
				return -1
			}
		}
	}
	return index
}

func findHorizontalMirrorSmudge(n Note) int {
	for y := range n {
		for y2 := y + 1; y2 < len(n); y2 += 2 {
			x := findDiffIndex(n[y], n[y2])
			mirrorLine := (y + y2) / 2
			if x != -1 && hasMirror(n.Unsmudge(y, x), mirrorLine) {
				return mirrorLine
			}
		}
	}
	return -1
}

func (n Note) GetSmudgeMirrorLine() int {
	mirror := findHorizontalMirrorSmudge(n)

	if mirror != -1 {
		return 100 * (mirror + 1)
	}

	mirror = findHorizontalMirrorSmudge(transpose(n))
	if mirror == -1 {
		// panic("0")
		return -1
	}
	return mirror + 1
}
