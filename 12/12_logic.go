package quest12

import (
	"fmt"
	"strings"

	"github.com/Evokoo/EBC_2026_Go/utils"
)

// ========================
// PARSER
// ========================

type Coordinate [2]int

type Shape struct {
	pos       Coordinate
	rows, col int
}

func parseInput(file string) []Shape {
	data := utils.ReadFile(file)
	points := make(utils.Set[Coordinate])

	for y, line := range strings.Split(data, "\n") {
		for x, char := range line {
			if char == '#' {
				points.Add([2]int{x, y})
			}
		}
	}

	shapes := []Shape{}
	seen := make(utils.Set[Coordinate])

	for point := range points {
		if seen.Has(point) {
			continue
		}

		shape := Shape{pos: point, rows: 1, col: 1}
		queue := []Coordinate{point}

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]

			seen.Add(current)

			for i, dir := range []Coordinate{{0, 1}, {1, 0}} {
				next := Coordinate{dir[0] + current[0], dir[1] + current[1]}
				if points.Has(next) {
					queue = append(queue, next)
					switch i {
					case 0:
						shape.col++
					case 1:
						shape.rows++
					}
				}
			}
		}

		shapes = append(shapes, shape)
	}

	fmt.Println(shapes)

	return []Shape{}
}
