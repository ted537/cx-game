package world

import (
	"log"
	"github.com/skycoin/cx-game/world/tiling"
)

type Connections struct { Up, Left, Right, Down bool }

func (c Connections) Bits() [4]bool {
	return [4]bool { c.Up, c.Left, c.Right, c.Down }
}

func (c Connections) Valid(possible Connections) bool {
	isValid :=
		( possible.Up || !c.Up  ) &&
		( possible.Down || !c.Down  ) &&
		( possible.Left || !c.Left  ) &&
		( possible.Right || !c.Right  )
	return isValid
}

func ConnectionsFromNeighbours(n tiling.DetailedNeighbours) Connections {
	s := n.Simplify()
	return Connections { Up: s.Up, Left: s.Left, Right: s.Right, Down: s.Down }
}

func ConnectedNeighbours(
		connections Connections, neighbours tiling.DetailedNeighbours,
) tiling.DetailedNeighbours {
	connectedNeighbours := neighbours // copy
	// hide neighbours which we shouldn't see due to connections
	if !connections.Up { connectedNeighbours.Up = tiling.None }
	if !connections.Down { connectedNeighbours.Down = tiling.None }
	if !connections.Left { connectedNeighbours.Left = tiling.None }
	if !connections.Right { connectedNeighbours.Right = tiling.None }
	return connectedNeighbours
}

func composeBits(bits []bool) int {
	place := 1
	sum := 0
	for _,bit := range bits {
		if bit { sum += place }
		place *= 2
	}
	return sum
}

func decomposeBits(composed int, bits []bool) {
	place := 1
	for idx := range bits {
		bits[idx] = composed&place != 0
		place *= 2
	}
}

// given some current connection state, cycles to another connection state.
// loops over all possible states eventually
func (c Connections) Next(valid Connections) Connections {
	bits := []bool { c.Up, c.Left, c.Right, c.Down }
	log.Printf("bits=%v",bits)
	i := composeBits(bits)
	for true {
		decomposed := [4]bool{}
		decomposeBits(i+1,decomposed[:])
		d := decomposed
		candidate := Connections { d[0], d[1], d[2], d[3] }
		if candidate.Valid(valid) { return candidate }
		i = (i+1) % 16
	}
	return Connections{} // unreachable anyway
}
