package inventory;

import "log"

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

func (inventory Inventory) DrawGrid() {
	for y:=0;y<inventory.Height;y++ {
		for x:=0;x<inventory.Width;x++ {
			// TODO
			_=x
			_=y
		}
	}
}

func (inv Inventory) DrawBar() {
	for idx,slot := range inv.getBarSlots() {
		_ = idx
		// TODO
		log.Print(slot)
	}
}
