package aoc2023

import (
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

	res := possibleMatches(chunksDot, r.counts)
	return res
}

func minLengthForTail(counts []int) int {
	return Sum(counts) + len(counts) // we have a dot afer each part plus the sentinel dot
}

func matchesMinPattern(str string, counts []int) bool {
	i := 0
	for _, count := range counts {
		maxI := i + count
		for ; i < maxI; i++ {
			if str[i] == '.' {
				return false
			}
		}
		if str[i] == '#' {
			return false
		}
		i++
	}
	return true
}

func possibleMatches(chunks []string, counts []int) int {
	if len(chunks) == 0 {
		if len(counts) == 0 {
			return 1
		} else {
			return 0
		}
	} else if len(chunks) == 1 {
		// final chunk - we are looking to get rid of all the counts, if it doesn't fit, it's not a match
		degreesOfFreedom := len(chunks[0]) - minLengthForTail(counts)
		if degreesOfFreedom < 0 {
			return 0
		} else if degreesOfFreedom == 0 {
			if matchesMinPattern(chunks[0], counts) {
				return 1
			} else {
				return 0
			}
		}
	}

	results := prefixMatches(0, chunks[0], counts)

	sum := 0
	for _, res := range results {
		sum += possibleMatches(chunks[1:], counts[len(counts)-res.suffixLen:]) * res.matches
	}
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

var cache = map[string][]MatchResult{}

func makeCacheKey(hashRun int, str string, counts []int) string {
	builder := strings.Builder{}
	builder.Grow(32)
	builder.WriteRune(rune(hashRun))
	builder.WriteString(str)
	builder.WriteByte(':')
	for _, count := range counts {
		builder.WriteRune(rune(count))
	}
	return builder.String()
}

func prefixMatches(hashRun int, str string, counts []int) []MatchResult {
	key := makeCacheKey(hashRun, str, counts)
	res, found := cache[key]
	if found {
		return res
	}

	res = prefixMatchesInt(hashRun, str, counts)
	cache[key] = res

	return res
}

func prefixMatchesInt(hashRun int, str string, counts []int) []MatchResult {
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
				return prefixMatches(hashRun+1, str[1:], counts)
			} else {
				// forced to terminate the run using dot, no need to explore another #
				return prefixMatches(0, str[1:], counts[1:])
			}
		}
		dotRes := prefixMatches(hashRun, "."+str[1:], counts)
		hashRes := prefixMatches(hashRun+1, str[1:], counts)

		return merge(dotRes, hashRes, len(counts))
	case '.':
		if hashRun > 0 {
			if len(counts) == 0 || counts[0] != hashRun {
				return []MatchResult{}
			}
			return prefixMatches(0, str[1:], counts[1:])
		}
		return prefixMatches(0, str[1:], counts)
	case '#':
		return prefixMatches(hashRun+1, str[1:], counts)
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
