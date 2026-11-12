package quest09

func Solve(file string, part int) int {
	seeds := parseInput(file)

	switch part {
	case 1:
		return I(seeds, 1)
	case 2:
		return I(seeds, 2)
	}

	panic("Invalid quest part")
}
