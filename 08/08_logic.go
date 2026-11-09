package quest08

import (
	"strings"

	"github.com/Evokoo/EBC_2026_Go/utils"
)

// ========================
// PART I
// ========================
func I(tiles Tiles) int {
	edge := tiles.LocateEdges()["DR"]
	total := 0

	for _, tile := range edge {
		powered := 1
		shadow := tile.value

		for {
			if next, valid := tiles.Move(tile, DIRECTIONS["DR"]); valid {
				if next.value >= shadow {
					powered++
				}
				shadow = max(shadow-1, next.value)
				tile = next
			} else {
				break
			}
		}

		total += powered
	}
	return total
}

// ========================
// PART II
// ========================

func II(tiles Tiles) int {
	total := 0

	for key, edge := range tiles.LocateEdges() {
		for _, tile := range edge {
			powered := 1
			shadow := tile.value

			for {
				if next, valid := tiles.Move(tile, DIRECTIONS[key]); valid {
					if next.value >= shadow {
						powered++
					}
					shadow = max(shadow-1, next.value)
					tile = next
				} else {
					break
				}
			}

			total += powered
		}
	}
	return total
}

// ========================
// PARSER
// ========================

var DIRECTIONS = map[string][2]int{
	"R":  {1, 0},
	"DR": {1, 1},
	"UR": {0, -1},
	"L":  {-1, 0},
	"DL": {0, 1},
	"UL": {-1, -1},
}

type Tile struct {
	q, r, value int
}

type Tiles map[[2]int]Tile

func (t Tiles) Move(current Tile, direction [2]int) (Tile, bool) {
	next := [2]int{current.q + direction[0], current.r + direction[1]}
	tile, valid := t[next]
	return tile, valid
}

type EdgeParts struct {
	key   string
	sides [2]int
}
type EdgeMap map[string][]Tile

func (t Tiles) LocateEdges() EdgeMap {
	tile := t[[2]int{0, 0}]
	sides := make(map[int][]Tile)
	for i, dir := range []string{"R", "DR", "DL", "L", "UL", "UR"} {
		side := []Tile{tile}

		for {
			if next, ok := t.Move(tile, DIRECTIONS[dir]); ok {
				side = append(side, next)
				tile = next
			} else {
				break
			}
		}
		sides[i] = side
	}

	edges := make(EdgeMap)
	parts := []EdgeParts{
		{"DR", [2]int{5, 0}},
		{"DL", [2]int{0, 1}},
		{"L", [2]int{1, 2}},
		{"UL", [2]int{2, 3}},
		{"UR", [2]int{3, 4}},
		{"R", [2]int{4, 5}},
	}

	for _, part := range parts {
		var tiles []Tile

		for i, n := range part.sides {
			side := sides[n]

			if i == 1 {
				tiles = append(tiles, side[1:]...)
			} else {
				tiles = append(tiles, side...)
			}
		}
		edges[part.key] = tiles
	}

	return edges
}

func (t Tiles) LocateEdge() []Tile {
	start := t[[2]int{0, 0}]
	edge := []Tile{start}

	current := start
	for {
		if next, ok := t.Move(current, DIRECTIONS["R"]); ok {
			edge = append(edge, next)
			current = next
		} else {
			break
		}
	}

	current = start
	for {
		if next, ok := t.Move(current, DIRECTIONS["DL"]); ok {
			edge = append(edge, next)
			current = next
		} else {
			break
		}
	}

	return edge
}

func parseInput(file string) Tiles {
	data := utils.ReadFile(file)
	offset := len(utils.QuickMatch(data, `^\.+`)[0])

	tiles := make(Tiles)
	for r, line := range strings.Split(data, "\n") {
		for i, char := range line {
			if char != '.' {
				n := int(char - '0')
				q := (i - (offset - r)) / 2
				tiles[[2]int{q, r}] = Tile{q, r, n}
			}
		}
	}
	return tiles
}
