package quest10

func Solve(file string, part int) int {
	field := parseInput(file)

	switch part {
	case 1:
		return I(field, 10)
	case 2:
		return II(field, 10000)
	case 3:
		return II(field, 1000000000000)
	}

	panic("Invalid quest part")
}
