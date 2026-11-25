package quest13

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Evokoo/EBC_2026_Go/utils"
)

// ========================
// I
// ========================
func I(dial Dial) int {
	result := 0

	for {
		if overlaps, canMove := dial.Update(); canMove {
			fmt.Println(overlaps)
			result += overlaps
		} else {
			return result
		}
	}
}

// ========================
// PARSER
// ========================

type Dial []*Hand

func (d Dial) Update() (int, bool) {
	//Check for valid moves
	remaining := 0
	for _, hand := range d {
		remaining += hand.moves
	}

	if remaining == 0 {
		return 0, false
	}

	//Make moves
	occupied := make(map[int]int)
	for _, hand := range d {
		hand.Update()
		occupied[hand.current]++
	}

	//Check for overlaps
	fmt.Println(occupied)
	overlaps := 0
	for _, count := range occupied {
		if count > 1 {
			overlaps += count - 1
		}
	}

	return overlaps, true
}

type Hand struct {
	positions, current, moves int
	direction                 string
}

func (h *Hand) Update() {
	if h.moves == 0 {
		return
	}

	step := h.moves % h.positions
	if h.direction == "CW" {
		h.current = (h.current + step) % h.positions
	} else {
		h.current = (h.current - step + h.positions) % h.positions
	}

	h.moves--
}

func parseInput(file string) Dial {
	data := utils.ReadFile(file)

	var dial Dial
	for range 3 {
		dial = append(dial, &Hand{})
	}

	for i, line := range strings.Split(data, "\n") {
		details := strings.Split(line, " ")
		values := strings.Split(details[1], ":")
		for j, value := range values {
			switch i {
			case 0:
				n, _ := strconv.Atoi(value)
				dial[j].positions = n
			case 1:
				dial[j].direction = value
			case 2:
				n, _ := strconv.Atoi(value)
				dial[j].current = n
			case 3:
				n, _ := strconv.Atoi(value)
				dial[j].moves = n
			}
		}
	}

	return dial
}
