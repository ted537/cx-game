package ui;

import (
	"github.com/skycoin/cx-game/spriteloader"
	"github.com/go-gl/mathgl/mgl32"
)

var charCodeToSpriteMap = make(map[int]int)
// TODO line wrapping
func drawString(text string, transform mgl32.Mat4) {
	pos := mgl32.Vec2 { 0,0 }
	for idx, charCode := range text {
		// font bitmap starts at ASCII code 33 (!)
		charIdx := int(charCode - 33)
		spriteId,ok := charCodeToSpriteMap[charIdx]
		if ok {
			letterTransform := transform.
				Mul4(mgl32.Translate3D(pos.X(),pos.Y(),0))
			spriteloader.DrawSpriteQuadMatrix(letterTransform, spriteId)
		}
		// TODO variable width fonts
		idx += 1
	}
}
