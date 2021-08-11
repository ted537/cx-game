package render

import (
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type Window struct {
	Width,Height    int
	PhysicalWidth,PhysicalHeight int
	Resizable bool
	Window    *glfw.Window
	context   Context
}


func NewWindow(width, height int, resizable bool) Window {
	glfwWindow := initGlfw(width, height, resizable)
	initOpenGL()

	InitQuadVao()

	//temporary, to set projection matrix

	window := Window{
		Width:     width,
		Height:    height,
		Resizable: resizable,
		Window:    glfwWindow,
	}
	window.context = window.DefaultRenderContext()

	return window
}

// initGlfw initializes glfw and returns a Window to use.
func initGlfw(width, height int, resizable bool) *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	var res int
	if resizable {
		res = glfw.True
	} else {
		res = glfw.False
	}

	glfw.WindowHint(glfw.Resizable, res)
	glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "CX Game", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

func initOpenGL() {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	InitDrawLines()
	lineProgram = CompileProgram(
		"./assets/shader/line.vert", "./assets/shader/line.frag")
}

var Projection mgl32.Mat4

func (window *Window) GetProjectionMatrix() mgl32.Mat4 {
	// return window.DefaultRenderContext().Projection
	return window.context.Projection
}

func (window *Window) SetProjectionMatrix(projection mgl32.Mat4) {
	window.context.Projection = projection
	Projection = projection
}
