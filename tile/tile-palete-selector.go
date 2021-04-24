package tile

import (
	"log"
	"github.com/skycoin/cx-game/sprite"
)

type TileMap struct {
	spritesheet *sprite.Spritesheet
	TileIds []int
	Width, Height int
}

type TilePaleteSelector struct {
	spritesheet sprite.Spritesheet
}

func (tilemap *TileMap) Draw() {
	for idx,id := range tilemap.TileIds {
		y := idx / tilemap.Width
		x := idx % tilemap.Width
		if id!=0 {
			log.Print(x,y,id)
		}
	}
}

func (selector *TilePaleteSelector) Draw() {

}
