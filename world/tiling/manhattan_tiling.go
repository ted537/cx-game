package tiling

const (
	manhattanTilingWidth int = 11
	manhattanTilingHeight int = 5
)

type ManhattanTiling struct {}

func (t ManhattanTiling) Count() int {
	return manhattanTilingWidth * manhattanTilingHeight
}

func (t ManhattanTiling) Index(n Neighbours) int {
	// block from (0,0) to (3,3)
	x := 1 + boolToInt(n.Left) - boolToInt(n.Right)
	y := 1 + boolToInt(n.Up) - boolToInt(n.Down)
	if !n.Left && !n.Right { x = 3 }
	if !n.Up && !n.Down { y = 3 }

	return manhattanTilingWidth*y + x
}
