package game

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"

	"github.com/skycoin/cx-game/components"
	"github.com/skycoin/cx-game/engine/ui"
	"github.com/skycoin/cx-game/item"
	"github.com/skycoin/cx-game/particles"
	"github.com/skycoin/cx-game/render"
	"github.com/skycoin/cx-game/render/worldctx"
	"github.com/skycoin/cx-game/stars/starfield"
	"github.com/skycoin/cx-game/world"
)

func Draw() {
	gl.ClearColor(7.0/255.0, 8.0/255.0, 25.0/255.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	render.SetCameraTransform(Cam.GetTransform())
	render.SetWorldWidth(float32(World.Planet.Width))

	baseCtx := win.DefaultRenderContext()
	baseCtx.Projection = Cam.GetProjectionMatrix()
	// camCtx := baseCtx.PushView(Cam.GetView())
	worldCtx := worldctx.NewWorldRenderContext(Cam, &World.Planet)

	//draw starfield
	starfield.DrawStarField()

	//queue-draw world
	World.Planet.Draw(Cam, world.BgLayer)
	World.Planet.Draw(Cam, world.PipeLayer)
	World.Planet.Draw(Cam, world.MidLayer)
	// draw lasers between mid and top layers.

	// particles.DrawTopParticles(camCtx)
	World.Planet.Draw(Cam, world.TopLayer)

	item.DrawWorldItems(Cam)

	topLeftCtx :=
		render.CenterToTopLeft(win.DefaultRenderContext())
	ui.DrawString(
		fmt.Sprint(fps.CurFps),
		mgl32.Vec4{1, 0.2, 0.3, 1},
		ui.AlignCenter,
		topLeftCtx.PushLocal(mgl32.Translate3D(1, -5, 0)),
	)

	// FIXME: draw dialogue boxes uses alternate projection matrix;
	// restore original projection matrix

	//fix dialogboxdraw
	// ui.DrawDialogueBoxes(camCtx)

	inventory := item.GetInventoryById(player.InventoryID)
	inventory.Draw(baseCtx, Cam.GetTransform().Inv())

	Console.Draw(win.DefaultRenderContext())

	components.Draw_Queued(&World.Entities, Cam)
	render.Flush(Cam.GetProjectionMatrix())

	//draw after flushing
	components.Draw(&World.Entities, Cam)
	ui.DrawAgentHUD(player)
	particles.DrawMidTopParticles(worldCtx)

	glfw.PollEvents()
	win.Window.SwapBuffers()
}
