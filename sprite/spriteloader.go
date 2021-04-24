package sprite

import (
	"log"
	"os"
	"image/png"
	"image"
	"image/draw"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"

	"github.com/skycoin/cx-game/render"
)

const SpriteRenderDistance = 10

var window *render.Window
// call this before loading any spritesheets
func InitSpriteloader(_window *render.Window) {
	window = _window
	quadVao = makeQuadVao()
}

type Spritesheet struct {
	tex uint32
	xScale, yScale float32
}

type Sprite struct {
	spriteSheetId int
	x, y int
}

var spritesheets = []Spritesheet{};
var sprites = []Sprite{};
var spriteIdsByName = make(map[string]int);

func LoadSpriteSheet(fname string) int {
	return LoadSpriteSheetWithSize(fname, 32,32)
}

func LoadSpriteSheetWithSize(fname string, spriteW,spriteH int) int {
	img := loadPng(fname)

	spritesheets = append(spritesheets, Spritesheet{
		xScale: float32(spriteW) / float32(img.Bounds().Dx()),
		yScale: float32(spriteH) / float32(img.Bounds().Dy()),
		tex: makeTexture(img),
	})

	return len(spritesheets)-1
}

func LoadSprite(spriteSheetId int, name string, x,y int) {
	sprites = append(sprites,Sprite{spriteSheetId,x,y})
	spriteIdsByName[name] = len(sprites)-1
}

func LoadSingleSprite(fname string, name string) {
	img := loadPng(fname)
	spritesheetId := LoadSpriteSheetWithSize(
		fname,
		img.Bounds().Dx(),img.Bounds().Dy(),
	)
	LoadSprite(spritesheetId,name,0,0)
}

func GetSpriteIdByName(name string) int {
	spriteId, ok := spriteIdsByName[name]
	if (!ok) {
		log.Fatalf("sprite with name [%v] does not exist",name)
	}
	return spriteId
}

func DrawSpriteQuad(xpos, ypos, xwidth, yheight float32, spriteId int) {
	sprite := sprites[spriteId]
	spritesheet := spritesheets[sprite.spriteSheetId]
	gl.UseProgram(window.Program)
	gl.Uniform1ui(
		gl.GetUniformLocation(window.Program, gl.Str("ourTexture\x00")),
		spritesheet.tex,
	)
	gl.Uniform2f(
		gl.GetUniformLocation(window.Program, gl.Str("texScale\x00")),
		spritesheet.xScale,spritesheet.yScale,
	)
	gl.Uniform2f(
		gl.GetUniformLocation(window.Program,gl.Str("texOffset\x00")),
		float32(sprite.x),float32(sprite.y),
	)

	worldTranslate := mgl32.Mat4.Mul4(
		mgl32.Translate3D(float32(xpos), float32(ypos), -SpriteRenderDistance),
		mgl32.Scale3D(float32(xwidth),float32(yheight),1),
	)
	gl.UniformMatrix4fv(
		gl.GetUniformLocation(window.Program,gl.Str("world\x00")),
		1,false,&worldTranslate[0],
	)

	gl.Enable(gl.TEXTURE_2D)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, spritesheet.tex)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	// sample nearest pixel such that we see nice pixel art
	// and not a blurry mess
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)

	gl.BindVertexArray(quadVao)
	gl.DrawArrays(gl.TRIANGLES, 0, 6)
}

// load a PNG image from disk into memory as RGBA
func loadPng(fname string) *image.RGBA {
	imgFile, err := os.Open(fname)

	if err != nil {
		log.Fatalln(err)
	}
	defer imgFile.Close()

	im, err := png.Decode(imgFile)
	if err != nil {
		log.Fatalln(err)
	}

	img := image.NewRGBA(im.Bounds())
	draw.Draw(img,img.Bounds(),im,image.Pt(0,0),draw.Src)
	return img
}

// upload an in-memory RGBA image to the GPU
func makeTexture(img *image.RGBA) uint32 {
	var tex uint32
	gl.GenTextures(1,&tex);

	gl.Enable(gl.TEXTURE_2D)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, tex)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	// sample nearest pixel such that we see nice pixel art
	// and not a blurry mess
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexImage2D(
		gl.TEXTURE_2D, 0, gl.RGBA,
		int32(img.Rect.Dx()), int32(img.Rect.Dy()), 0,
		gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(img.Pix),
	)

	return tex;
}

// x,y,z,u,v
var quadVertexAttributes = []float32{
	0.5, 0.5, 0, 1, 0,
	0.5, -0.5, 0, 1, 1,
	-0.5, 0.5, 0, 0, 0,

	0.5, -0.5, 0, 1, 1,
	-0.5, -0.5, 0, 0, 1,
	-0.5, 0.5, 0, 0, 0,
}

var quadVao uint32
func makeQuadVao() uint32 {
	var vbo uint32
	gl.GenBuffers(1,&vbo);

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(
		gl.ARRAY_BUFFER,
		4*len(quadVertexAttributes),
		gl.Ptr(quadVertexAttributes),
		gl.STATIC_DRAW,
	)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(4*3))
	gl.EnableVertexAttribArray(1)

	return vao
}


