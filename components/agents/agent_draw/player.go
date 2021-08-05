package agent_draw

import (
	"github.com/skycoin/cx-game/agents"
	"github.com/skycoin/cx-game/spriteloader"
)

const (
	playerHeadSize float32 = 1.5
)

func PlayerDrawHandler(agents []*agents.Agent, ctx DrawHandlerContext) {
	// TODO is this assumed??? can we omit this check?
	if len(agents)==0 { return }
	drawOpts := spriteloader.NewDrawOptions()
	for _, agent := range agents {
		//drawOpts.Alpha = alphaForAgent(agent)
		spriteloader.DrawSpriteQuadOptions(
			agent.PhysicsState.Pos.X-ctx.Camera.X,
			agent.PhysicsState.Pos.Y-ctx.Camera.Y,
			agent.PhysicsState.Size.X, agent.PhysicsState.Size.Y,
			agent.PlayerData.SuitSpriteID, drawOpts,
		)
		spriteloader.DrawSpriteQuadOptions(
			agent.PhysicsState.Pos.X-ctx.Camera.X,
			agent.PhysicsState.Pos.Y-ctx.Camera.Y,
			agent.PhysicsState.Size.X, agent.PhysicsState.Size.Y,
			agent.PlayerData.HelmetSpriteID, drawOpts,
		)
	}
}
