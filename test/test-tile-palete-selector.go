package main

import (
	"log"
	"runtime"

	//"github.com/skycoin/cx-game/spriteloader"
	"github.com/skycoin/cx-game/render"
	"github.com/skycoin/cx-game/tile"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called 
	// from the main thread.
	runtime.LockOSThread()
}

var isFreeCamera = false
var isDebugMode = false

func keyCallBack(w *glfw.Window, k glfw.Key, s int, a glfw.Action, mk glfw.ModifierKey) {
	if a == glfw.Press {
		if k == glfw.KeyEscape {
			w.SetShouldClose(true)
		}
		if k == glfw.KeyF2 {
			isFreeCamera = !isFreeCamera
			log.Print("free cam is now",isFreeCamera)
		}
		if k == glfw.KeyF3 {
			isDebugMode = !isDebugMode
			log.Print("debug mode is now",isDebugMode)
		}
	}
}

func main() {
	log.Print("running test")
	log.Print("you should see a tile palete selector")
	win := render.NewWindow(640,480,true)
	window := win.Window
	window.SetKeyCallback(keyCallBack)
	defer glfw.Terminate()
	tilemap := tile.TileMap {
		TileIds: []int{1,0,3,4},
		Width: 2, Height: 2,
	}
	for !window.ShouldClose() {
		gl.ClearColor(1,1,1,1)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		tilemap.Draw()
		glfw.PollEvents()
		window.SwapBuffers()
	}
	tilemap.Draw()
}
