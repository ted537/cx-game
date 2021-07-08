package pathfinding

import (
	"github.com/go-gl/mathgl/mgl32"
)

type BehaviourContext struct {
	Self *Enemy
	PlayerPos mgl32.Vec2
}

type Instruction struct {
	ShouldJump bool
	Direction float32
}

type Behaviour interface {
	Follow(BehaviourContext) Instruction
}
type BehaviourID uint32

var behaviours = []Behaviour{}

func (id BehaviourID) Get() Behaviour {
	return behaviours[id]
}

func RegisterBehaviour(behaviour Behaviour) BehaviourID {
	id := BehaviourID(len(behaviours))
	behaviours = append(behaviours,behaviour)
	return id
}
