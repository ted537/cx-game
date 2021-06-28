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

func (n Neighbours) countInnerCorners() int {
	return boolToInt(n.upLeftInnerCorner()) +
		boolToInt(n.upRightInnerCorner()) +
		boolToInt(n.downLeftInnerCorner()) +
		boolToInt(n.downRightInnerCorner())
}

func boolToInt(x bool) int {
	if x { return 0 } else { return 1 }
}

func (n Neighbours) blobCoords() (int,int) {
	innerCorners := n.countInnerCorners()
	if innerCorners == 0 {

	}
	if innerCorners == 1 {

	}
	if innerCorners == 2 {

	}
	if innerCorners == 3 {

	}
	if innerCorners == 4 {

	}
	return -1,-1
}

func DrawBlobTile(mvp mgl32.Mat4, neighbours Neighbours) {
	
}
