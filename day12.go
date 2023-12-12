package aoc2023

import (
	"regexp"
	"strings"
)

type SpringRecord struct {
	record string
	counts []int
}

var runRegexp = regexp.MustCompile("#+")

func (r SpringRecord) Matches() bool {
	runs := runRegexp.FindAllString(r.record, -1)
	if len(runs) != len(r.counts) {
		return false
	}
	for i, run := range runs {
		if len(run) != r.counts[i] {
			return false
		}
	}
	return true
}

func (r SpringRecord) PossibleMatches() int {
	before, after, found := strings.Cut(r.record, "?")
	if !found {
		if r.Matches() {
			return 1
		} else {
			return 0
		}
	}
	runs := runRegexp.FindAllString(before, -1)
	if len(runs) > len(r.counts) {
		return 0
	}
	for i, run := range runs {
		// a short run might still turn into a longer one, but a mismatched too-long run cannot be repaired
		if len(run) > r.counts[i] {
			return 0
		}
	}

	resDot := SpringRecord{before + "." + after, r.counts}.PossibleMatches()
	resHash := SpringRecord{before + "#" + after, r.counts}.PossibleMatches()

	return resDot + resHash
}

func ReadSpringRecords(path string) []SpringRecord {
	lines := ReadFile(path)

	res := []SpringRecord{}
	for _, line := range lines {
		record, countStrs, found := strings.Cut(line, " ")
		if !found {
			continue
		}
		counts := ParseNumList(countStrs)

		res = append(res, SpringRecord{record, counts})
	}
	return res
}
