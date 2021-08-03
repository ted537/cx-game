package item

import (
	"github.com/skycoin/cx-game/world"
	"github.com/skycoin/cx-game/ids"
)

func AddDrops() {
	world.AddDrop(world.IDFor("regolith"), world.Drop {
		Count: 1, Probability: 0.5, Item: ids.ItemTypeID(RockDustItemTypeID),
	})
}
