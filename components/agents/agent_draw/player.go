package agent_draw

import (
	"github.com/go-gl/mathgl/mgl32"

	"github.com/skycoin/cx-game/agents"
	"github.com/skycoin/cx-game/spriteloader"
	"github.com/skycoin/cx-game/render"
)

const (
	playerHeadSize float32 = 1.5
)

func drawPlayerSprite(
		agent *agents.Agent, ctx DrawHandlerContext, 
		spriteID spriteloader.SpriteID,
) {
	translate := mgl32.Translate3D(
		agent.PhysicsState.Pos.X-ctx.Camera.X,
		agent.PhysicsState.Pos.Y-ctx.Camera.Y,
		0,
	)
	scaleX :=
		-agent.PhysicsState.Size.X * agent.PhysicsState.Direction

	scale := mgl32.Scale3D( scaleX, agent.PhysicsState.Size.Y, 1)
	renderCtx := render.Context {
		World: translate.Mul4(scale),
		Projection: spriteloader.Window.GetProjectionMatrix(),
	}
	drawOpts := spriteloader.NewDrawOptions()
	spriteloader.DrawSpriteQuadContext(renderCtx, spriteID, drawOpts)
}

func PlayerDrawHandler(agents []*agents.Agent, ctx DrawHandlerContext) {
	if len(agents)==0 { return }
	for _, agent := range agents {

		drawPlayerSprite(agent, ctx, agent.PlayerData.SuitSpriteID)
		drawPlayerSprite(agent, ctx, agent.PlayerData.HelmetSpriteID)
	}
}
