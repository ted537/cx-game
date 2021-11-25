package world

import (
	"log"
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

func (planet *Planet) electricTileIndices() []int {
	indices := []int{}
	for y := 0 ; y < int(planet.Height) ; y++ {
		for x := 0 ; x < int(planet.Width) ; x++ {
			tile,ok := planet.GetTile(x,y, TopLayer)
			if ok && tile.IsElectric() {
				idx := planet.GetTileIndex(x,y)
				indices = append(indices, idx)
			}
		}
	}
	return indices
}

func (planet *Planet) DetectCircuits() {
	for _,tileIdx := range planet.electricTileIndices() {
		log.Printf("detected electric tile at idx = %d", tileIdx)
	}
}
