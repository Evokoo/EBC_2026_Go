package quest08

func Solve(file string, part int) int {
	tiles := parseInput(file)
	tiles.LocateEdges()

	switch part {
	case 1:
		return I(tiles)
	case 2:
		return II(tiles)
	}

	panic("Invalid quest part!")
}
