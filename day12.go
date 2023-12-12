package aoc2023

import (
	"fmt"
	"regexp"
	"strings"
)

type SpringRecord struct {
	record string
	counts []int
}

var runRegexp = regexp.MustCompile("[#\\?]+")

func (r SpringRecord) PossibleMatches() int {
	calls := 0
	res := possibleMatches2(&calls, 0, r.record+".", r.counts)
	fmt.Println(r, "calls", calls)
	return res
}

func possibleMatches2(calls *int, hashRun int, str string, counts []int) int {
	*calls++
	if len(str) == 0 {
		if len(counts) == 0 {
			return 1
		} else {
			return 0
		}
	}
	if len(counts) > (len(str)+1)/2 {
		// no way we can still match
		return 0
	}

	switch str[0] {
	case '?':
		return possibleMatches2(calls, hashRun, "."+str[1:], counts) + possibleMatches2(calls, hashRun+1, str[1:], counts)
	case '.':
		if hashRun > 0 {
			if len(counts) == 0 || counts[0] != hashRun {
				return 0
			}
			return possibleMatches2(calls, 0, str[1:], counts[1:])
		}
		return possibleMatches2(calls, 0, str[1:], counts)
	case '#':
		return possibleMatches2(calls, hashRun+1, str[1:], counts)
	default:
		panic("bad character " + str[0:0])
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
