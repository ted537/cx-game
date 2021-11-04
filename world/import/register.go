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

func registerTilesetTile(
	layerTile *tiled.LayerTile, opts TileRegistrationOptions,
) world.TileTypeID {
	tilesetTileImage := imageForTilesetTile(layerTile, opts)
	texture :=
		spriteloader.LoadTextureFromFileToGPUCached(tilesetTileImage.Path)
	name := fmt.Sprintf("%v:%v", opts.Tileset.Name, opts.LayerTile.ID)
	sprite := render.Sprite{
		Name:      name,
		Transform: tilesetTileImage.SpriteTransform,
		Model:     tilesetTileImage.Model(),
		Texture:   render.Texture{Target: gl.TEXTURE_2D, Texture: texture.Gl},
	}
	spriteID := render.RegisterSprite(sprite)

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
		world.RegisterTileType(name, tileType, defaltToolForLayer(layerID))

	return tileTypeID
}

func imageForTilesetTile(
	layerTile *tiled.LayerTile, opts TileRegistrationOptions,
) TilesetTileImage {
	if opts.TilesetTile != nil && opts.TilesetTile.Image != nil {
		tileImg := opts.TilesetTile.Image
		imgPath := path.Join(opts.TmxPath, "..", tileImg.Source)
		spriteTransform := mgl32.Ident3()
		model := modelFromSize(tileImg.Width, tileImg.Height)
		return TilesetTileImage {
			Path: imgPath,
			SpriteTransform: spriteTransform,
			Width: tileImg.Width, Height: tileImg.Height,
		}
	} else {
		imgPath := path.Join(opts.TmxPath, "..", opts.Tileset.Image.Source)
		spriteRect := opts.Tileset.GetTileRect(layerTile.ID)
		tileset := opts.Tileset
		tilesetDims := image.Point { tileset.Image.Width, tileset.Image.Height }
		spriteTransform := rectTransform( spriteRect, tilesetDims )
		model := mgl32.Ident4()
		return TilesetTileImage {
			Path: imgPath,
			SpriteTransform: spriteTransform,
			Width: 1, Height: 1,
		}
	}
}


func modelFromSize(dx int, dy int) mgl32.Mat4 {
	return mgl32.Scale2D(float32(dx)/16, float32(dy)/16).Mat4()
}
