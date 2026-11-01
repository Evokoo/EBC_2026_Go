package quest05

import (
	"strings"

	"github.com/Evokoo/EBC_2026_Go/utils"
)

// ========================
// PART I
// ========================

func I(input Cells) int {
	x := input.cols / 2
	y := input.rows / 2

	var queue [][2]int
	var seen map[string]struct{}

	count := 0
	queue = append(queue, [2]int{x, y})

	for len(queue) > 0 {

	}

	return 0
}

// ========================
// PARSER
// ========================

type Cells struct {
	grid       [][]rune
	rows, cols int
}

func parseInput(file string) Cells {
	data := utils.ReadFile(file)

	var result Cells
	for line := range strings.SplitSeq(data, "\n") {
		row := []rune(line)

		if result.cols == 0 {
			result.cols = len(row)
		}
		result.rows++
		result.grid = append(result.grid, row)
	}
	return result
}

// ========================
// Queue
// ========================
type Queue [][2]int

func newQueue() Queue {
	return Queue{}
}

func (q *Queue) Push(pair [2]int) {
	(*q) = append((*q), pair)
}

func (q *Queue) Pop() [2]int {
	item := (*q)[0]
	*q = (*q)[1:]
	return item
}

func (q Queue) isEmpty() bool {
	return len(q) > 0
}

// ========================
// Set
// ========================
type Set map[string]struct{}

func (s Set) Add(item string) {
	s[item] = struct{}{}
}
func (s Set) Has(item string) bool {
	_, found := s[item]
	return found
}
func (s Set) Remove(item string) {
	delete(s, item)
}
