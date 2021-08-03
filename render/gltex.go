package render

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// wrapper over an opengl texture

type Texture struct {
	Target uint32
	Texture uint32
}

// may need to adjust this interface if we ever need simultaneous texture units

func (t Texture) Bind() {
	gl.BindTexture(t.Target, t.Texture)
}

func (t Texture) Unbind() {
	gl.BindTexture(t.Target, 0)
}
