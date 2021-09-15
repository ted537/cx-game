package item

// TODO where to store state about previous mouse position

import (
	"github.com/go-gl/mathgl/mgl32"

	"github.com/skycoin/cx-game/render"
)

func pipetoolMouseDown(info ItemUseInfo) {

}

func pipetoolMouseDrag(info ItemUseInfo, prev mgl32.Vec2) {

}

func RegisterPipeToolItemType() ItemTypeID {
	itemtype := NewItemType(render.GetSpriteIDByName("dev-pipe-place-tool"))
	itemtype.Name = "Dev Pipe Place Tool"
	itemtype.Category = BuildTool
	itemtype.Use = UseBuildTool
	itemtype.OnDrag = pipetoolMouseDrag
	return AddItemType(itemtype)
}
