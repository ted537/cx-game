package agent_ai

import (
	"log"
	"github.com/skycoin/cx-game/agents"
	"github.com/skycoin/cx-game/constants"
	"github.com/skycoin/cx-game/components/types"
)

type AiHandler func(*agents.Agent)
var aiHandlers [constants.NUM_AI_HANDLERS]AiHandler

func Init() {
	RegisterAiHandler(constants.AI_HANDLER_NULL,AiHandlerNull)
}

func assertAllAiHandlersRegistered() {
	for _, handler := range aiHandlers {
		if handler == nil {
			log.Fatalf("Did not initialize all agent ai handlers")
		}
	}
}

func RegisterAiHandler(id types.AgentAiHandlerID, newHandler AiHandler) {
	aiHandlers[id] = newHandler
}
