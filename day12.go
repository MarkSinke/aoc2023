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

var chunkRegexp = regexp.MustCompile("[#\\?]+")

func (r SpringRecord) PossibleMatches() int {
	chunks := chunkRegexp.FindAllString(r.record, -1)
	var chunksDot = []string{}
	for _, ch := range chunks {
		chunksDot = append(chunksDot, ch+".")
	}

	calls := 0
	res := possibleMatches(&calls, chunksDot, r.counts)
	fmt.Printf("(calls: %v)", calls)
	return res
}

func minLengthForTail(counts []int) int {
	return Sum(counts) + len(counts) // we have a dot afer each part plus the sentinel dot
}

func possibleMatches(calls *int, chunks []string, counts []int) int {
	// fmt.Println("pm2", chunks, counts)
	if len(chunks) == 0 {
		if len(counts) == 0 {
			return 1
		} else {
			return 0
		}
	} else if len(chunks) == 1 {
		// final chunk - we are looking to get rid of all the counts, if it doesn't fit, it's not a match
		if len(chunks[0]) < minLengthForTail(counts) {
			return 0
		}
	}

	results := prefixMatches(calls, 0, chunks[0], counts)
	// fmt.Println(" prefixes", results)

	sum := 0
	for _, res := range results {
		sum += possibleMatches(calls, chunks[1:], counts[len(counts)-res.suffixLen:]) * res.matches
	}
	// fmt.Println("pm2 ret", sum)
	return sum
}

type MatchResult struct {
	suffixLen int
	matches   int
}

func merge(res0 []MatchResult, res1 []MatchResult, max int) []MatchResult {
	var countList = make([]int, max+1)
	for _, r := range res0 {
		countList[r.suffixLen] += r.matches
	}
	for _, r := range res1 {
		countList[r.suffixLen] += r.matches
	}
	res := make([]MatchResult, 0, max)
	for len, r := range countList {
		if r != 0 {
			res = append(res, MatchResult{len, r})
		}
	}
	return res
}

func prefixMatches(calls *int, hashRun int, str string, counts []int) []MatchResult {
	*calls++
	if hashRun > 0 {
		if len(counts) == 0 || hashRun > counts[0] {
			return []MatchResult{}
		}
	}
	if len(str) == 0 {
		if hashRun != 0 {
			panic("hashRun != 0")
		}
		return []MatchResult{{len(counts), 1}}
	}

	switch str[0] {
	case '?':
		if hashRun > 0 {
			if hashRun < counts[0] {
				// forced to continue the run using another # - no need to explore the dot
				return prefixMatches(calls, hashRun+1, str[1:], counts)
			} else {
				// forced to terminate the run using dot, no need to explore another #
				return prefixMatches(calls, 0, str[1:], counts[1:])
			}
		}
		dotRes := prefixMatches(calls, hashRun, "."+str[1:], counts)
		hashRes := prefixMatches(calls, hashRun+1, str[1:], counts)

		return merge(dotRes, hashRes, len(counts))
	case '.':
		if hashRun > 0 {
			if len(counts) == 0 || counts[0] != hashRun {
				return []MatchResult{}
			}
			return prefixMatches(calls, 0, str[1:], counts[1:])
		}
		return prefixMatches(calls, 0, str[1:], counts)
	case '#':
		return prefixMatches(calls, hashRun+1, str[1:], counts)
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
