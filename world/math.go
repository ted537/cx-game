package world

type BoolDiff int
const (
	NO_CHANGE BoolDiff = iota
	SET_ON
	SET_OFF
)
func computeBoolDiff(x,y bool) BoolDiff {
	if y && !x { return SET_ON }
	if !y && x { return SET_OFF }
	return NO_CHANGE
}
func applyBoolDiff(x bool, diff BoolDiff) bool {
	if diff == SET_OFF { return false }
	if diff == SET_ON { return true }
	return x
}

type ConnectionDiff struct { Up,Left,Right,Down BoolDiff }

type ManhattanNeighbour struct {
	X,Y int
	ConnectionDiff ConnectionDiff
}

func manhattanNeighbours(
		x,y int, connections, oldConnections Connections,
) []ManhattanNeighbour {
	diff := oldConnections.Diff(connections)
	return []ManhattanNeighbour {
		ManhattanNeighbour {
			x-1, y,
			ConnectionDiff { Right: diff.Left },
		},
		ManhattanNeighbour {
			x+1, y,
			ConnectionDiff { Left: diff.Right },
		},
		ManhattanNeighbour {
			x, y-1,
			ConnectionDiff { Up: diff.Down },
		},
		ManhattanNeighbour {
			x, y+1,
			ConnectionDiff { Down: diff.Up },
		},
	}
}
