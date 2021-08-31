package world

import (
	"github.com/skycoin/cx-game/render"
	"github.com/skycoin/cx-game/world/tiling"
)

type TileCategory uint32

const (
	TileCategoryNone TileCategory = iota
	TileCategoryNormal
	TileCategoryMulti
	TileCategoryChild
	TileCategoryLiquid
)

type TileCollisionType uint32

const (
	TileCollisionTypeSolid TileCollisionType = iota
	TileCollisionTypePlatform
)

func (tt TileCategory) ShouldRender() bool {
	return tt != TileCategoryNone
}

type Connections struct { Up, Left, Right, Down bool }

func ConnectionsFromNeighbours(n tiling.DetailedNeighbours) Connections {
	s := n.Simplify()
	return Connections { Up: s.Up, Left: s.Left, Right: s.Right, Down: s.Down }
}

func ConnectedNeighbours(
		connections Connections, neighbours tiling.DetailedNeighbours,
) tiling.DetailedNeighbours {
	connectedNeighbours := neighbours // copy
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
	return place
}

func decomposeBits(composed int, bits []bool) {
	place := 2
	for idx := range bits {
		bits[idx] = composed%place == 1
	}
}

// given some current connection state, cycles to another connection state.
// loops over all possible states eventually
func (c Connections) Next(valid Connections) Connections {
	bits := []bool { c.Up, c.Left, c.Right, c.Down }
	i := composeBits(bits)
	decomposed := [4]bool{}
	decomposeBits(i,decomposed[:])
	d := decomposed
	return Connections { d[0], d[1], d[2], d[3] }
}

type Tile struct {
	SpriteID          render.SpriteID
	TileCategory      TileCategory
	TileCollisionType TileCollisionType
	TileTypeID        TileTypeID
	Name              string
	OffsetX           int8
	OffsetY           int8
	Durability        int8
	Connections       Connections
}

func NewEmptyTile() Tile {
	return Tile{TileCategory: TileCategoryNone}
}
