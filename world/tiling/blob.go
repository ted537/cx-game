package tiling

import (
	"log"
)


const (
	BlobSheetWidth = 11
	BlobSheetHeight = 5

	SimpleBlobSheetWidth = 4
	SimpleBlobSheetHeight = 4
)

func ApplyBlobTiling(neighbours Neighbours) int {
	x,y := neighbours.blobCoords()
	idx := y * BlobSheetWidth + x
	return idx
}

// https://github.com/skycoin/cx-game/issues/205
func ApplySimpleBlobTiling(neighbours Neighbours) int {
	x,y := neighbours.simpleBlobCoords()
	idx := y * SimpleBlobSheetWidth + x
	return idx
}

type TilingType int
const (
	FullBlobTiling TilingType = iota
	SimpleBlobTiling 
)

func ApplyTiling(tt TilingType, neighbours Neighbours) int {
	if tt == FullBlobTiling { return ApplyBlobTiling(neighbours) }
	if tt == SimpleBlobTiling { return ApplySimpleBlobTiling(neighbours) }
	log.Fatalf("unknown tiling type")
	return -1
}
