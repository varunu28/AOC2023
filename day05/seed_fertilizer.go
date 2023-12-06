package day05

import (
	"strconv"
	"strings"
)

type seedPair struct {
	start int
	end   int
}

type rangePair struct {
	source      int
	destination int
	rangeLength int
}

// FindLowestLocationNumberForSeeds finds the lowest location number for seeds after processing
// the mapping from almanac
func FindLowestLocationNumberForSeeds(input []string) int {
	// parse seeds
	seeds := make(map[int]int)
	line := input[0]
	splits := strings.Split(line, ":")
	seedSplits := strings.Fields(splits[1])
	for _, seedSplit := range seedSplits {
		seed, _ := strconv.Atoi(seedSplit)
		seeds[seed] = seed
	}
	for i := 1; i < len(input); i++ {
		line := input[i]
		if line == "" {
			continue
		}
		idx := i + 1
		tmp := make(map[int]int)
		for idx < len(input) && input[idx] != "" {
			splits := strings.Fields(input[idx])
			destination, _ := strconv.Atoi(splits[0])
			source, _ := strconv.Atoi(splits[1])
			mapRange, _ := strconv.Atoi(splits[2])
			for k := range seeds {
				if k >= source && k < source+mapRange {
					tmp[k] = destination + (k - source)
				}
			}
			idx++
		}
		i = idx
		// reset seeds map
		for k := range seeds {
			if _, ok := tmp[k]; !ok {
				tmp[k] = k
			}
		}
		seeds = make(map[int]int)
		for _, v := range tmp {
			seeds[v] = v
		}
	}

	lowest := -1
	for _, v := range seeds {
		if lowest == -1 || v < lowest {
			lowest = v
		}
	}
	return lowest
}

// FindLowestLocationNumberForSeedRange finds the lowest location number for seed ranges after processing
// the mapping from almanac
func FindLowestLocationNumberForSeedRange(input []string) int {
	// parse seeds
	splits := strings.Split(input[0], ":")
	seedSplits := strings.Fields(splits[1])
	seeds := make([]seedPair, 0)
	for i := 0; i < len(seedSplits); i += 2 {
		start, _ := strconv.Atoi(seedSplits[i])
		end, _ := strconv.Atoi(seedSplits[i+1])
		seeds = append(seeds, seedPair{start, start + end})
	}
	for i := 2; i < len(input); i++ {
		idx := i + 1
		// parse ranges in a mapping
		ranges := make([]rangePair, 0)
		for idx < len(input) && input[idx] != "" {
			splits := strings.Fields(input[idx])
			destination, _ := strconv.Atoi(splits[0])
			source, _ := strconv.Atoi(splits[1])
			mapRange, _ := strconv.Atoi(splits[2])
			ranges = append(ranges, rangePair{source, destination, mapRange})
			idx++
		}
		tmp := make([]seedPair, 0)
		for len(seeds) > 0 {
			// pop the last seed
			s := seeds[len(seeds)-1]
			seeds = seeds[:len(seeds)-1]
			overlap := false
			for _, r := range ranges {
				overlapStart := max(s.start, r.source)
				overlapEnd := min(s.end, r.source+r.rangeLength)
				if overlapStart < overlapEnd {
					tmp = append(tmp, seedPair{overlapStart - r.source + r.destination, overlapEnd - r.source + r.destination})
					if overlapStart > s.start {
						seeds = append(seeds, seedPair{s.start, overlapStart})
					}
					if s.end > overlapEnd {
						seeds = append(seeds, seedPair{overlapEnd, s.end})
					}
					overlap = true
					break
				}
			}
			if !overlap {
				tmp = append(tmp, s)
			}
		}
		seeds = tmp
		i = idx
	}
	lowest := seeds[0].start
	for _, s := range seeds {
		if s.start < lowest {
			lowest = s.start
		}
	}
	return lowest
}
