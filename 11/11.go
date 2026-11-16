package quest11

func Solve(file string, part int) int {
	switch part {
	case 1:
		return I(parseInput(file, 1))
	case 2:
		return I(parseInput(file, 3))
	}

	panic("Invalid quest part")
}
