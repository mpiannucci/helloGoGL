package main

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// Get a new triangle object
func CreateTriangle() *polygon2d {
	return CreatePolygon(triangle)
}

// Get a new Square object
func CreateSquare() *polygon2d {
	return CreatePolygon(square)
}

// Get a new rectangle object
func CreateRectangle() *polygon2d {
	return CreatePolygon(rectangle)
}

// Get a new shape object of your choice
func CreatePolygon(shape PolygonShape) *polygon2d {
	p := new(polygon2d)
	p.SetShape(shape)
	p.InitBuffers()
	return p
}

// Abstract interface for OpenGL compatible drawable objects
type drawable interface {
	ID() string
	SetID(id string)
	SetTranslation(x, y, z float32)
	SetRotation(angle float32)
	SetColor(r, g, b float32)
	UpdateMVPMatrix()
	InitBuffers()
	BindBuffers()
	Draw()
}

// PloygonShape type to create polygon instances with
type PolygonShape int

// Types of polygons
const (
	triangle  PolygonShape = iota
	square    PolygonShape = iota
	rectangle PolygonShape = iota
)

// Two Dimensional Polygon Drawable
type polygon2d struct {
	id       string
	vertices []float32
	indices  []gl.GLuint
	shape    PolygonShape

	// Buffers
	vertexArray   gl.VertexArray
	vertexBuffer  gl.Buffer
	elementBuffer gl.Buffer

	// Shaders
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

// Get the id of the polygon
func (p *polygon2d) GetID() string {
	return p.id
}

// Set the id of the polygon
func (p *polygon2d) SetID(id string) {
	p.id = id
}

// Set the translation of the polygon
func (p *polygon2d) SetTranslation(x, y, z float32) {
	p.xyOffset = mgl32.Vec2{x, y}
}

// Set the rotation of the polygon
func (p *polygon2d) SetRotation(angle float32) {
	p.rotAngle = angle
}

// Set the color to draw the polygon
func (p *polygon2d) SetColor(r, g, b float32) {
	p.color = mgl32.Vec3{r, g, b}
}

// Set the shape of the polygon
func (p *polygon2d) SetShape(shape PolygonShape) {
	p.shape = shape
	switch p.shape {
	case triangle:
		p.vertices = []float32{
			0.0, 0.0, 0.0,
			1.0, 0.0, 0.0,
			0.5, 1.0, 0.0}
		p.indices = []gl.GLuint{0, 1, 2}
	case square:
		p.vertices = []float32{
			0.0, 0.0, 0.0,
			1.0, 0.0, 0.0,
			1.0, 1.0, 0.0,
			0.0, 1.0, 0.0}
		p.indices = []gl.GLuint{
			0, 1, 2,
			0, 2, 3}
	case rectangle:
		p.vertices = []float32{
			0.0, 0.0, 0.0,
			2.0, 0.0, 0.0,
			2.0, 1.0, 0.0,
			0.0, 1.0, 0.0}
		p.indices = []gl.GLuint{
			0, 1, 2,
			0, 2, 3}
	}
}

// Update the Model View Projection matrix for rendering in the shader
func (p *polygon2d) UpdateMVPMatrix() {
	p.model = mgl32.Ident4().Mul4(mgl32.HomogRotate3DZ(p.rotAngle)).Mul4(mgl32.Translate3D(p.xyOffset.X(), p.xyOffset.Y(), 0))
	p.mvp = p.projection.Mul4(p.model)
}

// Initialize the buffers
func (p *polygon2d) InitBuffers() {
	// Create and Bind Vertex Arrays
	p.vertexArray = gl.GenVertexArray()
	p.vertexArray.Bind()

	// Load shaders
	p.shader = MakeShaderProgram("shape.vs", "shape.fs")

	// Get the uniform locations
	p.mvpUniform = p.shader.GetUniformLocation("MVP")
	p.colorUniform = p.shader.GetUniformLocation("ColorVector")

	// Initialize projection matrices
	p.projection = mgl32.Ortho2D(-10, 10, -10, 10)
	p.view = mgl32.LookAt(0.0, 0.0, 5.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0)

	// Set Some defaults
	p.SetTranslation(0.0, 0.0, 0.0)
	p.SetRotation(0.0)
	p.SetColor(0.0, 0.0, 0.0)
}

// Bind the Buffers
func (p *polygon2d) BindBuffers() {
	// Create and bind vertex buffers
	p.vertexBuffer = gl.GenBuffer()
	p.vertexBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(p.vertices)*4, p.vertices, gl.STATIC_DRAW)

	// Create and bind the element buffers
	p.elementBuffer = gl.GenBuffer()
	p.elementBuffer.Bind(gl.ELEMENT_ARRAY_BUFFER)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(p.indices)*4, p.indices, gl.STATIC_DRAW)
}

// Render the polygon
func (p *polygon2d) Draw() {
	p.UpdateMVPMatrix()
	p.BindBuffers()

	// Load Shaders
	p.shader.Use()

	// Pass uniforms so the shader
	p.mvpUniform.UniformMatrix4fv(false, p.mvp)
	p.colorUniform.Uniform3f(p.color.X(), p.color.Y(), p.color.Z())

	// Load Arrays
	attribLoc := gl.AttribLocation(0)
	attribLoc.EnableArray()
	defer attribLoc.DisableArray()

	// Bind the buffer again
	p.vertexBuffer.Bind(gl.ARRAY_BUFFER)
	defer p.vertexBuffer.Unbind(gl.ARRAY_BUFFER)
	attribLoc.AttribPointer(3, gl.FLOAT, false, 0, nil)

	// Bind the element buffer
	p.elementBuffer.Bind(gl.ELEMENT_ARRAY_BUFFER)
	defer p.elementBuffer.Unbind(gl.ELEMENT_ARRAY_BUFFER)

	// Draw the arrays
	gl.DrawElements(gl.TRIANGLES, len(p.indices), gl.UNSIGNED_INT, nil)

	// Clean up
	attribLoc.DisableArray()
}
