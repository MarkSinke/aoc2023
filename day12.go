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
	// fmt.Println("matching", r)
	before, after, found := strings.Cut(r.record, "?")
	if !found {
		// fmt.Println("  !found", r.Matches())
		if r.Matches() {
			return 1
		} else {
			return 0
		}
	}
	options := strings.Count(after, "?") + strings.Count(after, "#")
	runs := runRegexp.FindAllString(before, -1)
	// fmt.Println("  runs", runs)
	if len(runs) == 0 {
		return SpringRecord{after, r.counts}.PossibleMatches() + SpringRecord{"#" + after, r.counts}.PossibleMatches()
	}
	if len(runs) > len(r.counts) {
		return 0
	}
	// very rough upper bound to avoid lots of recursion for long count lists
	if len(r.counts)-len(runs)-1 > options {
		// fmt.Println("shortcut", r, before, after, runs, options)
		return 0
	}

	for i := 0; i < len(runs)-1; i++ {
		if len(runs[i]) != r.counts[i] {
			return 0
		}
	}
	lastIndex := len(runs) - 1
	lastRun := runs[lastIndex]
	// fmt.Println("  lI, lR", lastIndex, lastRun)
	return r.possibleMatchesDot(lastIndex, lastRun, after) + r.possibleMatchesHash(before, lastIndex, lastRun, after)
}

func (r SpringRecord) possibleMatchesDot(lastIndex int, lastRun string, after string) int {
	if len(lastRun) != r.counts[lastIndex] {
		return 0
	}

	rec := SpringRecord{after, r.counts[lastIndex+1:]}
	// fmt.Println("recurse .", rec)
	return rec.PossibleMatches()
}

func (r SpringRecord) possibleMatchesHash(before string, lastIndex int, lastRun string, after string) int {
	if before[len(before)-1] == '#' {
		if len(lastRun) > r.counts[lastIndex] {
			// a short run might still turn into a longer one, but a mismatched too-long run cannot be repaired
			return 0
		}

		rec := SpringRecord{lastRun + "#" + after, r.counts[lastIndex:]}
		// fmt.Println("recurse #1", rec)
		return rec.PossibleMatches()
	} else {
		if len(lastRun) != r.counts[lastIndex] {
			return 0
		}

		rec := SpringRecord{"#" + after, r.counts[lastIndex+1:]}
		// fmt.Println("recurse #2", rec)
		return rec.PossibleMatches()
	}
}

func (r SpringRecord) Unfold() SpringRecord {
	newRecord := r.record + "?" + r.record + "?" + r.record + "?" + r.record + "?" + r.record
	newCounts := append(r.counts, r.counts...)
	newCounts = append(newCounts, r.counts...)
	newCounts = append(newCounts, r.counts...)
	newCounts = append(newCounts, r.counts...)
	return SpringRecord{newRecord, newCounts}
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
