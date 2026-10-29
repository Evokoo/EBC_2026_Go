package quest02

func Solve(file string, part int) int {
	alloys := parseInput(file)

	switch part {
	case 1:
	case 2:
		return I(alloys)
	case 3:
		return III(alloys)
	}

	panic("Quest part invalid")
}
