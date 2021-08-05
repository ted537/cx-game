package agent_ai

import (
	"github.com/skycoin/cx-game/agents"
	"github.com/skycoin/cx-game/input"
	"github.com/skycoin/cx-game/cxmath"
	"github.com/skycoin/cx-game/cxmath/math32"
)

const playerWalkSpeed float32 = 5
const frictionFactor float32 = 3

func AiHandlerPlayer(player *agents.Agent, ctx AiContext) {
	// TODO
	inputXAxis := input.GetAxis(input.HORIZONTAL)
	player.PhysicsState.Vel.X +=
		inputXAxis * playerWalkSpeed

	friction :=
		cxmath.Sign(player.PhysicsState.Vel.X) * frictionFactor

	if math32.Abs(friction) < math32.Abs(player.PhysicsState.Vel.X) {
		player.PhysicsState.Vel.X -= friction
	}
}
