package quest04

func Solve(file string, part int) int {
	data := parseInput(file)

	switch part {
	case 1:
		return I(data)
	case 2:
		return II(data)
	case 3:
		return III(data)
	}

	panic("Invalid quest part")
}
