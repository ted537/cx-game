package agents

import (
	"github.com/skycoin/cx-game/constants"
)

type AgentList struct {
	// profile this to see if reducing indirection
	// would help with performance in a significant way
	Agents []*Agent
}

func NewAgentList() *AgentList {
	return &AgentList{
		Agents: make([]*Agent, 0),
	}
}

func NewDevAgentList() *AgentList {
	agentList := NewAgentList()
	player := newAgent(len(agentList.Agents))
	player.AgentCategory = constants.AGENT_CATEGORY_PLAYER
	agentList.CreateAgent(player)
	enemy := newAgent(len(agentList.Agents))
	enemy.AgentCategory = constants.AGENT_CATEGORY_ENEMY_MOB
	agentList.CreateAgent(enemy)

	return agentList
}

func (al *AgentList) CreateAgent(agent *Agent) bool {
	//for now
	if len(al.Agents) > constants.MAX_AGENTS {
		return false
	}
	al.Agents = append(al.Agents, agent)
	return true
}

func (al *AgentList) DestroyAgent(agentId int) bool {
	if agentId < 0 || agentId >= len(al.Agents) {
		return false
	}

	al.Agents = append(al.Agents[:agentId], al.Agents[agentId+1:]...)
	return false
}

func (al *AgentList) Spawn(
		agentTypeID constants.AgentTypeID, opts AgentCreationOptions,
) int {
	agent := GetAgentType(agentTypeID).CreateAgent(opts)
	agent.FillDefaults()
	agent.Validate()
	agent.AgentId = len(al.Agents)
	al.Agents = append(al.Agents, agent)
	return agent.AgentId
}

func (al *AgentList) Get() []*Agent { return al.Agents }

func (al *AgentList) FromID(id int) *Agent { return al.Get()[id] }

func (al *AgentList) TileIsClear(x,y int) bool {
	for _,agent := range al.Get() {
		if agent.PhysicsState.Contains(float32(x), float32(y)) { return false }
	}
	return true
}
