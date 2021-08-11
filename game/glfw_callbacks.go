package game

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/skycoin/cx-game/cxmath"
	"github.com/skycoin/cx-game/engine/input"
	"github.com/skycoin/cx-game/item"
)

func mouseButtonCallback(
	w *glfw.Window, b glfw.MouseButton, a glfw.Action, mk glfw.ModifierKey,
) {
	if a == glfw.Press {
		mousePressCallback(w, b, a, mk)
	}
	if a == glfw.Release {
		mouseReleaseCallback(w, b, a, mk)
	}
}

// mouse position relative to screen
func screenPos() (float32, float32) {
	w := float32(win.Width)
	h := float32(win.Height)

	// adjust mouse position with zoom
	screenX :=
		((input.GetMouseX()-float32(widthOffset))/scale - w/2) / Cam.Zoom
	screenY :=
		-((input.GetMouseY()-float32(heightOffset))/scale - h/2) / Cam.Zoom
	return screenX, screenY
}

func mouseReleaseCallback(
	w *glfw.Window, b glfw.MouseButton, a glfw.Action, mk glfw.ModifierKey,
) {
	screenX, screenY := screenPos()

	inventory := item.GetInventoryById(player.InventoryID)
	player := findPlayer()
	inventory.OnReleaseMouse(screenX, screenY, Cam, &World.Planet, player)
}

func mousePressCallback(
	w *glfw.Window, b glfw.MouseButton, a glfw.Action, mk glfw.ModifierKey,
) {
	// we only care about mousedown events for now
	if a != glfw.Press {
		return
	}

	screenX, screenY := screenPos()

	inventory := item.GetInventoryById(player.InventoryID)
	clickedSlot :=
		inventory.TryClickSlot(screenX, screenY, Cam, &World.Planet, player)
	if clickedSlot {
		return
	}

	player := World.Entities.Agents.FromID(playerAgentID)
	item.GetInventoryById(player.InventoryID).
		TryUseItem(screenX, screenY, Cam, &World, player)
}

var (
	// what actually are these???
	widthOffset, heightOffset int32 
	scale                     float32 = 1
)

func windowSizeCallback(window *glfw.Window, width, height int) {

	scaleToFitWidth := float32(width) / float32(win.Width)
	scaleToFitHeight := float32(height) / float32(win.Height)
	scale = cxmath.Min(scaleToFitHeight, scaleToFitWidth)

	widthOffset = int32((float32(width) - float32(win.Width)*scale) / 2)
	heightOffset = int32((float32(height) - float32(win.Height)*scale) / 2)
	//correct mouse offsets
	input.UpdateMouseCoords(widthOffset, heightOffset, scale)

	viewportWidth := int32(float32(win.Width)*scale)
	viewportHeight := int32(float32(win.Height)*scale)
	gl.Viewport( widthOffset, heightOffset, viewportWidth, viewportHeight )
}

func scrollCallback(w *glfw.Window, xOff, yOff float64) {
	Cam.SetCameraZoomPosition(float32(yOff))
	input.SetCamZoom(Cam.Zoom)
}
