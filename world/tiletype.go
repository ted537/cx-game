package world

import (
	"github.com/go-gl/mathgl/mgl32"

	"github.com/skycoin/cx-game/components/types"
	"github.com/skycoin/cx-game/cxmath"
	"github.com/skycoin/cx-game/render"
	"github.com/skycoin/cx-game/world/tiling"
)

type TileTypeID uint32
type TileType struct {
	Name          string
	Layer         LayerID
	ToolType      types.ToolType
	Placer        Placer
	Invulnerable  bool
	ID            TileTypeID
	MaterialID    MaterialID
	Width, Height int32
	Drops         Drops
	ItemSpriteID  render.SpriteID
	LightSource   bool
	LightAmount   uint8 // from 0 to 15
	NeedsGround   bool
}

func (tt TileType) Size() cxmath.Vec2i {
	return cxmath.Vec2i{tt.Width, tt.Height}
}

func (tt *TileType) Transform() mgl32.Mat4 {
	translate := mgl32.Translate3D(
		-0.5+float32(tt.Width)/2,
		-0.5+float32(tt.Height)/2,
		0,
	)
	scale := mgl32.Scale3D(
		float32(tt.Width), float32(tt.Height), 1)
	return translate.Mul4(scale)
}

type TileCreationOptions struct {
	Neighbours tiling.DetailedNeighbours
	FlipTransform mgl32.Mat4
}

func NewTileCreationOptions() TileCreationOptions {
	return TileCreationOptions {
		FlipTransform: mgl32.Ident4(),
	}
}

type TileUpdateOptions struct {
	Neighbours tiling.DetailedNeighbours
	Tile       *Tile
	Cycling    bool
}

func (tt TileType) CreateTile(opts TileCreationOptions) Tile {
	return tt.Placer.CreateTile(tt, opts)
}

func (tt TileType) UpdateTile(opts TileUpdateOptions) {
	tt.Placer.UpdateTile(tt, opts)
}

var tileTypes = make([]TileType, 0)
var tileTypeIDsByName = make(map[string]TileTypeID)

func RegisterTileType(
	name string, tileType TileType, ToolType types.ToolType,
) TileTypeID {
	id := TileTypeID(len(tileTypes))
	tileType.ID = id
	// fill in default size
	if tileType.Width == 0 {
		tileType.Width = 1
	}
	if tileType.Height == 0 {
		tileType.Height = 1
	}
	tileType.ItemSpriteID = tileType.Placer.ItemSpriteID()
	if tileType.Drops == nil {
		tileType.Drops = Drops{}
	}
	tileType.ToolType = ToolType
	tileTypes = append(tileTypes, tileType)
	tileTypeIDsByName[name] = id
	return id
}

func NextTileTypeID() TileTypeID {
	return TileTypeID(len(tileTypes))
}

func GetTileTypeByID(id TileTypeID) (TileType, bool) {
	ok := id < TileTypeID(len(tileTypes))
	if ok {
		return tileTypes[id], true
	} else {
		return TileType{}, false
	}
}

func IDFor(name string) (TileTypeID, bool) {
	id, ok := tileTypeIDsByName[name]
	return id, ok
}

func (id TileTypeID) Get() *TileType {
	return &tileTypes[id]
}

func AddDrop(id TileTypeID, drop Drop) {
	tileTypes[id].Drops = append(tileTypes[id].Drops, drop)
}

// not including air
func AllTileTypeIDs() []TileTypeID {
	ids := make([]TileTypeID, 0, len(tileTypes))
	for idx := TileTypeID(2); int(idx) < len(tileTypes); idx++ {
		ids = append(ids, idx)
	}
	return ids
}

func TileTypeIDsForToolType(toolType types.ToolType) []TileTypeID {
	ids := []TileTypeID{}
	for i := TileTypeID(2); int(i) < len(tileTypes); i++ {
		if tileTypes[i].ToolType == toolType {
			ids = append(ids, i)
		}
	}
	return ids
}
