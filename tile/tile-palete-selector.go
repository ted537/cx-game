package tile

import (
	//"log"
	"math"
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

const paleteXOffset = 0.0
const paleteYOffset = -3.0
func (selector *TilePaleteSelector) Draw() {
	numTiles := float64(len(selector.Tiles))
	if numTiles>0 {
		width := math.Ceil(math.Sqrt(numTiles))
		scale := float32(1.0/width)
		for idx,tile := range selector.Tiles {
			yLocal := float32(idx/int(width))*scale
			xLocal := float32(idx%int(width))*scale
			//log.Print(tile.SpriteId)
			//log.Print(tile.Name)
			sprite.DrawSpriteQuad(
				float32(paleteXOffset+xLocal),
				float32(paleteYOffset+yLocal),
				scale,scale,
				tile.SpriteId,
			)
		}
	}
}
