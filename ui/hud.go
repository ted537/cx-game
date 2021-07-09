package ui

import (
	"github.com/go-gl/mathgl/mgl32"

	"github.com/skycoin/cx-game/render"
	"github.com/skycoin/cx-game/spriteloader"
	"github.com/skycoin/cx-game/utility"
)

type HealthBar struct {}
func NewHealthBar() HealthBar {
	return HealthBar{}
}

func (bar HealthBar) Draw(ctx render.Context,x float32) {
	// TODO
	utility.DrawColorQuad(ctx, mgl32.Vec4{1,0,0,1})
}

type CircleIndicator struct {
	spriteID spriteloader.SpriteID
}
func NewCircleIndicator(spriteID spriteloader.SpriteID) CircleIndicator {
	return CircleIndicator { spriteID: spriteID }
}


// x describes how full circle is
func (indicator CircleIndicator) Draw(ctx render.Context,x float32) {
	DrawArc(ctx.MVP(), mgl32.Vec4{1,1,1,1}, x)
	spriteloader.DrawSpriteQuadContext(ctx, indicator.spriteID)
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

	hpIconSpriteID spriteloader.SpriteID
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

		hpIconSpriteID: spriteloader.LoadSingleSprite(
			"./assets/hud/hud_hp_icon.png", "hp_icon" ),
	}
}

func DrawHUD(state HUDState) {
	hud.Draw(state)
}

const hudPadding = 1
const circleYOffset = float32(-1)
func (h HUD) Draw(state HUDState) {
	y := circleYOffset
	ctx := render.CenterToTopLeft(spriteloader.Window.DefaultRenderContext()).
		// padding
		PushLocal(mgl32.Translate3D(hudPadding,-hudPadding,0))

	spriteloader.DrawSpriteQuadContext(ctx, h.hpIconSpriteID)

	h.Health.Draw(
		ctx.PushLocal(mgl32.Translate3D(1,0,0)),state.Health)
	// TODO offset these
	h.Fullness.Draw(
		ctx.PushLocal(mgl32.Translate3D(0,y,0)),
		state.Fullness,
	)
	h.Hydration.Draw(
		ctx.PushLocal(mgl32.Translate3D(1,y,0)),state.Hydration)
	h.Oxygen.Draw(
		ctx.PushLocal(mgl32.Translate3D(2,y,0)),state.Oxygen)
	h.Fuel.Draw(
		ctx.PushLocal(mgl32.Translate3D(3,y,0)),state.Fuel)
}
