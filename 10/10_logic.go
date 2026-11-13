package quest10

import (
	"strings"

	"github.com/Evokoo/EBC_2026_Go/utils"
)

// ========================
// I
// ========================

func I(field Field, rounds int) int {

	for range rounds {
		field.Update()
	}

	return field.collisions

}

// ========================
// PARSER
// ========================

var DIRECTION = map[int]Coordinate{
	0:   {0, -1},
	90:  {1, 0},
	180: {0, 1},
	270: {-1, 0},
}

type Coordinate struct {
	x, y int
}

type Ghost struct {
	id, bearing int
	pos         Coordinate
}

type Field struct {
	ghosts     []*Ghost
	walls      utils.Set[Coordinate]
	collisions int
}

func (f *Field) Update() {
	//Move Ghosts and handle wall collisions
	for _, g := range f.ghosts {
		dir := DIRECTION[g.bearing]
		next := Coordinate{g.pos.x + dir.x, g.pos.y + dir.y}

		if f.walls.Has(next) {
			f.collisions++
			g.bearing = (g.bearing + 90) % 360
		} else {
			g.pos.x += dir.x
			g.pos.y += dir.y
		}
	}

	//Check Ghost collisions
	occupied := make(map[Coordinate][]*Ghost)
	for _, g := range f.ghosts {
		occupied[g.pos] = append(occupied[g.pos], g)
	}

	for _, ghosts := range occupied {
		if len(ghosts) > 1 {
			for _, g := range ghosts {
				f.collisions++
				g.bearing = (g.bearing + 180) % 360
			}
		}
	}
}

func parseInput(file string) Field {
	data := utils.ReadFile(file)

	var ghosts []*Ghost
	walls := make(utils.Set[Coordinate])

	for y, row := range strings.Split(data, "\n") {
		for x, char := range row {
			switch char {
			case '#':
				walls.Add(Coordinate{x, y})
			case '<', '>', 'v', '^':
				id := len(ghosts)
				bearing := map[rune]int{'<': 270, '>': 90, '^': 0, 'v': 180}[char]
				ghosts = append(ghosts, &Ghost{id, bearing, Coordinate{x, y}})
			}
		}
	}
	return Field{ghosts, walls, 0}
}
