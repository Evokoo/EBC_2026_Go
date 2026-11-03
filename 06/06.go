package quest06

func Solve(file string, part int) int {
	recipes := parseInput(file)

	switch part {
	case 1:
		return I(recipes, "NULL")
	case 2:
		return II(recipes, "NULL")
	case 3:
		return III(recipes)
	}

	panic("Invalid quest part")
}
