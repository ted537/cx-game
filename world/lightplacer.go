package world

import (
	"github.com/skycoin/cx-game/render"
)

type LightPlacer struct {
	Tile Tile
	OnSpriteID, OffSpriteID render.SpriteID
}

func (placer *LightPlacer) CreateTile(
	tt TileType, opts TileCreationOptions,
) Tile {
	tile := placer.Tile
	if tile.
}
