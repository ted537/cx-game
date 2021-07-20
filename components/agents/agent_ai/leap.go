package agent_ai

import (
	"github.com/skycoin/cx-game/agents"
	"github.com/skycoin/cx-game/cxmath/math32"
)

const (
	verticalJumpSpeed float32 = 15
	horizontalJumpSpeed float32 = 5
)

func AiHandlerLeap(agent *agents.Agent, ctx AiContext) {
	directionX := math32.Sign( ctx.PlayerPos.X() - agent.PhysicsState.Pos.X )

	onGround := agent.PhysicsState.Collisions.Below
	// TODO add a timer to delay leaps
	canJump := onGround
	if canJump {
		agent.PhysicsState.Vel.X = directionX * horizontalJumpSpeed
		agent.PhysicsState.Vel.Y = verticalJumpSpeed
	}
}
