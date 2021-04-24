package tile

import (
	//"log"
	"github.com/skycoin/cx-game/sprite"
)

type Tile struct {
	Name string
	SpriteId int
}

type TileMap struct {
	// store all the tiles with names
	Tiles []Tile
	// layout the stored tiles in some manner
	TileIds []int
	Width, Height int
}

type TilePaleteSelector struct {
	// store tiles for (1) displaying selector and (2) placing tiles
	Tiles []Tile
}

func (tilemap *TileMap) Draw() {
	for idx,tileId := range tilemap.TileIds {
		y := float32(idx / tilemap.Width)
		x := float32(idx % tilemap.Width)
		if tileId>=0 {
			spriteId := tilemap.Tiles[tileId].SpriteId
			sprite.DrawSpriteQuad(x,y,1,1,spriteId)
		}
	}
}

func (selector *TilePaleteSelector) Draw() {
	for _,tile := range selector.Tiles {
		sprite.DrawSpriteQuad(0.0,-1.0,1.0,1.0,tile.SpriteId)
	}
}
