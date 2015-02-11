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
		-1.0, 0.0, 0.0,
		1.0, 0.0, 0.0,
		0.0, 2.0, 0.0}

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

type rectangle struct {
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

func (r *rectangle) GetID() string {
	return r.id
}

func (r *rectangle) SetID(id string) {
	r.id = id
}

func (r *rectangle) SetTranslation(x, y, z float32) {
	r.xyOffset = mgl32.Vec2{x, y}
}

func (r *rectangle) SetRotation(angle float32) {
	r.rotAngle = angle
}

func (R *rectangle) SetColor(r, g, b float32) {
	R.color = mgl32.Vec3{r, g, b}
}

func (r *rectangle) UpdateMVPMatrix() {
	r.model = mgl32.Ident4().Mul4(mgl32.HomogRotate3DZ(r.rotAngle)).Mul4(mgl32.Translate3D(r.xyOffset.X(), r.xyOffset.Y(), 0))
	r.mvp = r.projection.Mul4(r.model)
}

func (r *rectangle) InitBuffers() {
	// Initialize to a basic triangle
	r.bufferData = []float32{
		0.0, 0.0, 0.0,
		0.0, 2.0, 0.0,
		2.0, 0.0, 0.0,
		2.0, 2.0, 0.0}

	// Create and Bind Vertex Arrays
	r.vertexArray = gl.GenVertexArray()
	r.vertexArray.Bind()

	// Load shaders
	r.shader = MakeShaderProgram("shape.vs", "shape.fs")

	// Get the uniform locations
	r.mvpUniform = r.shader.GetUniformLocation("MVP")
	r.colorUniform = r.shader.GetUniformLocation("ColorVector")

	// Initialize projection matrices
	r.projection = mgl32.Ortho2D(-10, 10, -10, 10)
	r.view = mgl32.LookAt(0.0, 0.0, 5.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0)

	// Set Some defaults
	r.SetTranslation(0.0, 0.0, 0.0)
	r.SetRotation(0.0)
	r.SetColor(0.0, 0.0, 0.0)
}

func (r *rectangle) BindBuffers() {
	// Create and bind data buffers
	r.buffer = gl.GenBuffer()
	r.buffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(r.bufferData)*4, &r.bufferData[0], gl.STATIC_DRAW)
}

func (r *rectangle) Draw() {
	r.UpdateMVPMatrix()
	r.BindBuffers()

	// Load Shaders
	r.shader.Use()

	// Pass uniforms so the shader
	r.mvpUniform.UniformMatrix4fv(false, r.mvp)
	r.colorUniform.Uniform3f(r.color.X(), r.color.Y(), r.color.Z())

	// Load Arrays
	attribLoc := gl.AttribLocation(0)
	attribLoc.EnableArray()
	r.buffer.Bind(gl.ARRAY_BUFFER)
	attribLoc.AttribPointer(3, gl.FLOAT, false, 0, nil)

	// Draw the arrays
	gl.DrawArrays(gl.TRIANGLE_STRIP, 0, 6)

	// Clean up
	attribLoc.DisableArray()
}
