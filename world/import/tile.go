package worldimport

import (
	"image"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/lafriks/go-tiled"

	"github.com/skycoin/cx-game/world"
	"github.com/skycoin/cx-game/components/types"
	"github.com/skycoin/cx-game/constants"
)


func defaltToolForLayer(layerID world.LayerID) types.ToolType {
	if layerID == world.BgLayer {
		return constants.BG_TOOL
	}
	return constants.FURNITURE_TOOL
}

func findTilesetTileForLayerTile(
	layerTile *tiled.LayerTile,
) (*tiled.TilesetTile, bool) {
	for _, tilesetTile := range layerTile.Tileset.Tiles {
		if tilesetTile.ID == layerTile.ID {
			return tilesetTile, true
		}
	}
	return nil, false
}

// compute a 3x3 matrix which maps every point in the
// "parent" rectangle to a corresponding point in the "here" rectangle
func rectTransform(here image.Rectangle, parentDims image.Point) mgl32.Mat3 {
	scaleX := float32(here.Dx()) / float32(parentDims.X)
	scaleY := float32(here.Dy()) / float32(parentDims.Y)
	scale := mgl32.Scale2D(scaleX, scaleY)

	translate := mgl32.Translate2D(
		float32(here.Min.X)/float32(parentDims.X),
		float32(here.Min.Y)/float32(parentDims.Y))

	return translate.Mul3(scale)
}


type TilesetIDKey struct {
	tileset *tiled.Tileset
	id      uint32
}

var tilesetAndIDToCXTile = map[TilesetIDKey]world.TileTypeID{}

func directPlacerForLayerTiles(layerTiles []tiled.LayerTile) world.Placer {
	spriteID := layerTiles[0].SpriteID
	return world.DirectPlacer { SpriteID: spriteID, Tile: tile }
}

func powerPlacerForLayerTiles(layerTiles []tiled.LayerTile) world.Placer {

}

func placerForLayerTiles(layerTiles []tiled.LayerTile) world.Placer {
	for _,layerTile := range layerTiles {
		metadata := parseMetadataFromLayerTile(layerTile)
		if metadata.Powered.Set {
			return powerPlacerForLayerTiles(layerTiles)
		}
	}
	return directPlacerForLayerTiles(layerTiles)
}

func registerTiledTile(layerTiles []world.LayerTile) world.TileTypeID {
	placer := placerForLayerTiles(layerTiles)

}

func getTileTypeID(
	layerTile *tiled.LayerTile, tmxPath string, layerID world.LayerID,
	tiledSprites TiledSprites,
) world.TileTypeID {
	tileset := layerTile.Tileset
	// nil entry => empty layer tile
	if tileset == nil {
		return world.TileTypeIDAir
	}

	// search for tile in existing tiles
	tilesetTile, foundTilesetTile := findTilesetTileForLayerTile(layerTile)

	if foundTilesetTile {
		cxtile := tilesetTile.Properties.GetString("cxtile")
		tileTypeID, foundTileTypeID := world.IDFor(cxtile)
		if foundTileTypeID {
			return tileTypeID
		}
	}

	flipX,flipY := scaleFromFlipFlags(layerTile)
	flipTransform := mgl32.Scale2D( float32(flipX), float32(flipY) )
	key := TilesetIDKey{ tileset, layerTile.ID }
	cachedTileTypeID, hitCache := tilesetAndIDToCXTile[key]
	if hitCache {
		return cachedTileTypeID
	}

	// did not find - register new tile type
	return registerTilesetTile(layerTile, TileRegistrationOptions {
		TmxPath: tmxPath, LayerID: layerID, Tileset: tileset,
		LayerTile: layerTile, TilesetTile: tilesetTile,
		FlipTransform: flipTransform,
		TiledSprites: tiledSprites,
	})
}
