package pathfinding

import (
	"github.com/skycoin/cx-game/cxmath/math32"
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
	directionX := math32.Sign(ctx.PlayerPos.X()-ctx.SelfPos.X())
	ctx.Enemy.Vel.X = directionX * wb.walkSpeed

	if wb.shouldJump(ctx) {
		ctx.Enemy.Vel.Y = wb.jumpSpeed
	} else {
		ctx.Enemy.Vel.Y = physics.Gravity * dt
	}
}

var WalkingBehaviourID BehaviourID
func init() {
	WalkingBehaviourID = RegisterBehaviour(WalkingBehaviour{})
}
