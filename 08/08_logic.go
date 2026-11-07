package quest08

import (
	"fmt"
	"strings"

	"github.com/Evokoo/EBC_2026_Go/utils"
)

// ========================
// PART I
// ========================
func I(tiles Tiles) int {
	edge := tiles.LocateEdge()
	total := 0

	for _, tile := range edge {
		count := 0
		height := tile.value

		for {
			if next, valid := tiles.Move(tile, DIRECTIONS["DR"]); valid {
				if next.value < height {
					count++
					height = max(height-1, 0)
				} else {
					height = max(height, next.value)
				}

				tile = next
			} else {
				break
			}
		}

		total += count
	}

	fmt.Println(total)
	return 0
}

// ========================
// PARSER
// ========================

var DIRECTIONS = map[string][2]int{
	"R":  {1, 0}, //{q,r}
	"DR": {1, 1},
	"DL": {0, 1},
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
	tiles := make(Tiles)
	for r, line := range strings.Split(data, "\n") {
		for i, char := range line {
			if char != '.' {
				n := int(char - '0')
				q := (i - (4 - r)) / 2
				tiles[[2]int{q, r}] = Tile{q, r, n}
			}
		}
	}

	tiles.LocateEdge()
	return tiles
}
