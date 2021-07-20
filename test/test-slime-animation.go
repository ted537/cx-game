package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/skycoin/cx-game/render"
	"github.com/skycoin/cx-game/spriteloader"
)

var spriteAnimated *spriteloader.SpriteAnimated

func init() {
	runtime.LockOSThread()
}

func keyCallBack(w *glfw.Window, k glfw.Key, s int, a glfw.Action, mk glfw.ModifierKey) {
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	win := render.NewWindow(600, 400, true)
	window := win.Window
	window.SetKeyCallback(keyCallBack)
	spriteloader.InitSpriteloader(&win)
	spriteAnimated = spriteloader.
		NewSpriteAnimated("./assets/slime.json")
	action := spriteAnimated.Action("Idle")
	time := glfw.GetTime()
	for !window.ShouldClose() {
		gl.ClearColor(1, 1, 1, 1)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		newTime := glfw.GetTime()
		dt := float32(newTime-time)
		time = newTime
		action.Update(dt)

		log.Printf("slime has sprite id = %v", action.SpriteID())

		spriteloader.DrawSpriteQuad(0,0,3,2,action.SpriteID())

		glfw.PollEvents()
		window.SwapBuffers()
	}

}
