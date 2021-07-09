package ui

import (
	"math"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"

	"github.com/skycoin/cx-game/utility"
	"github.com/skycoin/cx-game/cxmath/math32"
)

const arcTriangles = 50
var arcVAO uint32
var arcVBO uint32
var arcShader *utility.Shader

func createArcVertexAttributes() []float32 {
	attributes := make([]float32,arcTriangles*3*5)
	i := 0
	for tri := 0 ; tri < arcTriangles; tri++ {
		angle := 2 * math.Pi * float32(tri) / float32(arcTriangles)
		x := math32.Sin(angle)
		y := math32.Cos(angle)
		z := float32(0)
		attributes[i] = x
		attributes[i+1] = y
		attributes[i+2] = z
		// arc is currently untextured so values of u and v are not important
		i += 5
	}
	return attributes
}

func initArcVAO() {
	var arcVBO uint32
	gl.GenBuffers(1,&arcVBO)
	gl.GenVertexArrays(1,&arcVAO)
	gl.BindVertexArray(arcVAO)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER,arcVBO)

	vertexAttributes := createArcVertexAttributes()

	gl.BufferData(
		gl.ARRAY_BUFFER,
		4*len(vertexAttributes),
		gl.Ptr(vertexAttributes),
		gl.STATIC_DRAW,
	)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(4*3))
	gl.EnableVertexAttribArray(1)
	//unbind
	gl.BindVertexArray(0)
}

func InitArc() {
	initArcVAO()
	arcShader = utility.NewShader(
		"./assets/shader/mvp.vert", "./assets/shader/color.frag" )
}

func DrawArc(mvp mgl32.Mat4, color mgl32.Vec4, fullness float32) {
	arcShader.Use()
	defer arcShader.StopUsing()
	arcShader.SetMat4("mvp",&mvp)
	arcShader.SetVec4("color",&color)
	gl.Disable(gl.DEPTH_TEST)
	gl.BindVertexArray(arcVAO)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(fullness*arcTriangles*3))
}
