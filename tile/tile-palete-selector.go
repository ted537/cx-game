package tile

import (
	"log"
	"github.com/skycoin/cx-game/spriteloader"
)

type TileMap struct {
	TileIds []int
	Width, Height int
}

type TilePaleteSelector struct {
	spritesheet spriteloader.Spritesheet
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
