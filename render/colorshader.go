// super simple shader for drawing colors directly.
// intended for UI.
package render;

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"

	"github.com/skycoin/cx-game/constants"
	"github.com/skycoin/cx-game/cxmath/math32i"
)

func NewColorShader() Program {
	return CompileProgram(
		"./assets/shader/color.vert",
		"./assets/shader/color.frag",
	)
}
var colorProgram Program
var colorProgramInit bool = false

type ColorUniforms struct {
	Count int32
	ModelViews []mgl32.Mat4
	Colors []mgl32.Vec4
}
var colorUniforms ColorUniforms

func (c *ColorUniforms) Add(modelView mgl32.Mat4, color mgl32.Vec4) {
	c.Count++
	c.ModelViews = append( c.ModelViews, modelView )
	c.Colors = append( c.Colors, color )
}

func (c *ColorUniforms) Clear() {
	c.Count = 0
	c.ModelViews = c.ModelViews[:0]
	c.Colors = c.Colors[:0]
}

// n is batch size
func (c ColorUniforms) Batch(n int32) []ColorUniforms {
	batchCount := divideRoundUp(c.Count,n)
	batches := make([]ColorUniforms, batchCount)
	for i := int32(0) ; i < n ; i++ {
		start := n * i
		stop := math32i.Min(n * (i+1), c.Count)
		batches[i] = c.Range(start,stop)
	}
	return batches
}

func (c ColorUniforms) Range(start,stop int32) ColorUniforms {
	return ColorUniforms {
		Count: stop-start,
		ModelViews: c.ModelViews[start:stop],
		Colors: c.Colors[start:stop],
	}
}

func (c ColorUniforms) Set(program Program) {
	program.SetMat4s("modelviews", c.ModelViews)
	program.SetVec4s("colors", c.Colors)
}

func DrawColorQuad(modelView mgl32.Mat4, color mgl32.Vec4) {
	colorUniforms.Add(modelView,color)
	/*
	if !colorProgramInit {
		colorProgram = NewColorShader()
		colorProgramInit = true
	}
	// setup features
	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	// update uniforms
	colorProgram.Use()
	defer colorProgram.StopUsing()
	colorProgram.SetVec4("colour",&color)
	colorProgram.SetMat4("mvp",&mvp)
	// draw
	gl.BindVertexArray(QuadVao)
	gl.DrawArrays(gl.TRIANGLES,0,6)
	*/
}

func flushColorDraws(projection mgl32.Mat4) {
	colorProgram.Use()
	defer colorProgram.StopUsing()

	colorProgram.SetMat4("projection", &projection)

	batchSize := constants.DRAW_COLOR_BATCH_SIZE
	for _,batch := range colorUniforms.Batch(batchSize) {
		batch.Set(colorProgram)
		gl.DrawArraysInstanced(gl.TRIANGLES, 0, 6, batchSize)
	}
	colorUniforms.Clear()
}
