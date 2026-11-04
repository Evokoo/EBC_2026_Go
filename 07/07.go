package quest07

func Solve(file string, part int) int {
	clouds := parseInput(file)

	switch part {
	case 1:
		return I(clouds, 5)
	case 2:
		return I(clouds, 10)
	case 3:
		return III(clouds, 10, 100)
	}

	panic("Invalid quest part")
}
