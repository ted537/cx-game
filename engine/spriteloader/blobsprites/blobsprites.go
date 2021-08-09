package blobsprites

import (
	"fmt"

<<<<<<< HEAD:spriteloader/blobsprites/blobsprites.go
	"github.com/skycoin/cx-game/render"
)

type BlobSpritesID uint32
var allBlobSprites = make(map[BlobSpritesID]([]render.SpriteID))
var nextBlobSpriteId = BlobSpritesID(1)

/*
func LoadBlobSprites(fname string, w,h int, name string) BlobSpritesID {
	spritesheetId := spriteloader.LoadSpriteSheetByColRow(
		fname, h, w )
	blobSprites := []render.SpriteID{}
	for idx:=0; idx < w*h; idx++ {
=======
	"github.com/skycoin/cx-game/engine/spriteloader"
	"github.com/skycoin/cx-game/render/blob"
)

type BlobSpritesID uint32

var allBlobSprites = make(map[BlobSpritesID]([]spriteloader.SpriteID))
var nextBlobSpriteId = BlobSpritesID(1)

func LoadBlobSprites(fname string, w, h int, name string) BlobSpritesID {
	spritesheetId := spriteloader.LoadSpriteSheetByColRow(
		fname, h, w)
	blobSprites := []spriteloader.SpriteID{}
	for idx := 0; idx < w*h; idx++ {
>>>>>>> main:engine/spriteloader/blobsprites/blobsprites.go
		y := idx / w
		x := idx % w
		name := fmt.Sprint("blob_%d", idx)
		blobSprites =
			append(blobSprites, spriteloader.LoadSprite(spritesheetId, name, x, y))
	}
	blobSpriteId := nextBlobSpriteId
	allBlobSprites[blobSpriteId] = blobSprites
	nextBlobSpriteId += 1
	return blobSpriteId
}

func LoadFullBlobSprites(fname string, name string) BlobSpritesID {
	return LoadBlobSprites(fname,
		blob.BlobSheetWidth, blob.BlobSheetHeight,
		name)
}

func LoadSimpleBlobSprites(fname string, name string) BlobSpritesID {
	return LoadBlobSprites(
		fname,
		blob.SimpleBlobSheetWidth, blob.SimpleBlobSheetHeight,
		name,
	)
}
*/

func GetBlobSpritesById(id BlobSpritesID) []render.SpriteID {
	return allBlobSprites[id]
}

func LoadIDFromSpritename(name string, n int) BlobSpritesID {
	sprites := make([]render.SpriteID, n)
	for idx := 0; idx < n; idx++ {
<<<<<<< HEAD:spriteloader/blobsprites/blobsprites.go
		spritename := fmt.Sprintf("%v:%d",name,idx)
		sprites[idx] = render.GetSpriteIDByName(spritename)
=======
		spritename := fmt.Sprintf("%v:%d", name, idx)
		sprites[idx] = spriteloader.GetSpriteIdByName(spritename)
>>>>>>> main:engine/spriteloader/blobsprites/blobsprites.go
	}
	id := nextBlobSpriteId
	allBlobSprites[id] = sprites
	nextBlobSpriteId++
	return id
}
