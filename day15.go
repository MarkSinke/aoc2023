package aoc2023

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hash(str string) int {
	val := 0
	for _, r := range str {
		val += int(r)
		val *= 17
		val %= 256
	}
	return val
}

func ReadInitInstructions(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic("cannot read file " + path)
	}
	str, _ := strings.CutSuffix(string(data), "\n")
	return strings.Split(str, ",")
}

func SumHashes(strs []string) int {
	sum := 0
	for _, str := range strs {
		sum += hash(str)
	}
	return sum
}

type Lens struct {
	label string
	focal int
}

type Box []Lens

func (b Box) remove(label string) Box {
	var res Box
	for _, lens := range b {
		if lens.label != label {
			res = append(res, lens)
		}
	}
	return res
}

func (b Box) replace(label string, focal int) Box {
	var res Box
	replaced := false
	for _, lens := range b {
		if lens.label == label {
			res = append(res, Lens{label, focal})
			replaced = true
		} else {
			res = append(res, lens)
		}
	}
	if !replaced {
		res = append(res, Lens{label, focal})
	}
	return res
}

func remove(boxes []Box, label string) {
	boxIndex := hash(label)
	boxes[boxIndex] = boxes[boxIndex].remove(label)
}

func replace(boxes []Box, label string, focal int) {
	boxIndex := hash(label)
	boxes[boxIndex] = boxes[boxIndex].replace(label, focal)
}

func RunInit(instructions []string) []Box {
	boxes := make([]Box, 256)

	for _, inst := range instructions {
		minusStr, minusFound := strings.CutSuffix(inst, "-")
		if minusFound {
			remove(boxes, minusStr)
		} else {
			label, focalStr, _ := strings.Cut(inst, "=")
			focal, _ := strconv.Atoi(focalStr)
			replace(boxes, label, focal)
		}
	}

	return boxes
}

func PrintBoxes(boxes []Box) {
	for i, box := range boxes {
		if len(box) > 0 {
			fmt.Println("Box", i, box)
		}
	}
}

func SumFocusingPower(boxes []Box) int {
	sum := 0
	for i, box := range boxes {
		for j, lens := range box {
			sum += (i + 1) * (j + 1) * lens.focal
		}
	}
	return sum
}
