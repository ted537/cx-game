package worldimport

import (
	"github.com/lafriks/go-tiled"
)

// a sprite registered from a tiled import
type TiledSprite struct {
	Image    TilesetTileImage
	Metadata TiledMetadata
}

// properties on a Tiled tileset tile that are relevant to cx-game
type TiledMetadata struct {
	Powered OptionalBool
	Name string
}

type OptionalBool struct {
	Set bool
	Value bool
}

type TileSprites []TiledSprite
type TiledSprites map[string]TileSprites

func RegisterTiledSprites(tiledMap *tiled.Map, mapDir string) TiledSprites {
	tiledSprites := TiledSprites{}

	for _,tileset := range tiledMap.Tilesets {
		for _,tilesetTile := range tileset.Tiles {
			metadata := parseMetadata(tilesetTile.Properties)
			image := imageForTilesetTile(
				tileset, tilesetTile.ID, tilesetTile, mapDir )
			tiledSprite := TiledSprite { Image: image, Metadata: metadata }
			tiledSprites[metadata.Name] =
				append(tiledSprites[metadata.Name], tiledSprite)
		}
	}

	return tiledSprites
}

func hasProperty(properties tiled.Properties, name string) bool {
	for _,property := range properties {
		if property.Name == name {
			return true
		}
	}
	return false
}

func parseMetadata(properties tiled.Properties) TiledMetadata {
	metadata := TiledMetadata{}
	if hasProperty(properties, "powered") {
		metadata.Powered.Set = true
		metadata.Powered.Value = properties.GetBool("powered")
	}
	return metadata
}
