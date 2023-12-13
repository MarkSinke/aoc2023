package aoc2023

import "fmt"

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

	return findHorizontalMirror(transpose(n)) + 1
}

func (n Note) Print() {
	for _, line := range n {
		fmt.Println(line)
	}
}

func SumMirrors(notes []Note) int {
	sum := 0
	for _, n := range notes {
		sum += n.GetMirrorLine()
	}
	return sum
}
