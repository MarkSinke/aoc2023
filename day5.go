package aoc2023

import (
	"strings"
)

func CutLabel(str string) string {
	parts := strings.Split(str, ":")
	return parts[1]
}

type SeedMapEntry struct {
	dest   int
	src    int
	length int
}

type Range struct {
	start int
	end   int
}

func (r Range) IsEmpty() bool {
	return r.start >= r.end
}

func (e SeedMapEntry) MapRange(r Range) (Range, []Range) {
	rLeft := Range{r.start, min(r.end, e.src)}
	rOverlap := Range{max(e.src, r.start), min(r.end, e.src+e.length)}
	rRight := Range{max(e.src+e.length, r.start), r.end}

	var mapped Range
	var toMapNext []Range

	if !rLeft.IsEmpty() {
		toMapNext = append(toMapNext, rLeft)
	}
	if !rOverlap.IsEmpty() {
		mapped = Range{rOverlap.start + e.dest - e.src, rOverlap.end + e.dest - e.src}
	}
	if !rRight.IsEmpty() {
		toMapNext = append(toMapNext, rRight)
	}

	return mapped, toMapNext
}

type SeedMap struct {
	entries []SeedMapEntry
}

func (m SeedMap) Map(from int) int {
	for _, entry := range m.entries {
		if from >= entry.src && from < entry.src+entry.length {
			return from - entry.src + entry.dest
		}
	}

	// evervything not explicitly mapped is a 1:1
	return from
}

func (m SeedMap) MapRanges(toMapIn []Range) []Range {
	toMap := toMapIn
	var mapped []Range

	// first map the range through the entriesm amending the to-map range as we go
	for _, entry := range m.entries {
		var toMapNext []Range
		for _, r := range toMap {
			mappedEntry, toMapNextEntry := entry.MapRange(r)
			if !mappedEntry.IsEmpty() {
				mapped = append(mapped, mappedEntry)
			}
			toMapNext = append(toMapNext, toMapNextEntry...)
		}
		toMap = toMapNext
	}
	// evervything not explicitly mapped is a 1:1
	mapped = append(mapped, toMap...)

	return mapped
}

func (m SeedMap) IsEmpty() bool {
	return len(m.entries) == 0
}

func (m *SeedMap) AddEntry(dest int, src int, length int) {
	m.entries = append(m.entries, SeedMapEntry{dest, src, length})
}

func ReadSeedMaps(path string) ([]int, []SeedMap) {
	lines := ReadFile(path)

	var seeds []int
	var maps []SeedMap
	var currentMap SeedMap

	for _, line := range lines {
		switch {
		case strings.Contains(line, "seeds:"):
			seeds = ParseNumList(CutLabel(line))

		case strings.Contains(line, "map:"):
			currentMap = SeedMap{}

		case line == "":
			if !currentMap.IsEmpty() {
				maps = append(maps, currentMap)
				currentMap = SeedMap{}
			}

		case true:
			nums := ParseNumList(line)
			currentMap.AddEntry(nums[0], nums[1], nums[2])
		}
	}
	return seeds, maps
}

func MapSeedThroughMaps(seed int, maps []SeedMap) int {
	cur := seed
	for _, m := range maps {
		cur = m.Map(cur)
	}
	return cur
}

func MapSeedsThroughMaps(seeds []int, maps []SeedMap) []int {
	var res = make([]int, len(seeds))
	for i, seed := range seeds {
		res[i] = MapSeedThroughMaps(seed, maps)
	}
	return res
}

func MapRangesThroughMaps(rs []Range, maps []SeedMap) []Range {
	cur := rs
	for _, m := range maps {
		cur = m.MapRanges(cur)
	}
	return cur
}
