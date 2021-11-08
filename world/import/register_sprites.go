package worldimport

import (
	"log"

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

func NewTiledMetadata(name string) TiledMetadata {
	return TiledMetadata { Name: name }
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
		log.Printf("registering sprites for tileset %v", tileset.Name)
		registeredTileIDs := map[uint32]bool{}
		for _,tilesetTile := range tileset.Tiles {
			name := nameForTilesetTile(tileset.Name, tilesetTile.ID)
			metadata := NewTiledMetadata(name)
			metadata.ParseFrom(tilesetTile.Properties)
			image := imageForTilesetTile(
				tileset, tilesetTile.ID, tilesetTile, mapDir )
			tiledSprite := TiledSprite { Image: image, Metadata: metadata }
			tiledSprites[metadata.Name] =
				append(tiledSprites[metadata.Name], tiledSprite)
			registeredTileIDs[tilesetTile.ID] = true
		}
		if tileset.Image != nil {
			for id := uint32(0) ; id < uint32(tileset.TileCount) ; id++ {
				name := nameForTilesetTile(tileset.Name, id)
				metadata := NewTiledMetadata(name)
				isRegistered,_ := registeredTileIDs[id]
				if !isRegistered {
					image :=
						imageForTilesetTile(tileset, uint32(id), nil, mapDir)
					tiledSprite :=
						TiledSprite { Image: image, Metadata: metadata }
					tiledSprites[metadata.Name] =
						append(tiledSprites[metadata.Name], tiledSprite)
				}
			}
		}
	}

	for name,_ := range tiledSprites {
		log.Printf("TMX: %v",name)
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

func (metadata *TiledMetadata) ParseFrom(properties tiled.Properties) {
	if hasProperty(properties, "powered") {
		metadata.Powered.Set = true
		metadata.Powered.Value = properties.GetBool("powered")
	}
}

func parseMetadataFromLayerTile(layerTile *tiled.LayerTile) TiledMetadata {
	name := nameForLayerTile(layerTile)
	metadata := NewTiledMetadata(name)
	tilesetTile, ok := findTilesetTileForLayerTile(layerTile)
	if ok {
		metadata.ParseFrom(tilesetTile.Properties)
	}
	return metadata
}
