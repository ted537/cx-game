package world

import (
	"log"
	"github.com/skycoin/cx-game/cxmath"
	"github.com/skycoin/cx-game/constants"
)

type CircuitID uint32
type Circuit []uint32 // list of tile indices
type Circuits map[CircuitID]Circuit

func (c Circuit) Bind(planet *Planet) BoundCircuit {
	return BoundCircuit { Circuit:c, Planet: planet }
}

// circuit bound to a planet
type BoundCircuit struct {
	Circuit Circuit
	Planet  *Planet
}

func (bc *BoundCircuit) Tiles() []*Tile {
	topLayerTiles := bc.Planet.GetLayerTiles(TopLayer)
	tiles := []*Tile{}
	for _,tileIdx := range bc.Circuit {
		tile := &topLayerTiles[tileIdx]
		tiles = append(tiles,tile)
	}
	return tiles
}

func (bc *BoundCircuit) Wattage() int {
	wattage := 0
	for _,tile := range bc.Tiles() {
		wattage += tile.Power.Wattage
	}
	return wattage
}

func (bc *BoundCircuit) FixedTick() {
	active := bc.Wattage() > 0
	bc.Toggle(active)
}

func (bc *BoundCircuit) Toggle(active bool) {
	for _,tile := range bc.Tiles() {
		tile.Power.On = active
		tileType,ok := GetTileTypeByID(tile.TileTypeID)
		if ok {
			tileType.UpdateTile(TileUpdateOptions{Tile:tile})
		}
	}
}

func (planet *Planet) UpdateCircuits() {
	for _,circuit := range planet.Circuits {
		boundCircuit := circuit.Bind(planet)
		boundCircuit.FixedTick()
	}
}

func (planet *Planet) electricTilePositions() []cxmath.Vec2i {
	positions := []cxmath.Vec2i{}
	for y := 0 ; y < int(planet.Height) ; y++ {
		for x := 0 ; x < int(planet.Width) ; x++ {
			tile,ok := planet.GetTile(x,y, MidLayer)
			if ok && tile.IsElectric() {
				position := cxmath.Vec2i { int32(x), int32(y) }
				positions = append(positions, position)
			}
		}
	}
	return positions
}

func (planet *Planet) DetectCircuits() {
	log.Printf("detecting circuits")
	positions := planet.electricTilePositions()
	clusters := cxmath.FindClusters(positions, constants.POWER_REACH_RADIUS)
	log.Printf("found clusters: %v", clusters)
}
