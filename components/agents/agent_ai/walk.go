package agent_ai

import "github.com/skycoin/cx-game/agents"

const (
	walkSpeed float32 = 1
	jumpSpeed float32 = 15
)

/*
func shouldJump(agent *agents.Agent) {
	// TODO
	return false
}
*/

func AiHandlerWalk(agent *agents.Agent) {
	agent.PhysicsState.Vel.X = 1
}
