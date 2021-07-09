package ui

import (
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"

	"github.com/skycoin/cx-game/spriteloader"
	"github.com/skycoin/cx-game/render"
	"github.com/skycoin/cx-game/utility"
)

func newStretchingNineSliceVao(w,h float32) uint32 {
	geometry := utility.NewGeometry()

	// TODO read from struct / config file
	left := float32(1.0/8.0)
	right := left
	top := float32(1.0/8.0)

	geometry.AddQuadFromCorners(
		utility.Vert { 0,0,0, 0,0 },
		utility.Vert { left,-top,0, left,top },
	)
	geometry.AddQuadFromCorners(
		utility.Vert { left,0,0, left,0 },
		utility.Vert { w-right,-top,0, 1-right,top },
	)

	log.Printf("nineslice geometry has %d verts",geometry.Verts())

	return geometry.Upload()
}

type StretchingNineSlice struct {
	sprite spriteloader.SpriteID
	vao uint32
	shader *utility.Shader
}

func NewStretchingNineSlice(
		sprite spriteloader.SpriteID, w,h float32,
) StretchingNineSlice {
	return StretchingNineSlice {
		sprite: sprite,
		vao: newStretchingNineSliceVao(w,h),
		shader: utility.NewShader(
			"./assets/shader/mvp.vert", "./assets/shader/tex.frag" ),
	}
}

func (nine StretchingNineSlice) Draw(ctx render.Context) {
	metadata := spriteloader.GetSpriteMetadata(nine.sprite)
	gl.ActiveTexture(gl.TEXTURE0)
	nine.shader.Use()
	defer nine.shader.StopUsing()

	nine.shader.SetUint("tex", metadata.GpuTex)
	gl.BindTexture(gl.TEXTURE_2D, metadata.GpuTex)

	nine.shader.SetVec4F("color",1,1,1,1)
	nine.shader.SetVec4F("colour",1,1,1,1)
	mvp := ctx.MVP()
	nine.shader.SetMat4("mvp", &mvp)

	nine.shader.SetVec2F("offset", 0,0)
	nine.shader.SetVec2F("scale", 1,1)

	gl.Disable(gl.DEPTH_TEST)
	gl.BindVertexArray(nine.vao)
	// draw 9 quads = 18 triangles = 54 verts
	gl.DrawArrays(gl.TRIANGLES,0,2*3*2)
}
