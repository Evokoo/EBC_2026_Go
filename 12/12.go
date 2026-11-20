package quest12

func Solve(file string, part int) int {
	shapes := parseInput(file)

	switch part {
	case 1:
		return I(shapes)
	}

	panic("Invalid quest part")
}
