package quest11

import (
	"strconv"
	"strings"

	"github.com/Evokoo/EBC_2026_Go/utils"
)

// ========================
// PARSER
// ========================

type Portal struct {
	id          int
	connections []*Portal
}
type Ring map[int]*Portal

func parseInput(file string) Ring {
	data := utils.ReadFile(file)

	var count int
	var links [][2]int

	for i, line := range strings.Split(data, "\n") {
		if len(line) == 0 {
			continue
		}

		values := utils.QuickMatch(line, `\d+`)

		switch i {
		case 0:
			n, _ := strconv.Atoi(values[0])
			count = n
		default:
			a, _ := strconv.Atoi(values[0])
			b, _ := strconv.Atoi(values[1])
			links = append(links, [2]int{a, b})
		}
	}

	ring := make(Ring)
	for i := range count {
		ring[i+1] = &Portal{id: i + 1, connections: []*Portal{}}
	}

	for i := range count {
		curr := i + 1
		next := (curr % count) + 1
		prev := (curr-2+count)%count + 1
		ring[curr].connections = append(ring[curr].connections, []*Portal{ring[prev], ring[next]}...)
	}

	for _, pair := range links {
		a, b := pair[0], pair[1]
		ring[a].connections = append(ring[a].connections, ring[b])
		ring[b].connections = append(ring[b].connections, ring[a])
	}

	return ring
}
