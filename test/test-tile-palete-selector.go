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

var win render.Window
var tilemap tile.TileMap
var tilePaleteSelector tile.TilePaleteSelector

func tick() {

}

func draw() {
	win.UpdateProjectionMatrix()
	gl.ClearColor(1,1,1,1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	tilemap.Draw()
	if isDebugMode {
		tilePaleteSelector.Draw()
	}
}

func wait() {
	glfw.PollEvents()
	win.Window.SwapBuffers()
}

func load() {
	log.Print("running test")
	log.Print("you should see a tile palete selector")
	win = render.NewWindow(640,480,true)
	window := win.Window
	window.SetKeyCallback(keyCallBack)
	window.SetCursorPosCallback(cursorPosCallback)
	window.SetMouseButtonCallback(mouseButtonCallback)

	sprite.InitSpriteloader(&win)
	sprite.LoadSingleSprite(
		"./assets/tile/test-tile-blue-01.png",
		"blue1",
	)
	sprite.LoadSingleSprite(
		"./assets/tile/test-tile-blue-02.png",
		"blue2",
	)
	sprite.LoadSingleSprite(
		"./assets/tile/test-tile-blue-03.png",
		"blue3",
	)
	sprite.LoadSingleSprite(
		"./assets/tile/test-tile-stone-01.png",
		"stone1",
	)
	sprite.LoadSingleSprite(
		"./assets/tile/test-tile-stone-02.png",
		"stone2",
	)
	sprite.LoadSingleSprite(
		"./assets/tile/test-tile-stone-03.png",
		"stone3",
	)
	sprite.LoadSingleSprite(
		"./assets/tile/test-tile-wood-02.png",
		"wood2",
	)
	sprite.LoadSingleSprite(
		"./assets/tile/test-tile-wood-03.png",
		"wood3",
	)

	tiles := []tile.Tile {tile.Tile{
		Name: "real tile",
		SpriteId: sprite.GetSpriteIdByName("wood3"),
	}}
	tilemap = tile.TileMap {
		Tiles: tiles,
		TileIds: []int{-1,0,0,-1},
		Width: 2, Height: 2,
	}
	tilePaleteSelector = tile.TilePaleteSelector {
		Tiles: tiles,
	}
}

func main() {
	load()
	defer glfw.Terminate()
	for !win.Window.ShouldClose() {
		tick()
		draw()
		wait()
	}
}
