package ui

import (
	"github.com/skycoin/cx-game/spriteloader"
)

type HealthBar struct {}
func NewHealthBar() HealthBar {
	return HealthBar{}
}

func (bar HealthBar) Draw(x float32) {
	// TODO
}

type CircleIndicator struct {
	spriteID spriteloader.SpriteID
}
func NewCircleIndicator(spriteID spriteloader.SpriteID) CircleIndicator {
	return CircleIndicator { spriteID: spriteID }
}


// x describes how full circle is
func (indicator CircleIndicator) Draw(x float32) {

}

// all values are normalized to [1,1] range
type HUDState struct {
	Health float32

	Fullness float32 // opposite of hunger
	Hydration float32
	Oxygen float32
	Fuel float32
}

type HUD struct {
	Health HealthBar

	Fullness CircleIndicator
	Hydration CircleIndicator
	Oxygen CircleIndicator
	Fuel CircleIndicator
}
var hud HUD

func InitHUD() {
	hud = HUD {
		Health: NewHealthBar(),

		Fullness: NewCircleIndicator(spriteloader.LoadSingleSprite(
			"./assets/hud/hud_status_food.png", "status_food")),
		Hydration: NewCircleIndicator(spriteloader.LoadSingleSprite(
			"./assets/hud/hud_status_water.png", "status_water")),
		Oxygen: NewCircleIndicator(spriteloader.LoadSingleSprite(
			"./assets/hud/hud_status_oxygen.png", "status_oxygen")),
		Fuel: NewCircleIndicator(spriteloader.LoadSingleSprite(
			"./assets/hud/hud_status_fuel.png", "status_fuel")),
	}
}

func DrawHUD(state HUDState) {
	hud.Draw(state)
}

func (h HUD) Draw(state HUDState) {
	h.Health.Draw(state.Health)
	h.Fullness.Draw(state.Fullness)
	h.Hydration.Draw(state.Hydration)
	h.Oxygen.Draw(state.Oxygen)
	h.Fuel.Draw(state.Fuel)
}
