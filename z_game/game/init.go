package game

import (
	"log"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/skycoin/cx-game/agents"
	"github.com/skycoin/cx-game/camera"
	"github.com/skycoin/cx-game/components"
	"github.com/skycoin/cx-game/constants"
	"github.com/skycoin/cx-game/cxmath"
	"github.com/skycoin/cx-game/input"
	"github.com/skycoin/cx-game/item"
	"github.com/skycoin/cx-game/models"
	"github.com/skycoin/cx-game/particles"
	"github.com/skycoin/cx-game/render"
	"github.com/skycoin/cx-game/sound"
	"github.com/skycoin/cx-game/spriteloader"
	"github.com/skycoin/cx-game/spriteloader/anim"
	"github.com/skycoin/cx-game/starfield"
	"github.com/skycoin/cx-game/ui"
	"github.com/skycoin/cx-game/ui/console"
	"github.com/skycoin/cx-game/world"
	"github.com/skycoin/cx-game/world/mapgen"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

const (
	WINDOW_WIDTH  = 800
	WINDOW_HEIGHT = 600
)

var (
	Console console.Console
	Cam    *camera.Camera
	win    render.Window
	window *glfw.Window
	fps    *models.Fps
	player *agents.Agent

	World              world.World
	DrawCollisionBoxes = false
	FPS                int

	//unused
	isTileSelectorVisible = false
	worldItem             *item.WorldItem
)

func Init() {
	win = render.NewWindow(WINDOW_WIDTH, WINDOW_HEIGHT, true)
	// defer glfw.Terminate()

	window = win.Window

	window.SetMouseButtonCallback(mouseButtonCallback)
	window.SetScrollCallback(scrollCallback)
	window.SetSizeCallback(windowSizeCallback)

	input.Init(&win)
	sound.Init()
	spriteloader.DEBUG = false
	spriteloader.InitSpriteloader(&win)
	spriteloader.LoadSpritesFromConfigs()
	anim.InitAnimatedSpriteLoader()
	world.Init()
	item.InitWorldItem()
	ui.Init()
	particles.InitParticles()
	item.RegisterItemTypes()

	//models.Init()
	//player = models.NewPlayer()

	fps = models.NewFps(false)
	Cam = camera.NewCamera(&win)
	//World.Planet = world.NewDevPlanet()

	// TODO move this to the world package or similar
	World = world.World{
		Entities: world.Entities{
			Agents: *agents.NewAgentList(),
		},
		Planet: *mapgen.GeneratePlanet(),
	}
	components.ChangeWorld(&World)

	//World.Planet = *mapgen.GeneratePlanet()
	Cam.PlanetWidth = float32(World.Planet.Width)


	starfield.InitStarField(&win, Cam)


	worldTiles := World.Planet.GetAllTilesUnique()
	log.Printf("Found [%v] unique tiles in the world", len(worldTiles))

	spawnX := int(20)
	Cam.SetCameraPosition(float32(spawnX), 5)

	spawnPos := cxmath.Vec2 {
		float32(spawnX),
		float32(World.Planet.GetHeight(spawnX) + 10),
	}

	World.Entities.Agents.Spawn(
		constants.AGENT_TYPE_SLIME, agents.AgentCreationOptions{
			X: spawnPos.X - 6, Y: spawnPos.Y,
		},
	)
	World.Entities.Agents.Spawn(
		constants.AGENT_TYPE_SPIDER_DRILL, agents.AgentCreationOptions{
			X: spawnPos.X + 6, Y: spawnPos.Y,
		},
	)
	playerAgentID = World.Entities.Agents.Spawn(
		constants.AGENT_TYPE_PLAYER, agents.AgentCreationOptions {
			X: spawnPos.X, Y: spawnPos.Y,
		},
	)
	player = findPlayer()
	player.InventoryID = item.NewDevInventory()
	components.Init(&World, Cam, player)

	sound.LoadSound("player_jump", "jump.wav")
	Console = console.New()

}
