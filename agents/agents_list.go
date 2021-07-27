package agents

import "github.com/skycoin/cx-game/constants"

type AgentList struct {
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

//  agentType - constants.AGENT_CATEGORY_*desired type*
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
