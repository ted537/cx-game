package render

import (
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
	UVTransform mgl32.Mat3
	Options SpriteDrawOptions
}
var spriteDrawsPerAtlas = map[Texture][]SpriteDraw{}

func drawSprite(modelView mgl32.Mat4, id SpriteID, opts SpriteDrawOptions) {
	sprite := sprites[id]
	atlas := sprite.Texture
	spriteDrawsPerAtlas[atlas] = append( spriteDrawsPerAtlas[atlas],
		SpriteDraw {
			Sprite: sprite,
			ModelView: modelView,
		} )
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
	for atlas,spriteDraws := range spriteDrawsPerAtlas {
		drawAtlasSprites(atlas, spriteDraws)
	}
	spriteDrawsPerAtlas = make(map[Texture][]SpriteDraw)
}

func drawAtlasSprites(atlas Texture, spriteDraws []SpriteDraw) {
	atlas.Bind()
	defer atlas.Unbind()

	uniforms := extractUniforms(spriteDraws)
	for _,batch := range uniforms.Batch(constants.DRAW_SPRITE_BATCH_SIZE) {
		drawInstancedQuads(batch)
	}
}

func extractUniforms(spriteDraws []SpriteDraw) Uniforms {
	uniforms := NewUniforms(len(spriteDraws))
	for _,spriteDraw := range spriteDraws {

	}
	return uniforms
}

type Uniforms struct {
	Count int
	ModelViews []mgl32.Mat4
	UVTransforms []mgl32.Mat3
}

func NewUniforms(count int) Uniforms {
	return Uniforms {
		Count: count,
		ModelViews: make([]mgl32.Mat4, count),
		UVTransforms: make([]mgl32.Mat3, count),
	}
}

func (u Uniforms) Batch(batchSize int) []Uniforms {
	numBatches := divideRoundUp(u.Count, batchSize)
	batches := make([]Uniforms,numBatches)

	for i := batchSize ; i < u.Count ; i+= batchSize {
		batches[i] = u.Range(i-batchSize,i)
	}

	return batches
}

func (u Uniforms) Range(start,stop int) Uniforms {
	return Uniforms {
		Count: stop-start+1,
		ModelViews: u.ModelViews[start:stop],
		UVTransforms: u.UVTransforms[start:stop],
	}
}

func divideRoundUp(a,b int) int {
	if a%b == 0 {
		return a/b
	} else {
		return a/b+1
	}
}

func drawInstancedQuads(batch Uniforms) {

}
