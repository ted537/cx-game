package enemies

import (
	"github.com/skycoin/cx-game/agents"
	"github.com/skycoin/cx-game/camera"
	"github.com/skycoin/cx-game/constants"
	"github.com/skycoin/cx-game/cxmath"
	"github.com/skycoin/cx-game/models"
	"github.com/skycoin/cx-game/physics"
	"github.com/skycoin/cx-game/spriteloader"
	"github.com/skycoin/cx-game/world"
	"github.com/skycoin/cx-game/enemies/pathfinding"
)

const TimeBeforeFadeout = float32(1.0) // in seconds
const TimeDuringFadeout = float32(1.0) // in seconds

type Enemy struct {
	physics.Body
	SpriteID uint32
	Health int
	TimeSinceLastJump float32
	PathfindingBehaviourID pathfinding.BehaviourID
	TimeSinceDeath float32 // 
}

func (enemy *Enemy) Alpha() float32 {
	if enemy.TimeSinceDeath < TimeBeforeFadeout { return 1 }
	x := enemy.TimeSinceDeath - TimeBeforeFadeout
	return 1 - x / TimeDuringFadeout
}

// the enemy is either alive or has recently died and is fading out
func (enemy *Enemy) Exists() bool {
	return enemy.Alpha() > 0
}

func (enemy *Enemy) IsAlive() bool {
	return enemy.Health > 0
}

func InitBasicEnemies() {
	basicEnemySpriteId = spriteloader.
		LoadSingleSprite("./assets/enemies/basic-enemy.png", "basic-enemy")
}

// TODO load an actual sprite here
var basicEnemySpriteId spriteloader.SpriteID
var basicEnemyMovSpeed = float32(1)
var basicEnemies = []*Enemy{}

const enemyJumpVelocity = 15

// TODO create a system to handle projectiles, melee attacks, etc
var playerStrikeRange = float32(1)

func TickBasicEnemies(
	world *world.Planet, dt float32,
	player *models.Player, playerIsAttacking bool,
) {
	nextEnemies := []*Enemy{}
	for idx, _ := range basicEnemies {
		enemy := basicEnemies[idx]
		enemy.Tick(world, dt, player, playerIsAttacking)
		if enemy.Exists() {
			nextEnemies = append(nextEnemies, enemy)
		} else {
			enemy.Deleted = true
		}
	}
	basicEnemies = nextEnemies
}

func DrawBasicEnemies(cam *camera.Camera) {
	for _, enemy := range basicEnemies {
		if cam.IsInBounds(int(enemy.Pos.X), int(enemy.Pos.Y)) {
			enemy.Draw(cam)
		}
	}
}

// TODO either create enemy as a side effect or return instance
func NewBasicEnemy(x, y float32) *agents.Agent {
	agent := agents.Agent {
		AgentType: constants.AGENT_ENEMY_MOB,
		AiHandlerID: constants.AI_HANDLER_WALK,
		DrawHandlerID: constants.DRAW_HANDLER_QUAD,
		PhysicsState: physics.Body{
			Size: cxmath.Vec2{X: 3.0, Y: 3.0},
			Pos:  cxmath.Vec2{X: x, Y: y},
		},
		HealthComponent: agents.NewHealthComponent(5),
	}
	physics.RegisterBody(&agent.PhysicsState)
	return &agent
}

func NewLeapingEnemy(x,y float32) *agents.Agent {
	agent := agents.Agent {
		AgentType: constants.AGENT_ENEMY_MOB,
		AiHandlerID: constants.AI_HANDLER_LEAP,
		DrawHandlerID: constants.DRAW_HANDLER_QUAD,
		PhysicsState: physics.Body{
			Size: cxmath.Vec2{X:2.0, Y: 2.0},
			Pos:  cxmath.Vec2{X: x, Y: y},
		},
		HealthComponent: agents.NewHealthComponent(5),
	}
	physics.RegisterBody(&agent.PhysicsState)
	return &agent
}

func (enemy Enemy) isStuck() bool {
	return (enemy.Collisions.Left || enemy.Collisions.Right) &&
		!enemy.Collisions.Below
}

func (enemy *Enemy) Tick(
	world *world.Planet, dt float32, player *models.Player,
	playerIsAttacking bool,
) {
	if enemy.IsAlive() {
		enemy.PathfindingBehaviourID.Get().Follow(pathfinding.BehaviourContext{
			Self: &enemy.Body,
			PlayerPos: player.Pos.Mgl32(),
		})
	} else {
		// dead men don't walk
		enemy.PathfindingBehaviourID = pathfinding.FreeBehaviourID
		enemy.Body.Vel.X = 0
		enemy.TimeSinceDeath += dt
	}
}

func (enemy *Enemy) Draw(cam *camera.Camera) {
	camX := enemy.Pos.X - cam.X
	camY := enemy.Pos.Y - cam.Y

	drawOpts := spriteloader.NewDrawOptions()
	drawOpts.Alpha = enemy.Alpha()
	spriteloader.DrawSpriteQuadOptions(
		camX, camY,
		enemy.Size.X, enemy.Size.Y,
		basicEnemySpriteId, drawOpts,
	)
}
