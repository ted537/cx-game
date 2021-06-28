package render

import (
	"github.com/go-gl/mathgl/mgl32"
)

// https://www.tilesetter.org/docs/generating_tilesets
type Neighbours struct {
	Left, Right, Up, Down bool
	UpLeft, UpRight, DownLeft, DownRight bool
}

func (n Neighbours) upLeftInnerCorner() bool {
	return n.Up && n.Left && !n.UpLeft
}

func (n Neighbours) upRightInnerCorner() bool {
	return n.Up && n.Right && !n.UpRight
}

func (n Neighbours) downLeftInnerCorner() bool {
	return n.Down && n.Left && !n.DownLeft
}

func (n Neighbours) downRightInnerCorner() bool {
	return n.Down && n.Right && !n.DownRight
}

func (n Neighbours) hasLeftInnerCorner() bool {
	return n.upLeftInnerCorner() || n.downLeftInnerCorner()
}

func (n Neighbours) hasRightInnerCorner() bool {
	return n.upRightInnerCorner() || n.downRightInnerCorner()
}

func (n Neighbours) hasUpInnerCorner() bool {
	return n.upLeftInnerCorner() || n.upRightInnerCorner()
}

func (n Neighbours) hasDownInnerCorner() bool {
	return n.downLeftInnerCorner() || n.downRightInnerCorner()
}

func (n Neighbours) countInnerCorners() int {
	return boolToInt(n.upLeftInnerCorner()) +
		boolToInt(n.upRightInnerCorner()) +
		boolToInt(n.downLeftInnerCorner()) +
		boolToInt(n.downRightInnerCorner())
}

func (n Neighbours) countEdges() int {
	return boolToInt(n.Left) + boolToInt(n.Up) +
		boolToInt(n.Right) + boolToInt(n.Down)
}

func boolToInt(x bool) int {
	if x { return 0 } else { return 1 }
}

// compute coords on the maps TileSetter provides in their "Blob" format
func (n Neighbours) blobCoords() (x,y int) {
	innerCorners := n.countInnerCorners()
	// default to the solid square
	x = 1
	y = 1
	if innerCorners == 0 {
		x = 1 + boolToInt(n.Up) - boolToInt(n.Down)
		y = 1 + boolToInt(n.Up) - boolToInt(n.Down)
		if !n.Left && !n.Right { x = 3 }
		if !n.Up && !n.Down { y = 3 }

		return x,y
	}
	if innerCorners == 1 {
		if !n.Left { x = 4 }
		if n.hasRightInnerCorner() { x = 5 }
		if n.hasLeftInnerCorner() { x = 6 }
		if !n.Right { x = 7 }

		if !n.Up { y = 0 }
		if !n.hasDownInnerCorner() { y = 1 }
		if !n.hasUpInnerCorner() { y = 2 }
		if !n.Down { y = 3 }
	}
	if innerCorners == 2 {

	}
	if innerCorners == 3 {

	}
	if innerCorners == 4 {

	}
	return x,y
}

func DrawBlobTile(mvp mgl32.Mat4, neighbours Neighbours) {
	
}
