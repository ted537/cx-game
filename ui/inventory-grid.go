package ui

import (
	"github.com/skycoin/cx-game/world"
	"github.com/skycoin/cx-game/cxmath"
)

const InventoryGridWidth = 5

func binTileTypesByMaterial(
		tiletypes []world.TileType,
) map[world.MaterialID][]world.TileType {
	bins := make(map[world.MaterialID][]world.TileType)
	for _,tiletype := range tiletypes {
		_,ok := bins[tiletype.MaterialID]
		if !ok { bins[tiletype.MaterialID] = []world.TileType{} }
		bins[tiletype.MaterialID] =
			append(bins[tiletype.MaterialID], tiletype)
	}
	return bins
}

func LayoutTiletypes(tiletypes []world.TileType) []cxmath.Vec2i {
	bins := binTileTypesByMaterial(tiletypes)
	positions := []cxmath.Vec2i{}

	y := int32(0)
	x := int32(0)
	for _,bin := range bins {
		for _,tiletype := range bin {
			_ = tiletype // no use currently, might throw in struct later
			x++
			if x==InventoryGridWidth { x=0; y++ }
			positions = append(positions, cxmath.Vec2i{x,y})
		}
		if x>0 { y++ } // after each bin, we must go to the next row
	}
	
	return positions
}
