package main

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type drawable interface {
	GetID() string
	SetID(id string)
	SetTranslation(x, y, z float32)
	SetRotation(angle float32)
	SetColor(r, g, b float32)
	UpdateMVPMatrix()
	InitBuffers()
	BindBuffers()
	Draw()
}

type triangle struct {
	id           string
	bufferData   []float32
	vertexArray  gl.VertexArray
	buffer       gl.Buffer
	shader       gl.Program
	mvpUniform   gl.UniformLocation
	xyOffset     mgl32.Vec2
	rotAngle     float32
	colorUniform gl.UniformLocation
	color        mgl32.Vec3
	projection   mgl32.Mat4
	view         mgl32.Mat4
	model        mgl32.Mat4
	mvp          mgl32.Mat4
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

func (t *triangle) SetRotation(angle float32) {
	t.rotAngle = angle
}

func (t *triangle) SetColor(r, g, b float32) {
	t.color = mgl32.Vec3{r, g, b}
}

func (t *triangle) UpdateMVPMatrix() {
	t.model = mgl32.Ident4().Mul4(mgl32.HomogRotate3DZ(t.rotAngle)).Mul4(mgl32.Translate3D(t.xyOffset.X(), t.xyOffset.Y(), 0))
	t.mvp = t.projection.Mul4(t.model)
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

	// Get the uniform locations
	t.mvpUniform = t.shader.GetUniformLocation("MVP")
	t.colorUniform = t.shader.GetUniformLocation("ColorVector")

	// Initialize projection matrices
	t.projection = mgl32.Ortho2D(-10, 10, -10, 10)
	t.view = mgl32.LookAt(0.0, 0.0, 5.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0)

	// Set Some defaults
	t.SetTranslation(0.0, 0.0, 0.0)
	t.SetRotation(0.0)
	t.SetColor(0.0, 0.0, 0.0)
}

func (t *triangle) BindBuffers() {
	// Create and bind data buffers
	t.buffer = gl.GenBuffer()
	t.buffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(t.bufferData)*4, &t.bufferData[0], gl.STATIC_DRAW)
}

func (t *triangle) Draw() {
	t.UpdateMVPMatrix()
	t.BindBuffers()

	// Load Shaders
	t.shader.Use()

	// Pass uniforms so the shader
	t.mvpUniform.UniformMatrix4fv(false, t.mvp)
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
