package ui

import (
	"github.com/skycoin/cx-game/render"
)

// all values are normalized to [1,1] range
type HUD struct {
	Health float32

	Fullness float32 // opposite of hunger
	Hydration float32
	Oxygen float32
	Fuel float32
}

func (h HUD) Draw(ctx render.Context) {
	h.drawHealthBar(ctx)
	h.drawFullnessCircle(ctx)
	h.drawHydrationCircle(ctx)
	h.drawOxygenCircle(ctx)
	h.drawFuelCircle(ctx)
}

func DrawHUD(ctx render.Context, state HUDState) {
	
}
