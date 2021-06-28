package render

import (
	"testing"
)

func TestBlobCoords(t *testing.T) {
	var gotX,gotY int
	var expectedX,expectedY int
	n := NewSolidNeighbours()
	n.DownRight = false
	gotX,gotY = n.blobCoords()
	expectedX = 5; expectedY = 2
	if gotX != expectedX || gotY != expectedY {
		t.Errorf("got [%v,%v] instead of [%v,%v] for bottom right corner",
			gotX,gotY, expectedX,expectedY )
	}
}
