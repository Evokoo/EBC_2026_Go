package quest12

import (
	"strings"

	"github.com/Evokoo/EBC_2026_Go/utils"
)

// // ========================
// // I
// // ========================
// func I(shapes []Shape) int {
// 	spells := 0

// 	for _, shape := range shapes {
// 		spells += min(shape.cols, shape.rows)
// 	}

// 	return spells
// }

// ========================
// I & II
// ========================
func I(shapes []Shape) int {
	spells := 0

	for _, shape := range shapes {
		spells += min(shape.cols, shape.rows, shape.shortX+shape.shortY)
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
	pos            Coordinate
	points         utils.Set[Coordinate]
	rows, cols     int
	shortX, shortY int
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
			shape := Shape{pos: start, points: make(utils.Set[Coordinate])}
			queue := []Coordinate{start}

			minX, maxX := start.x, start.x
			minY, maxY := start.y, start.y

			for len(queue) > 0 {
				current := queue[0]
				queue = queue[1:]

				if visited.Has(current) {
					continue
				}
				visited.Add(current)
				shape.points.Add(current)

				minX = min(minX, current.x)
				maxX = max(maxX, current.x)
				minY = min(minY, current.y)
				maxY = max(maxY, current.y)

				for _, dir := range DIRECTIONS {
					nx, ny := dir.x+current.x, dir.y+current.y
					next := Coordinate{nx, ny}

					if nx < 0 || nx >= cols || ny < 0 || ny >= rows || grid[ny][nx] != '#' || visited.Has(next) {
						continue
					}

					queue = append(queue, next)
				}
			}

			shape.pos = Coordinate{x: minX, y: minY}
			shape.cols = maxX - minX + 1
			shape.rows = maxY - minY + 1
			shapes = append(shapes, shape)
		}
	}

	for i, shape := range shapes {
		posA := shape.pos
		posB := Coordinate{x: shape.pos.x, y: shape.pos.y + shape.rows - 1}
		a, b := 0, 0

		for range shape.cols {
			if grid[posA.y][posA.x] == '#' && shape.points.Has(posA) {
				a++
			}
			if grid[posB.y][posB.x] == '#' && shape.points.Has(posB) {
				b++
			}
			posA.x++
			posB.x++
		}

		shape.shortX = min(a, b)

		posA = shape.pos
		posB = Coordinate{x: shape.pos.x + shape.cols - 1, y: shape.pos.y}
		a, b = 0, 0

		for range shape.rows {
			if grid[posA.y][posA.x] == '#' && shape.points.Has(posA) {
				a++
			}
			if grid[posB.y][posB.x] == '#' && shape.points.Has(posB) {
				b++
			}
			posA.y++
			posB.y++
		}

		shape.shortY = min(a, b)
		shapes[i] = shape
	}

	return shapes
}
