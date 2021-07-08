package ui

import (
	"github.com/skycoin/cx-game/render"
)

// all values are normalized to [1,1] range
type HUDState struct {
	Health float32

	Fullness float32 // opposite of hunger
	Hydration float32
	Oxygen float32
	Fuel float32
}

type HUD struct {
	fullnessSpriteID uin
}

func (h HUD) drawHealthBar(health float32) {
	// TODO
}

func (h HUD) drawFullnessCircle(fullness float32) {

}

func (h HUD)

func (h HUD) Draw() {
	h.drawHealthBar()
	h.drawFullnessCircle()
	h.drawHydrationCircle()
	h.drawOxygenCircle()
	h.drawFuelCircle()
}
