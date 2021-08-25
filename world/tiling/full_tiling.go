package tiling

type FullTiling struct {}

func (t FullTiling) Count() int { return 11*5 }
func (t FullTiling) Index(n Neighbours) int {
	return 1
}
