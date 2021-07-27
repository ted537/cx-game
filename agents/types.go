package agents

import (
	"log"
	"github.com/skycoin/cx-game/constants"
)

type AgentCreationOptions struct {
	X,Y float32
}

type AgentType interface {
	CreateAgent(AgentCreationOptions) Agent
}

var agentTypes [constants.NUM_AGENT_TYPES]AgentType

func init() {
	defer assertAllAgentTypesRegistered()

}

func RegisterAgentType(id constants.AgentTypeID, agentType AgentType) {
	agentTypes[id] = agentType
}

func registerSlime() {
	
}

func assertAllAgentTypesRegistered() {
	for id,agentType := range agentTypes {
		if agentType == nil {
			log.Fatalf("did not register agent type for id=%d",id)
		}
	}
}
