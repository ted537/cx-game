package ui;

import (
	"log"
//	"math"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/skycoin/cx-game/cxmath"
	"github.com/skycoin/cx-game/spriteloader"
)

var asciiToCharDataMap = make(map[int]CharData)
var fontTex uint32
func InitTextRendering() {
	_,img,_ := spriteloader.LoadPng("assets/font/8bitoperator_jve.png")
	fontTex = spriteloader.MakeTexture(img)

	for _,charData := range charDatas {
		asciiToCharDataMap[charData.ascii] = charData
	}
}

func calculateLineWidth(text string) float32 {
	x := 0
	for _,charCode := range text {
		_ = charCode
		x += 1
		//x += asciiToCharDataMap[int(charCode)].w
	}
	return float32(x)
}
// TODO line wrapping
// TODO alignment options
func DrawString(text string, transform mgl32.Mat4) {
	// setup GPU params
	gl.Disable(gl.DEPTH_TEST)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, fontTex)
	// use the default program with support 
	// for scaled / offset texture lookups
	// TODO dedicate a program just for this
	program := spriteloader.Window.Program
	gl.UseProgram(program)
	gl.Uniform1ui(
		gl.GetUniformLocation(program, gl.Str("ourTexture\x00")),
		fontTex,
	)

	pos := mgl32.Vec2 {
		-calculateLineWidth(text)/2,
		0,
	}
	for _, charCode := range text {
		charData,ok := asciiToCharDataMap[int(charCode)]
		if ok {
			z := -spriteloader.SpriteRenderDistance
			letterTransform := transform.
				Mul4(cxmath.Scale(1)).
				Mul4(mgl32.Translate3D(pos.X(),pos.Y(),z))
			_ = letterTransform
			_ = charData

			texScale := float32(charData.h) / 256
			//texScale := float32(1/math.Max(float64(charData.w),float64(charData.h)))
			gl.Uniform2f(
				gl.GetUniformLocation(program, gl.Str("texScale\x00")),
				texScale, texScale,
			)
			// FIXME
			gl.Uniform2f(
				gl.GetUniformLocation(program, gl.Str("texOffset\x00")),
				float32(charData.tx)/256/texScale, float32(charData.ty)/256/texScale,
			)
			log.Printf("drawing character %v with u %v",string(charData.ascii), float32(charData.tx)/256)

			gl.UniformMatrix4fv(
				gl.GetUniformLocation(program, gl.Str("world\x00")),
				1, false, &letterTransform[0],
			)
			aspect := float32(spriteloader.Window.Width) / float32(spriteloader.Window.Height)
			projectTransform := mgl32.Perspective(
				mgl32.DegToRad(45), aspect, 0.1, 100.0,
			)
			gl.UniformMatrix4fv(
				gl.GetUniformLocation(program, gl.Str("projection\x00")),
				1, false, &projectTransform[0],
			)
			gl.BindVertexArray(spriteloader.QuadVao)
			gl.DrawArrays(gl.TRIANGLES, 0, 6)
			// restore texScale and texOffset to defaults
			// TODO separate GPU programs such that this becomes unecessary
			gl.Uniform2f(
				gl.GetUniformLocation(program, gl.Str("texScale\x00")),
				1, 1,
			)
			gl.Uniform2f(
				gl.GetUniformLocation(program, gl.Str("texOffset\x00")),
				0, 0,
			)
		}
		// TODO variable width fonts
		pos = pos.Add(mgl32.Vec2{1.1,0})
		log.Print(pos)
	}
}
