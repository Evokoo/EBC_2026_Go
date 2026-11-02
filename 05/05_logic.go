package quest05

import (
	"strings"

	"github.com/Evokoo/EBC_2026_Go/utils"
)

// ========================
// LANDMASS
// ========================
type Landmass struct {
	layout     [][]rune
	rows, cols int
}
type Coordinate struct {
	x, y int
}

var DIRECTIONS = []Coordinate{
	{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1},
}

func (l Landmass) GetCoordinate(coordinate Coordinate) rune {
	return l.layout[coordinate.y][coordinate.x]
}
func (l *Landmass) UpdateCoordinate(coordinate Coordinate) {
	(*l).layout[coordinate.y][coordinate.x] = '#'
}
func (l *Landmass) ValidateCoordinate(coordinate Coordinate) (bool, []Coordinate) {
	score := 0

	var neighbours []Coordinate
	for _, dir := range DIRECTIONS {
		coord := Coordinate{coordinate.x + dir.x, coordinate.y + dir.y}

		if coord.x >= 0 && coord.x < l.cols && coord.y >= 0 && coord.y < l.rows {
			tile := l.GetCoordinate(coord)

			switch tile {
			case '0':
				neighbours = append(neighbours, coord)
			case '1':
				neighbours = append(neighbours, coord)
				score++
			case '#':
				score++
			}
		}
	}

	return score >= 5, neighbours
}
func (l *Landmass) InitSearch() Queue {
	start := Coordinate{l.cols / 2, l.rows / 2}
	(*l).UpdateCoordinate(start)

	queue := newQueue()

	for _, dir := range DIRECTIONS {
		coord := Coordinate{start.x + dir.x, start.y + dir.y}
		if coord.x >= 0 && coord.x < l.cols && coord.y >= 0 && coord.y < l.rows {
			queue.Push(coord)
		}

	}

	return queue
}

// ========================
// PART I
// ========================

func I(lm Landmass) int {
	queue := lm.InitSearch()
	count := 1

	for !queue.isEmpty() {
		current := queue.Pop()

		if lm.GetCoordinate(current) == '#' {
			continue
		}

		if isLand, neighbours := lm.ValidateCoordinate(current); isLand {
			count++
			lm.UpdateCoordinate(current)

			for _, neighbour := range neighbours {
				queue.Push(neighbour)
			}
		}
	}

	return count
}

// ========================
// PARSER
// ========================

func parseInput(file string) Landmass {
	data := utils.ReadFile(file)

	var landmass Landmass
	for line := range strings.SplitSeq(data, "\n") {
		row := []rune(line)

		if landmass.cols == 0 {
			landmass.cols = len(row)
		}
		landmass.rows++
		landmass.layout = append(landmass.layout, row)
	}
	return landmass
}

// ========================
// Queue
// ========================
type Queue []Coordinate

func newQueue() Queue {
	return Queue{}
}
func (q *Queue) Push(coordinate Coordinate) {
	(*q) = append((*q), coordinate)
}
func (q *Queue) Pop() Coordinate {
	coordinate := (*q)[0]
	*q = (*q)[1:]
	return coordinate
}
func (q Queue) isEmpty() bool {
	return len(q) == 0
}

// ========================
// Set
// ========================
type Set map[Coordinate]struct{}

func (s Set) Add(coord Coordinate) {
	s[coord] = struct{}{}
}
func (s Set) Has(coord Coordinate) bool {
	_, found := s[coord]
	return found
}
func (s Set) Remove(coord Coordinate) {
	delete(s, coord)
}
