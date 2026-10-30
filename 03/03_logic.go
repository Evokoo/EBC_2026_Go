package quest03

import (
	"strconv"
	"strings"

	"github.com/Evokoo/EBC_2026_Go/utils"
)

// ========================
// PART I && II
// ========================
func I(pairs [][2]int) int {
	var result int

	for _, pair := range pairs {
		result += calculateISO(pair)
	}

	return result
}

// ========================
// PART III
// ========================
func III(pairs [][2]int) int {
	var result int

	for _, pair := range pairs {
		a := pair[0]
		b := pair[1]

		target := calculateISO([2]int{a, b})

		for b > 0 {
			b--
			next := calculateISO([2]int{a, b})

			if next == target {
				target = 0
				break
			}
		}
		result += target
	}

	return result
}

// ========================
// HELPERS
// ========================
func calculateISO(pair [2]int) int {
	a := pair[0]
	b := pair[1]
	ex := 0
	sum := 0

	for a > 0 && b > 0 {
		sum += b % 10 * toPower(a%10, ex)
		a /= 10
		b /= 10
		ex++
	}
	return sum
}

func toPower(n, ex int) int {
	result := 1
	for range ex {
		result *= n
	}
	return result
}

// ========================
// PARSER
// ========================

func parseInput(file string) [][2]int {
	data := utils.ReadFile(file)

	var pairs [][2]int

	for line := range strings.SplitSeq(data, "\n") {
		values := strings.Split(line, "-")
		matches := utils.QuickMatch(values[1], `\d+`)

		a := matches[0]
		b := matches[1]
		for len(a) < len(b) {
			a += a
		}

		n1, _ := strconv.Atoi(a[len(a)-len(b):])
		n2, _ := strconv.Atoi(b)
		pairs = append(pairs, [2]int{n1, n2})
	}

	return pairs
}
