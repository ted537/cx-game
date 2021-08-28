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
