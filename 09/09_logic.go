package quest09

import (
	"strconv"
	"strings"

	"github.com/Evokoo/EBC_2026_Go/utils"
)

// ========================
// I
// ========================

func I(seeds []Seed) int {
	result := 0

	for _, seed := range seeds {
		elements, seconds := 0, 0
		segments := []int{0}

		for elements < seed.limit {
			currentSegments := len(segments)
			for i := range currentSegments {
				segments[i]++
				elements++

				if segments[i] == seed.growth {
					segments[i] = seed.growth / 2
					segments = append(segments, seed.growth/2)
					segments = append(segments, 1)
					elements++
				}
			}
			seconds++
		}

		result += seconds
	}

	return result
}

// ========================
// PARSER
// ========================

type Seed struct {
	id, growth, limit int
}

func parseInput(file string) []Seed {
	data := utils.ReadFile(file)

	var seeds []Seed
	for line := range strings.SplitSeq(data, "\n") {
		seed := Seed{}

		for i, value := range utils.QuickMatch(line, `\d+`) {
			n, _ := strconv.Atoi(value)
			switch i {
			case 0:
				seed.id = n
			case 1:
				seed.growth = n
			case 2:
				seed.limit = n
			}
		}
		seeds = append(seeds, seed)
	}
	return seeds
}
