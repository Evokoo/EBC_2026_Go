package quest03

func Solve(file string, part int) int {
	pairs := parseInput(file)

	switch part {
	case 1, 2:
		return I(pairs)
	case 3:
		return III(pairs)
	}

	panic("Invalid quest part")
}
