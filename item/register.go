package item

var (
	LaserGunItemTypeID      ItemTypeID
	GunItemTypeID           ItemTypeID
	RockDustItemTypeID      ItemTypeID
	FurnitureToolItemTypeID ItemTypeID
	TileToolItemTypeID      ItemTypeID
	EnemyToolItemTypeID     ItemTypeID
	PipeToolItemTypeID      ItemTypeID
)

func RegisterItemTypes() {
	LaserGunItemTypeID = RegisterLaserGunItemType()
	GunItemTypeID = RegisterGunItemType()
	RockDustItemTypeID = RegisterRockDustItemType()
	FurnitureToolItemTypeID = RegisterFurnitureToolItemType()
	TileToolItemTypeID = RegisterTileToolItemType()
	PipeToolItemTypeID = RegisterPipeToolItemType()
	EnemyToolItemTypeID = RegisterEnemyToolItemType()

	AddDrops()
}
