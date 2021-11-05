package worldimport

import (
	"github.com/skycoin/cx-game/render"
)

// a sprite registered from a tiled import
type TiledSprite struct {
	SpriteID render.SpriteID
	Metadata TiledMetadata
}

// properties on a Tiled tileset tile that are relevant to cx-game
type TiledMetadata struct {
	Powered bool
	Name string
}
