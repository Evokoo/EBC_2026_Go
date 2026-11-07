package quest08

import (
	"fmt"
	"strings"

	"github.com/Evokoo/EBC_2026_Go/utils"
)

// ========================
// PARSER
// ========================

type Tile struct {
	q, r, value int
}

type Tiles map[[2]int]Tile

func parseInput(file string) Tiles {
	data := utils.ReadFile(file)
	tiles := make(Tiles)
	for r, line := range strings.Split(data, "\n") {
		col := 0

		for _, char := range line {
			if char != '.' {
				n := int(char - '0')
				q := col - (r / 2)
				tiles[[2]int{q, r}] = Tile{q, r, n}
				col++
			}
		}
	}
	fmt.Println(tiles)
	return tiles
}
