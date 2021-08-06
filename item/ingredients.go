package item

import (
	"github.com/skycoin/cx-game/render"
)

func RegisterRockDustItemType() ItemTypeID {
	/*
	sprite := spriteloader.LoadSingleSprite(
		"./assets/item/rock1_dust.png", "rock-dust" )
	*/
	itemtype := NewItemType(render.GetSpriteIDByName("rock-dust"))
	itemtype.Name = "Rock Dust"
	return AddItemType(itemtype)
}
