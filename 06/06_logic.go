package quest06

import (
	"strings"

	"github.com/Evokoo/EBC_2026_Go/utils"
)

// ========================
// PART I
// ========================

func I(recipes Recipes, target string) int {
	used := make(utils.Set[string])
	DFS(recipes[target], &used, &recipes)
	used.Remove(target)
	return len(used)
}

// ========================
// PART II
// ========================

func II(recipes Recipes, target string) int {
	result := 0

	for item := range recipes {
		if item != target {
			used := make(utils.Set[string])
			DFS(recipes[item], &used, &recipes)

			if used.Has(target) {
				result++
			}
		}
	}

	return result
}

// ========================
// PART III
// ========================

func III(recipes Recipes) int {
	result := 0

	for item := range recipes {
		used := make(utils.Set[string])
		DFS(recipes[item], &used, &recipes)

		count := 0
		for ingredient := range used {
			if ingredient == "NULLv1" || ingredient == "NULLv2" || ingredient == "NULLv3" {
				count++
			}
		}

		if count >= 2 {
			result++
		}

	}

	return result
}

// ========================
// DFS
// ========================
func DFS(recipe []string, used *utils.Set[string], recipes *Recipes) {
	for _, item := range recipe {
		if !used.Has(item) {
			used.Add(item)

			if ingredients, found := (*recipes)[item]; found {
				DFS(ingredients, used, recipes)
			}
		}
	}
}

// ========================
// PARSER
// ========================

type Recipes map[string][]string

func parseInput(file string) Recipes {
	data := utils.ReadFile(file)

	result := make(Recipes)
	for line := range strings.SplitSeq(data, "\n") {
		details := strings.Split(line, ":")
		result[details[0]] = strings.Split(details[1], ",")
	}

	return result
}
