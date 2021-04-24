package main

import (
	"log"
	"runtime"

	"github.com/skycoin/cx-game/sprite"
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
var mouseX = 0.0
var mouseY = 0.0

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

func mouseButtonCallback(w *glfw.Window, b glfw.MouseButton, a glfw.Action, mk glfw.ModifierKey) {
	log.Print("clicked at ",mouseX,mouseY)
}

func cursorPosCallback(w *glfw.Window, xpos, ypos float64) {
	mouseX = xpos
	mouseY = ypos
}

func main() {
	log.Print("running test")
	log.Print("you should see a tile palete selector")
	win := render.NewWindow(640,480,true)
	window := win.Window
	window.SetKeyCallback(keyCallBack)
	window.SetCursorPosCallback(cursorPosCallback)
	window.SetMouseButtonCallback(mouseButtonCallback)
	defer glfw.Terminate()

	sprite.InitSpriteloader(&win)
	spriteSheetId := sprite.
		LoadSpriteSheet("./assets/starfield/stars/planets.png")
	sprite.
		LoadSprite(spriteSheetId, "star", 2,1)
	spriteId := sprite.
		GetSpriteIdByName("star")

	tiles := []tile.Tile {tile.Tile{
		Name: "real tile",
		SpriteId: spriteId,
	}}
	tilemap := tile.TileMap {
		Tiles: tiles,
		TileIds: []int{-1,0,0,-1},
		Width: 2, Height: 2,
	}
	tilePaleteSelector := tile.TilePaleteSelector {
		Tiles: tiles,
	}

	for !window.ShouldClose() {
		gl.ClearColor(1,1,1,1)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		tilemap.Draw()
		if isDebugMode {
			tilePaleteSelector.Draw()
		}
		glfw.PollEvents()
		window.SwapBuffers()
	}
	tilemap.Draw()
}
