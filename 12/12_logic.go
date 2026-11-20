package quest12

import (
	"strings"

	"github.com/Evokoo/EBC_2026_Go/utils"
)

// ========================
// I
// ========================
func I(shapes []Shape) int {
	spells := 0

	for _, shape := range shapes {
		spells += min(shape.cols, shape.rows)
	}

	return spells
}

// ========================
// II
// ========================
func II(shapes []Shape) int {
	spells := 0

	for _, shape := range shapes {
		if shape.cols != shape.rows {
			dual := max(shape.cols, shape.rows) - min(shape.cols, shape.rows)

			if dual < shape.cols && dual < shape.rows {
				spells += dual
			} else {
				spells += min(shape.cols, shape.rows)
			}
		}
	}

	return spells
}

// ========================
// PARSER
// ========================

var DIRECTIONS = []Coordinate{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

type Coordinate struct {
	x, y int
}

type Shape struct {
	pos        Coordinate
	rows, cols int
}

func parseInput(file string) []Shape {
	data := utils.ReadFile(file)
	grid := strings.Split(data, "\n")
	rows, cols := len(grid), len(grid[0])

	shapes := []Shape{}
	visited := make(utils.Set[Coordinate])

	for y, row := range grid {
		for x, char := range row {
			if char != '#' || visited.Has(Coordinate{x, y}) {
				continue
			}

			start := Coordinate{x, y}
			shape := Shape{pos: start, rows: 0, cols: 0}
			queue := []Coordinate{start}

			for len(queue) > 0 {
				current := queue[0]
				queue = queue[1:]

				if visited.Has(current) {
					continue
				} else {
					visited.Add(current)
				}

				for _, dir := range DIRECTIONS {
					nx, ny := dir.x+current.x, dir.y+current.y

					if nx < 0 || nx >= cols || ny < 0 || ny >= rows || grid[ny][nx] != '#' {
						continue
					}

					queue = append(queue, Coordinate{nx, ny})
				}
				shape.cols = max(shape.cols, current.x-start.x+1)
				shape.rows = max(shape.rows, current.y-start.y+1)
			}
			shapes = append(shapes, shape)
		}
	}

	return shapes
}
