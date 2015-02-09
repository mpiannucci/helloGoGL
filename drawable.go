package main

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type drawable interface {
	GetID() string
	SetID(id string)
	SetTranslation(x, y, z float32)
	SetColor(r, g, b float32)
	InitBuffers()
	BindBuffers()
	Draw()
}

type triangle struct {
	id            string
	bufferData    []float32
	vertexArray   gl.VertexArray
	buffer        gl.Buffer
	shader        gl.Program
	offsetUniform gl.UniformLocation
	xyOffset      mgl32.Vec2
	colorUniform  gl.UniformLocation
	color         mgl32.Vec3
}

func (t *triangle) GetID() string {
	return t.id
}

func (t *triangle) SetID(id string) {
	t.id = id
}

func (t *triangle) SetTranslation(x, y, z float32) {
	t.xyOffset = mgl32.Vec2{x, y}
}

func (t *triangle) SetColor(r, g, b float32) {
	t.color = mgl32.Vec3{r, g, b}
}

func (t *triangle) InitBuffers() {
	// Initialize to a basic triangle
	t.bufferData = []float32{
		-3.0, 0.0, 0.0,
		3.0, 0.0, 0.0,
		0.0, 5.0, 0.0}

	// Create and Bind Vertex Arrays
	t.vertexArray = gl.GenVertexArray()
	t.vertexArray.Bind()

	// Load shaders
	t.shader = MakeShaderProgram("shape.vs", "shape.fs")

	t.offsetUniform = t.shader.GetUniformLocation("Offset")
	t.colorUniform = t.shader.GetUniformLocation("ColorVector")

	// Set Some defaults
	t.SetTranslation(0, 0, 0)
	t.SetColor(0, 0.0, 1.0)
}

func (t *triangle) BindBuffers() {
	// Create and bind data buffers
	t.buffer = gl.GenBuffer()
	t.buffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(t.bufferData)*4, &t.bufferData[0], gl.STATIC_DRAW)
}

func (t *triangle) Draw() {
	t.BindBuffers()

	// Load Shaders
	t.shader.Use()

	t.offsetUniform.Uniform2f(t.xyOffset.X(), t.xyOffset.Y())
	t.colorUniform.Uniform3f(t.color.X(), t.color.Y(), t.color.Z())

	// Load Arrays
	attribLoc := gl.AttribLocation(0)
	attribLoc.EnableArray()
	t.buffer.Bind(gl.ARRAY_BUFFER)
	attribLoc.AttribPointer(3, gl.FLOAT, false, 0, nil)

	// Draw the arrays
	gl.DrawArrays(gl.TRIANGLES, 0, 3)

	// Clean up
	attribLoc.DisableArray()
}
