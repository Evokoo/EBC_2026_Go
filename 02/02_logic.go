package quest02

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/Evokoo/EBC_2026_Go/utils"
)

// ========================
// PART I && II
// ========================

func I(alloys []string) int {
	var checksum int

	for _, line := range alloys {
		rank := make(map[int]int)

		for pair := range strings.SplitSeq(line, " ") {
			a, _ := strconv.Atoi(string(pair[0]))
			b, _ := strconv.Atoi(string(pair[2]))
			o := pair[1]

			if _, ok := rank[a]; !ok {
				rank[a] = 0
			}
			if _, ok := rank[b]; !ok {
				rank[b] = 0
			}

			if o == '>' {
				rank[a]++
				rank[b]--
			} else {
				rank[a]--
				rank[b]++
			}
		}

		fmt.Println(rank, scoreAlloy(rank))

		checksum += scoreAlloy(rank)

	}
	return checksum
}

// ========================
// PART III
// ========================

func III(alloys []string) int {
	return 0
}

func scoreAlloy(input map[int]int) int {
	keys := make([]int, 0, len(input))
	for key := range input {
		keys = append(keys, key)
	}

	slices.SortFunc(keys, func(a, b int) int {
		return input[b] - input[a]
	})

	var str strings.Builder
	for _, n := range keys {
		str.WriteString(strconv.Itoa(n))
	}

	value, _ := strconv.Atoi(str.String())
	return value
}

// ========================
// PARSER
// ========================

func parseInput(input string) []string {
	return strings.Split(utils.ReadFile(input), "\n")
}
