package main

import (
	"github.com/skycoin/cx-game/tile"
)

func main() {
	tilemap := tile.TileMap {
		[]int{1,0,3,4},
		2,2,
	}
	tilemap.Draw()
}
