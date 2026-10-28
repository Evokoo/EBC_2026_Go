package quest01

import (
	"math"
	"strconv"
	"strings"

	"github.com/Evokoo/EBC_2026_Go/utils"
)

// ========================
// PART I
// ========================

func I(values []int) int {
	var result int
	for i, n := range values {
		for !isSquare(n) {
			n++
		}
		result += n - values[i]
	}
	return result
}

func isSquare(n int) bool {
	if n < 0 {
		return false
	}
	root := int(math.Sqrt(float64(n)))
	return root*root == n
}

// ========================
// PART II
// ========================

func II(values []int) int {
	var result int
	for i, n := range values {
		for n < 8 || n%4 != 0 {
			n++
		}
		result += n - values[i]
	}
	return result
}

// ========================
// PART III
// ========================

func III(values []int) int {
	var result int
	for _, n := range values {
		r := int(math.Sqrt(float64(n)))

		if r%2 == 0 {
			r++
		}

		for r*r-1 < n {
			r += 2
		}

		result += r*r - 1 - n
	}
	return result
}

// ========================
// PARSER
// ========================

func parseInput(input string) []int {
	data := utils.ReadFile(input)

	var values []int
	for n := range strings.SplitSeq(data, ",") {
		num, _ := strconv.Atoi(n)
		values = append(values, num)
	}

	return values
}
