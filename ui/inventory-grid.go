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

func LayoutMaterials(tiletypes []world.TileType) []cxmath.Vec2i {
	bins := binTileTypesByMaterial(tiletypes)
	_ = bins
	positions := []cxmath.Vec2i{}
	
	return positions
}
