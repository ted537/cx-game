package render

import (
	"log"

	"github.com/go-gl/mathgl/mgl32"

	"github.com/skycoin/cx-game/cxmath/math32"
	"github.com/skycoin/cx-game/constants"
)

type SpriteID int

// Usage:
// SetWorldWidth(...)
// DrawSprite(...) - many times
// Flush()

var worldWidth float32
func SetWorldWidth(w float32) { worldWidth = w }

type SpriteDrawOptions struct {

}
func NewSpriteDrawOptions() SpriteDrawOptions{
	return SpriteDrawOptions{}
}

type SpriteDraw struct {
	Sprite Sprite
	ModelView mgl32.Mat4
	Options SpriteDrawOptions
}
var spriteDraws = [][]SpriteDraw{}

func drawSprite(modelView mgl32.Mat4, id SpriteID, opts SpriteDrawOptions) {
	sprite := sprites[id]
	atlas := sprite.Texture
	spriteDraws[atlas] = append( spriteDraws[atlas], SpriteDraw {
		Sprite: sprite,
		ModelView: modelView,
	})
}

// unaffected by camera movement
func DrawUISprite(transform mgl32.Mat4, id SpriteID) {

}

// affected by camera movement
// TODO implement wrap-around in here
func DrawWorldSprite(transform mgl32.Mat4, id SpriteID) {
	position := transform.Col(3)
	positiveXPosition := math32.PositiveModulo(position.X(), worldWidth)
}

func Flush() {
	flushSpriteDraws()
}

func flushSpriteDraws() {
	for atlasIndex := range spriteDraws {
		drawAtlasSprites(atlasIndex)
	}
	spriteDraws = make([][]SpriteDraw, len(atlases))
}

func drawAtlasSprites(atlasIndex int) {
	atlases[atlasIndex].Bind()
	defer atlases[atlasIndex].Unbind()

	uniforms := batchUniforms(spriteDraws[atlasIndex])
	for _,batch := range uniforms {
		drawInstancedQuads(batch)
	}
}

type UniformBatch struct {
	ModelViews [constants.DRAW_SPRITE_BATCH_SIZE]mgl32.Mat4
	ModelTransforms [constants.DRAW_SPRITE_BATCH_SIZE]mgl32.Mat3
}
func batchUniforms(spriteDraws []SpriteDraw) []UniformBatch {
	log.Fatal("batchUniforms() has not been implemented")
	return nil
}

func drawInstancedQuads(batch UniformBatch) {

}
