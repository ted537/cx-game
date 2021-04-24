package tile

import (
	"log"
	"math"

	"github.com/go-gl/mathgl/mgl32"

	"github.com/skycoin/cx-game/sprite"
)

type Tile struct {
	Name string
	SpriteId int
}

type TileMap struct {
	// store all the tiles with names
	Tiles []Tile
	// layout the stored tiles in some manner
	TileIds []int
	Width, Height int
}

type TilePaleteSelector struct {
	// store tiles for (1) displaying selector and (2) placing tiles
	Tiles []Tile
	Transform mgl32.Mat4
}

func (tilemap *TileMap) Draw() {
	for idx,tileId := range tilemap.TileIds {
		y := float32(idx / tilemap.Width)
		x := float32(idx % tilemap.Width)
		if tileId>=0 {
			spriteId := tilemap.Tiles[tileId].SpriteId
			sprite.DrawSpriteQuad(x,y,1,1,spriteId)
		}
	}
}

const paleteXOffset = 0.0
const paleteYOffset = -3.0
func (selector *TilePaleteSelector) Draw() {
	numTiles := float64(len(selector.Tiles))
	if numTiles>0 {
		width := math.Ceil(math.Sqrt(numTiles))
		scale := float32(1.0/width)
		for idx,tile := range selector.Tiles {
			yLocal := float32(idx/int(width))*scale
			xLocal := float32(idx%int(width))*scale
			localTransform := mgl32.Mat4.Mul4(
				selector.Transform,
				mgl32.Translate3D(xLocal,yLocal,0),
			)
			localPos := localTransform.Col(3)
			sprite.DrawSpriteQuad(
				localPos.X(),localPos.Y(),
				scale,scale,
				tile.SpriteId,
			)
		}
	}
}

func (selector *TilePaleteSelector) ClickHandler(x,y float32, projection mgl32.Mat4) {
	log.Print("tile palete selector is checking collisions for a click at ",x,y)
	homogenousClipCoords := mgl32.Vec4 { x,y,-1.0,1.0}
	cameraCoords := projection.Inv().Mul4x1(homogenousClipCoords)
	rayEye := mgl32.Vec4 { cameraCoords.X(), cameraCoords.Y(), -1.0, 0 }
	worldCoords := rayEye.Normalize().Mul(sprite.SpriteRenderDistance)

	paleteCoords := selector.Transform.Inv().Mul4x1(worldCoords)
	log.Print(paleteCoords)
}
