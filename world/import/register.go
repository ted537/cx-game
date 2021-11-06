package worldimport

import (
	"fmt"
	"path"
	"image"

	"github.com/lafriks/go-tiled"
	"github.com/go-gl/mathgl/mgl32"

	"github.com/skycoin/cx-game/world"
	"github.com/skycoin/cx-game/engine/spriteloader"
	"github.com/skycoin/cx-game/render"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type TileRegistrationOptions struct {
	TmxPath string
	LayerID world.LayerID

	Tileset *tiled.Tileset
	LayerTile *tiled.LayerTile
	TilesetTile *tiled.TilesetTile

	FlipTransform mgl32.Mat3
}

type TilesetTileImage struct {
	Path string
	SpriteTransform mgl32.Mat3
	Width int32 // measured in tiles
	Height int32
}

func (t TilesetTileImage) Model() mgl32.Mat4 {
	return mgl32.Scale3D(float32(t.Width), float32(t.Height), 1)
}

func (t TilesetTileImage) RegisterSprite(name string) render.SpriteID {
	texture :=
		spriteloader.LoadTextureFromFileToGPUCached(t.Path)
	sprite := render.Sprite{
		Name:      name,
		Transform: t.SpriteTransform,
		Model:     t.Model(),
		Texture:   render.Texture{Target: gl.TEXTURE_2D, Texture: texture.Gl},
	}
	return render.RegisterSprite(sprite)
}

func registerTilesetTile(
	layerTile *tiled.LayerTile, opts TileRegistrationOptions,
) world.TileTypeID {
	name := fmt.Sprintf("%v:%v", opts.Tileset.Name, opts.LayerTile.ID)
	pathPrefix := path.Join(opts.TmxPath, "..")
	tilesetTileImage := imageForTilesetTile(
		opts.Tileset, layerTile.ID, opts.TilesetTile, pathPrefix )
	spriteID := tilesetTileImage.RegisterSprite(name)

	tile := world.NewNormalTile()
	tile.Name = name
	tile.TileTypeID = world.NextTileTypeID()

	tileType := world.TileType{
		Name:   name,
		Layer:  opts.LayerID,
		Placer: world.DirectPlacer{SpriteID: spriteID, Tile: tile},
		Width:  tilesetTileImage.Width, Height: tilesetTileImage.Height,
	}

	tileTypeID :=
		world.RegisterTileType(name, tileType, defaltToolForLayer(opts.LayerID))

	return tileTypeID
}

func imageForTilesetTile(
	//layerTile *tiled.LayerTile, opts TileRegistrationOptions,
	tileset *tiled.Tileset, tileID uint32, tilesetTile *tiled.TilesetTile,
	pathPrefix string,
) TilesetTileImage {
	if tilesetTile != nil && tilesetTile.Image != nil {
		tileImg := tilesetTile.Image
		imgPath := path.Join(pathPrefix, tileImg.Source)
		return TilesetTileImage {
			Path: imgPath,
			SpriteTransform: mgl32.Ident3(),
			Width: int32(tileImg.Width)/16, Height: int32(tileImg.Height)/16,
		}
	} else {
		imgPath := path.Join(pathPrefix, tileset.Image.Source)
		spriteRect := tileset.GetTileRect(tileID)
		tilesetDims := image.Point { tileset.Image.Width, tileset.Image.Height }
		spriteTransform := rectTransform( spriteRect, tilesetDims )
		return TilesetTileImage {
			Path: imgPath,
			SpriteTransform: spriteTransform,
			Width: 1, Height: 1,
		}
	}
}
