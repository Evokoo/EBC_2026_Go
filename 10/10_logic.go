package quest10

import (
	"fmt"
	"math"
	"strings"

	"github.com/Evokoo/EBC_2026_Go/utils"
)

// ========================
// I
// ========================

func I(field Field, rounds int) int {

	for field.round < rounds {
		field.Update()
	}

	return field.CountCollisions()
}

// ========================
// II
// ========================

func II(field Field, limit int) int {
	field.SetCollisionLimit(limit)

	for field.activeGhosts > 0 {
		field.Update()
	}

	fmt.Println()
	return field.round
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
	id, bearing, collisions int
	pos                     Coordinate
	isActive                bool
}

type Field struct {
	ghosts                              []*Ghost
	walls                               utils.Set[Coordinate]
	round, collisionLimit, activeGhosts int
}

func (f *Field) SetCollisionLimit(limit int) {
	f.collisionLimit = limit
}

func (f *Field) Update() {
	//Move Ghosts and handle wall collisions
	for _, g := range f.ghosts {
		if !g.isActive {
			continue
		}

		dir := DIRECTION[g.bearing]
		next := Coordinate{g.pos.x + dir.x, g.pos.y + dir.y}

		if f.walls.Has(next) {
			g.collisions++
			g.bearing = (g.bearing + 90) % 360

			if g.collisions == f.collisionLimit {
				g.isActive = false
				f.activeGhosts--
			}

		} else {
			g.pos.x += dir.x
			g.pos.y += dir.y
		}
	}

	//Check Ghost collisions
	occupied := make(map[Coordinate][]*Ghost)
	for _, g := range f.ghosts {
		if !g.isActive {
			continue
		}
		occupied[g.pos] = append(occupied[g.pos], g)
	}

	for _, ghosts := range occupied {
		if len(ghosts) > 1 {
			for _, g := range ghosts {
				g.collisions++
				g.bearing = (g.bearing + 180) % 360

				if g.collisions == f.collisionLimit {
					g.isActive = false
					f.activeGhosts--
				}
			}
		}
	}

	f.round++
}

func (f Field) CountCollisions() int {
	count := 0
	for _, g := range f.ghosts {
		count += g.collisions
	}
	return count
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
				ghosts = append(ghosts, &Ghost{id, bearing, 0, Coordinate{x, y}, true})
			}
		}
	}
	return Field{
		ghosts:         ghosts,
		walls:          walls,
		round:          0,
		activeGhosts:   len(ghosts),
		collisionLimit: math.MaxInt,
	}
}
