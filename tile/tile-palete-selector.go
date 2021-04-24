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
	Width int
	SelectedTileIndex int
}

func MakeTilePaleteSelector(tiles []Tile) TilePaleteSelector {
	width := math.Ceil(math.Sqrt(float64(len(tiles))))
	scale := float32(1.0/width)
	return TilePaleteSelector {
		Tiles: tiles,
		Transform: mgl32.Mat4.Mul4(
			mgl32.Translate3D(0.0,-3.0,0.0),
			mgl32.Scale3D(scale,scale,scale),
		),
		Width: int(width),
		SelectedTileIndex: -1,
	}
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

func (selector *TilePaleteSelector) Draw() {
	numTiles := float64(len(selector.Tiles))
	if numTiles>0 {
		for idx,tile := range selector.Tiles {
			yLocal := float32(idx/int(selector.Width))
			xLocal := float32(idx%int(selector.Width))
			localTransform := mgl32.Mat4.Mul4(
				selector.Transform,
				mgl32.Translate3D(xLocal,yLocal,0),
			)
			localPos := localTransform.Col(3)
			scaleX,scaleY,_ := mgl32.Extract3DScale(localTransform)
			sprite.DrawSpriteQuad(
				localPos.X(),localPos.Y(),
				scaleX,scaleY,
				tile.SpriteId,
			)
		}
	}
}

func convertScreenCoordsToWorld(x,y float32, projection mgl32.Mat4) mgl32.Vec4 {
	homogenousClipCoords := mgl32.Vec4 { x,y,-1.0,1.0}
	cameraCoords := projection.Inv().Mul4x1(homogenousClipCoords)
	rayEye := mgl32.Vec4 { cameraCoords.X(), cameraCoords.Y(), -1.0, 0 }
	worldCoords := rayEye.Normalize().Mul(sprite.SpriteRenderDistance)
	worldCoords[3]=1
	return worldCoords
}

func (selector *TilePaleteSelector) ClickHandler(x,y float32, projection mgl32.Mat4) {
	worldCoords := convertScreenCoordsToWorld(x,y,projection)
	paleteCoords := selector.Transform.Inv().Mul4x1(worldCoords).Vec2()
	tileX := int(paleteCoords.X()+0.5)
	tileY := int(paleteCoords.Y()+0.5)
	if tileX>=0 && tileX<selector.Width && tileY>=0 && tileY<selector.Width {
		selector.SelectedTileIndex = tileY*selector.Width + tileX
	}
}

func (tilemap *TileMap) TryPlaceTile(x,y float32, projection mgl32.Mat4, tileId int) {
	// tilemap is drawn directly on the world - no need to convert further
	worldCoords := convertScreenCoordsToWorld(x,y,projection)
	tileX := int(worldCoords.X()+0.5)
	tileY := int(worldCoords.Y()+0.5)
	if tileX>=0 && tileX<tilemap.Width && tileY>=0 && tileY<tilemap.Width {
		log.Print("clicked world tile at ",tileX,tileY)
		tileIdx := tileY*tilemap.Width + tileX
		tilemap.TileIds[tileIdx] = tileId
	}
}
