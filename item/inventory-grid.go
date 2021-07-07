package item

import (
	"github.com/skycoin/cx-game/cxmath"
	"github.com/skycoin/cx-game/cxmath/math32i"
	"github.com/skycoin/cx-game/world"
)

const InventoryGridWidth = 5

func binTileTypeIDsByMaterial(
		tiletypeIDs []world.TileTypeID,
) map[world.MaterialID][]world.TileTypeID {
	bins := make(map[world.MaterialID][]world.TileTypeID)
	for _,tiletypeID := range tiletypeIDs {
		_,ok := bins[tiletypeID.Get().MaterialID]
		if !ok { bins[tiletypeID.Get().MaterialID] = []world.TileTypeID{} }
		bins[tiletypeID.Get().MaterialID] =
			append(bins[tiletypeID.Get().MaterialID], tiletypeID)
	}
	return bins
}

func GetTileTypesIDsForItemTypeIDs(
		itemtypeIDs []ItemTypeID,
) []world.TileTypeID {
	tiletypeIDs := []world.TileTypeID{}
	for _,itemtypeID := range itemtypeIDs {
		tiletypeID,ok := GetTileTypeIDForItemTypeID(itemtypeID)
		if ok { tiletypeIDs = append(tiletypeIDs, tiletypeID) }
	}
	return tiletypeIDs
}

type PositionedTileTypeID struct {
	TileTypeID world.TileTypeID
	Rect cxmath.Rect
}

func getTileTypeSizes(ids []world.TileTypeID) []cxmath.Vec2i{
	sizes := make([]cxmath.Vec2i,len(ids))
	for idx,id := range ids { sizes[idx] = id.Get().Size() }
	return sizes
}

func LayoutTiletypes(tiletypeIDs []world.TileTypeID) []PositionedTileTypeID {
	bins := binTileTypeIDsByMaterial(tiletypeIDs)
	positionedTileTypeIDs := make([]PositionedTileTypeID,len(tiletypeIDs))
	layoutIdx := 0

	materialYOffset := int32(0)

	for _,bin := range bins {
		sizes := getTileTypeSizes(bin)
		rects := cxmath.PackRectangles(InventoryGridWidth, sizes)
		for binIdx,_ := range rects {
			rect := &rects[binIdx]
			rect.Origin.Y += materialYOffset
			positionedTileTypeIDs[layoutIdx] = PositionedTileTypeID {
				TileTypeID: tiletypeIDs[layoutIdx],
				Rect: *rect,
			}
		}
		for _,rect := range rects {
			materialYOffset = math32i.Max(materialYOffset,rect.Bottom())
		}
	}
	
	return positionedTileTypeIDs
}
