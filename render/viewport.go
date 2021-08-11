package render

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"

	"github.com/skycoin/cx-game/cxmath"
)

type Viewport struct {
	X,Y,Width,Height int32
}

func (v Viewport) Use() {
	gl.Viewport(v.X, v.Y, v.Width, v.Height)
}

// fits target into frame, centered with black bars
// returns a transformation matrix
func fitCentered( virtualDims, physicalDims mgl32.Vec2) Viewport {
	// "physical" dimensions describe actual window size
	// "virtual" dimensions describe scaling of both world and UI
	// physical determines resolution.
	// virtual determines how big things are.
	physicalWidth := physicalDims.X()
	physicalHeight := physicalDims.Y()
	virtualWidth := virtualDims.X()
	virtualHeight := virtualDims.Y()

	scaleToFitWidth := physicalWidth / virtualWidth
	scaleToFitHeight := physicalHeight / virtualHeight
	// scale to fit entire virtual window in physical window
	scale := cxmath.Min(scaleToFitHeight, scaleToFitWidth)

	// scale up virtual dimensions to fit in physical dimensions.
	// in case of aspect ratio mismatch, black bars will appear
	viewportWidth := int32(virtualWidth*scale)
	viewportHeight := int32(virtualHeight*scale)

	// viewport offsets
	x := (int32(physicalWidth) - viewportWidth)/2
	y := (int32(physicalHeight) - viewportHeight)/2

	return Viewport { x, y, viewportWidth, viewportHeight }
}

func composePhysicalToViewportTransform(
		physicalDims, viewportDims mgl32.Vec2,
) mgl32.Mat4 {
	// TODO
	return mgl32.Ident4()
}
