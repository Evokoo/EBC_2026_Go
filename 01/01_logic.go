package quest01

import (
	"math"
	"strconv"
	"strings"

	"github.com/Evokoo/EBC_2026_Go/utils"
)

// ========================
// I
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
