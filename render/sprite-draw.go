package render

import (
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"

	"github.com/skycoin/cx-game/cxmath/math32"
	"github.com/skycoin/cx-game/cxmath/math32i"
	"github.com/skycoin/cx-game/constants"
)

type SpriteID int

// Usage:
// SetWorldWidth(...)
// DrawSprite(...) - many times
// Flush()

var worldWidth float32
var cameraTransform mgl32.Mat4
func SetWorldWidth(w float32) { worldWidth = w }
func SetCameraTransform(mat mgl32.Mat4) { cameraTransform = mat }

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
			UVTransform: sprite.Transform,
		} )
}

// unaffected by camera movement
func DrawUISprite(transform mgl32.Mat4, id SpriteID, opts SpriteDrawOptions) {
	drawSprite(transform, id, opts)
}

// affected by camera movement
// TODO implement wrap-around in here
func DrawWorldSprite(transform mgl32.Mat4, id SpriteID, opts SpriteDrawOptions) {
	position := transform.Col(3)
	positiveXPosition := math32.PositiveModulo(position.X(), worldWidth)
	// TODO
	_ = position; _ = positiveXPosition

	modelview := cameraTransform.Inv().Mul4(transform)
	drawSprite(modelview, id, opts)
}

func Flush(projection mgl32.Mat4) {
	flushSpriteDraws(projection)
}

func flushSpriteDraws(projection mgl32.Mat4) {
	spriteProgram.Use()
	defer spriteProgram.StopUsing()

	gl.BindVertexArray(QuadVao)

	spriteProgram.SetMat4("projection", &projection)

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
	uniforms := NewUniforms(int32(len(spriteDraws)))
	for idx,spriteDraw := range spriteDraws {
		uniforms.ModelViews[idx] = spriteDraw.ModelView
		uniforms.UVTransforms[idx] = spriteDraw.UVTransform
	}
	return uniforms
}

type Uniforms struct {
	Count int32
	ModelViews []mgl32.Mat4
	UVTransforms []mgl32.Mat3
}

func NewUniforms(count int32) Uniforms {
	return Uniforms {
		Count: count,
		ModelViews: make([]mgl32.Mat4, count),
		UVTransforms: make([]mgl32.Mat3, count),
	}
}

func (u Uniforms) Batch(batchSize int32) []Uniforms {
	numBatches := divideRoundUp(u.Count, batchSize)
	batches := make([]Uniforms,numBatches)

	for i := int32(0) ; i < numBatches ; i++ {
		start := batchSize * i
		stop := math32i.Min(batchSize * (i+1), u.Count)
		batches[i] = u.Range(start,stop)
	}

	return batches
}

func (u Uniforms) Range(start,stop int32) Uniforms {
	return Uniforms {
		Count: stop-start,
		ModelViews: u.ModelViews[start:stop],
		UVTransforms: u.UVTransforms[start:stop],
	}
}

func divideRoundUp(a,b int32) int32 {
	if a%b == 0 {
		return a/b
	} else {
		return a/b+1
	}
}

func drawInstancedQuads(batch Uniforms) {
	spriteProgram.SetMat4s("modelviews", batch.ModelViews)
	spriteProgram.SetMat3s("uvtransforms", batch.UVTransforms)
	gl.DrawArraysInstanced(gl.TRIANGLES, 0,6, batch.Count)
}
