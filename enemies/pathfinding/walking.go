package pathfinding

import (
	"github.com/go-gl/mathgl/mgl32"

	"github.com/skycoin/cx-game/cxmath/math32"
	"github.com/skycoin/cx-game/physics"
)

type WalkingBehaviour struct {
	walkSpeed float32
	jumpSpeed float32
}

func (wb WalkingBehaviour) shouldJump(ctx BehaviourContext) bool {
	needsToJumpUpLeftWall := ctx.Self.Collisions.Left && ctx.Self.Vel.X<0
	needsToJumpUpRightWall := ctx.Self.Collisions.Right && ctx.Self.Vel.X>0
	return needsToJumpUpLeftWall || needsToJumpUpRightWall
}

func (wb WalkingBehaviour) Follow(ctx BehaviourContext) Instruction {
	// TODO fix this
	dt := float32(1.0/30.0)
	velX := float32(0); velY := float32(0)
	directionX := math32.Sign(ctx.PlayerPos.X()-ctx.Self.Pos.X)
	velX = directionX * wb.walkSpeed

	if wb.shouldJump(ctx) {
		velY = wb.jumpSpeed
	} else {
		velY = physics.Gravity * dt
	}

	return Instruction { Velocity: mgl32.Vec2{ velX, velY } }
}

var WalkingBehaviourID BehaviourID
func init() {
	WalkingBehaviourID = RegisterBehaviour(WalkingBehaviour{})
}
