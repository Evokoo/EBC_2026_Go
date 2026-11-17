package quest11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Evokoo/EBC_2026_Go/utils"
)

// ========================
// I
// ========================

func I(ring Ring) int {
	return ring.CalculateMinRoute()
}

// ========================
// II
// ========================

func II(ring Ring) int {
	result := 0

	for a, p1 := range ring.portals {
		for b, p2 := range ring.portals {
			if b <= a || !ring.LinkPortals(p1, p2) {
				continue
			}
			result = max(ring.CalculateMinRoute(), result)
			ring.UnlinkPortals(p1, p2)
		}
	}

	return result
}

// ========================
// III
// ========================

func III(ring Ring) int {
	// Generate indexs and pair permutations
	steps := len(ring.portals) / 8
	indexs := make([]int, 0, steps)

	for i := 0; i < len(ring.portals); i += steps {
		indexs = append(indexs, i+1)
	}

	fmt.Println(indexs)

	return 0
}

// ========================
// PARSER
// ========================

type Portal struct {
	id          int
	connections utils.Set[*Portal]
}

type Ring struct {
	portals      map[int]*Portal
	baseDistance int
}

func newRing() Ring {
	return Ring{portals: make(map[int]*Portal), baseDistance: 0}
}

func (r Ring) DistanceMap(start int) int {
	queue := [][2]int{{start, 0}}
	distanceMap := make(map[int]int)

	for len(queue) > 0 {
		current := queue[0]
		id, dist := current[0], current[1]
		queue = queue[1:]

		if _, found := distanceMap[id]; found {
			continue
		} else {
			distanceMap[id] = dist
		}

		for portal := range r.portals[id].connections {
			queue = append(queue, [2]int{portal.id, dist + 1})
		}
	}

	total := 0
	for _, dist := range distanceMap {
		total += dist
	}

	return total
}
func (r Ring) CalculateMinRoute() int {
	result := 0
	for _, portal := range r.portals {
		result += r.DistanceMap(portal.id)
	}
	return r.baseDistance - result
}
func (r Ring) LinkPortals(a, b *Portal) bool {
	if a.connections.Has(b) {
		return false
	}
	a.connections.Add(b)
	b.connections.Add(a)
	return true
}
func (r Ring) UnlinkPortals(a, b *Portal) {
	a.connections.Remove(b)
	b.connections.Remove(a)
}

func parseInput(file string, spread int) Ring {
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

	ring := newRing()

	for i := range count {
		ring.portals[i+1] = &Portal{id: i + 1, connections: make(utils.Set[*Portal])}
	}

	for i := range count {
		curr := i + 1
		for d := 1; d <= spread; d++ {
			next := (i+d)%count + 1
			prev := (i-d+count)%count + 1
			ring.portals[curr].connections.Add(ring.portals[prev])
			ring.portals[curr].connections.Add(ring.portals[next])
		}
	}

	for _, portal := range ring.portals {
		ring.baseDistance += ring.DistanceMap(portal.id)
	}

	for _, pair := range links {
		a, b := pair[0], pair[1]
		ring.portals[a].connections.Add(ring.portals[b])
		ring.portals[b].connections.Add(ring.portals[a])
	}

	return ring
}
