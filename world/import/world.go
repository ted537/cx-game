package worldimport

import (
	"github.com/go-gl/mathgl/mgl32"

	"github.com/lafriks/go-tiled"
	"github.com/skycoin/cx-game/world"
)

func importTile(
	planet *world.Planet,
	tileIndex int, layerTile *tiled.LayerTile, tmxPath string,
	layerID world.LayerID, tiledSprites TiledSprites,
) {
	tileTypeID := getTileTypeID(layerTile, tmxPath, layerID, tiledSprites)
	if tileTypeID != world.TileTypeIDAir {

		// correct mismatch between Tiled Y axis (downwards)
		// and our Y axis  (upwards)
		y := int(planet.Height) - tileIndex/int(planet.Width)
		x := tileIndex % int(planet.Width)

		opts := world.NewTileCreationOptions()
		flipX, flipY := scaleFromFlipFlags(layerTile)
		opts.FlipTransform = mgl32.Scale3D( float32(flipX), float32(flipY), 1 )
		planet.PlaceTileType(tileTypeID, x, y, opts)
	}
}

func importLayer(
	planet *world.Planet, tiledLayer *tiled.Layer, tmxPath string,
	layerID world.LayerID, tiledSprites TiledSprites,
) {
	for idx, layerTile := range tiledLayer.Tiles {
		importTile(planet, idx, layerTile, tmxPath, layerID, tiledSprites)
	}
}

func findTiledSpritesInMap(
		allTiledSprites TiledSprites, tiledMap *tiled.Map,
) TiledSprites {
	mapTiledSprites := TiledSprites{}
	for _,layer := range tiledMap.Layers {
		for _,layerTile := range layer.Tiles {
			if !layerTile.Nil {
				name := nameForLayerTile(layerTile)
				mapTiledSprites[name] = allTiledSprites[name]
			}
		}
	}

	return mapTiledSprites
}
