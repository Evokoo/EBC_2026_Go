package quest04

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/Evokoo/EBC_2026_Go/utils"
)

// ========================
// TYPES
// ========================
type UserResult struct {
	parts, time, points int
}
type Completion struct {
	id, quest, part, time int
}

// ========================
// PART I
// ========================
func I(data []Completion) int {
	users := make(map[int]int)

	for _, completion := range data {
		users[completion.id] += completion.time
	}

	keys := make([]int, 0, len(users))
	for key := range users {
		keys = append(keys, key)
	}

	slices.SortFunc(keys, func(a, b int) int {
		return cmp.Compare(users[a], users[b])
	})

	var checksum int
	for i, key := range keys {
		checksum += (i + 1) * users[key]
	}

	return checksum
}

// ========================
// PART II
// ========================
func II(data []Completion) int {
	users := make(map[int]UserResult)

	for _, completion := range data {
		res := users[completion.id]

		res.parts++
		res.time += completion.time

		users[completion.id] = res
	}

	keys := make([]int, 0, len(users))
	for key := range users {
		keys = append(keys, key)
	}

	slices.SortFunc(keys, func(a, b int) int {
		if res := cmp.Compare(users[b].parts, users[a].parts); res != 0 {
			return res
		}
		return cmp.Compare(users[a].time, users[b].time)
	})

	var checksum int
	for i, key := range keys {
		checksum += (i + 1) * users[key].time
	}

	return checksum
}

// ========================
// PART III
// ========================
func III(data []Completion) int {
	users := make(map[int]UserResult)
	order := make(map[string][][2]int)

	for _, completion := range data {
		key := fmt.Sprintf("%d-%d", completion.quest, completion.part)
		order[key] = append(order[key], [2]int{completion.id, completion.time})

		res := users[completion.id]

		res.parts++
		res.time += completion.time

		users[completion.id] = res
	}

	//apply points
	for key := range order {
		keyParts := strings.Split(key, "-")
		userOrder := order[key]

		slices.SortFunc(userOrder, func(a, b [2]int) int {
			return cmp.Compare(a[1], b[1])
		})

		var points int
		switch keyParts[1] {
		case "1":
			points = 50
		case "2":
			points = 100
		case "3":
			points = 150
		}

		user := users[userOrder[0][0]]
		user.points += points
		users[userOrder[0][0]] = user
	}

	keys := make([]int, 0, len(users))
	for key := range users {
		keys = append(keys, key)
	}

	slices.SortFunc(keys, func(a, b int) int {
		return cmp.Compare(users[a].points, users[b].points)
	})

	for _, key := range keys {
		fmt.Println(users[key])
	}

	var checksum int
	for i, key := range keys {
		checksum += (i + 1) * users[key].time
	}

	return checksum
}

// ========================
// PARSER
// ========================
func parseInput(file string) []Completion {
	data := utils.ReadFile(file)

	var results []Completion
	for line := range strings.SplitSeq(data, "\n") {
		details := utils.QuickMatch(line, `\d+`)

		var result Completion
		for i, detail := range details {
			n, _ := strconv.Atoi(detail)
			switch i {
			case 0:
				result.id = n
			case 1:
				result.quest = n
			case 2:
				result.part = n
			case 3:
				result.time = n
			}
		}

		results = append(results, result)
	}

	return results
}
