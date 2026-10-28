package quest01

func Solve(file string, part int) int {
	values := parseInput(file)

	switch part {
	case 1:
		return I(values)
	case 2:
		return II(values)
	case 3:
		return III(values)
	}

	panic("Invalid Quest part")
}
