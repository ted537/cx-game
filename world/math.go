package world

type ManhattanNeighbour struct {
	X,Y int
	Mask Connections
}

func manhattanNeighbours(x,y int) []ManhattanNeighbour {
	return []ManhattanNeighbour {
		ManhattanNeighbour {
			x-1, y,
			Connections { Up: true, Left: true, Down: true, Right: false },
		},
		ManhattanNeighbour {
			x+1, y,
			Connections { Up: true, Left: false, Down: true, Right: true },
		},
		ManhattanNeighbour {
			x, y-1,
			Connections { Up: false, Left: true, Down: true, Right: true },
		},
		ManhattanNeighbour {
			x, y+1,
			Connections { Up: true, Left: true, Down: false, Right: true },
		},
	}
}
