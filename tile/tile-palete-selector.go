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
	spritesheet sprite.Spritesheet
}

func (tilemap *TileMap) Draw() {
	for idx,tileId := range tilemap.TileIds {
		y := idx / tilemap.Width
		x := idx % tilemap.Width
		if tileId>=0 {
			spriteId := tilemap.Tiles[tileId].SpriteId
			sprite.DrawSpriteQuad(x,y,1,1,spriteId)
		}
	}
}

func (selector *TilePaleteSelector) Draw() {

}
