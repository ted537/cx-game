package world

import (
	"github.com/skycoin/cx-game/render"
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
	LightSource       bool
    NeedsGround       bool
}

func NewEmptyTile() Tile {
	return Tile{TileCategory: TileCategoryNone}
}

func NewNormalTile() Tile {
    return Tile {
        TileCategory: TileCategoryNormal,
        TileCollisionType: TileCollisionTypeSolid,
    }
}
