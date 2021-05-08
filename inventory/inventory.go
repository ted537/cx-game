package inventory;

import (
	"log"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/skycoin/cx-game/spriteloader"
)

type InventorySlot struct {
	itemType uint32
	quantity uint32
}

type Inventory struct {
	Width, Height int
	slots []InventorySlot
}

var inventories = []Inventory{}

func NewInventory(width, height int) uint32 {
	inventories = append(inventories, Inventory {
		Width: width, Height: height,
		slots: make([]InventorySlot, width*height),
	})
	return uint32(len(inventories)-1)
}

func GetInventoryById(id uint32) Inventory {
	return inventories[id]
}

// TODO really need to figure out screen space solution

func (inventory Inventory) getBarSlots() []InventorySlot {
	start := inventory.Width*(inventory.Height-1)
	return inventory.slots[start:]
}

var gridScale float32 = 0.8
func (inventory Inventory) DrawGrid() {
	gridTransform :=
		mgl32.Translate3D(0,1,-spriteloader.SpriteRenderDistance).
		Mul4(mgl32.Scale3D(gridScale,gridScale,gridScale))
	for y:=0;y<inventory.Height;y++ {
		for x:=0;x<inventory.Width;x++ {
			xRender := float32(x) - float32(inventory.Width) / 2
			yRender := float32(y) - float32(inventory.Height) / 2
			slotTransform := gridTransform.
				Mul4(mgl32.Translate3D(xRender,yRender,0))
			spriteloader.DrawSpriteQuadMatrix(slotTransform, 0)
		}
	}
}

func (inv Inventory) DrawBar() {
	barTransform := mgl32.Translate3D(0,-3,-spriteloader.SpriteRenderDistance)
	barSlots := inv.getBarSlots()
	for idx,slot := range barSlots {
		x := float32(idx) - float32(len(barSlots)) / 2
		slotTransform := barTransform.
			Mul4(mgl32.Translate3D(x,0,0))
		_ = idx
		// TODO draw the correct sprite
		log.Print(slot)
		spriteloader.DrawSpriteQuadMatrix(slotTransform, 0)
	}
}
