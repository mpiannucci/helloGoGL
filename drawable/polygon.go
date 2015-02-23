package drawable

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// Two Dimensional Polygon Drawable
type Polygon2d struct {
	// Shared drawable attributes
	*Attributes

	// Shape specific properties
	vertices   []float32
	indices    []uint32
	shape_type ShapeType

	// Buffers
	vertexArray   uint32
	vertexBuffer  uint32
	elementBuffer uint32

	// Shaders
	shader       uint32
	mvpUniform   int32
	colorUniform int32
}

// Set the shape of the polygon
func (p *Polygon2d) SetShape(shape_type ShapeType) {
	p.shape_type = shape_type
	switch p.shape_type {
	case TriangleShape:
		p.vertices = []float32{
			0.0, 0.0, 0.0,
			1.0, 0.0, 0.0,
			0.5, 1.0, 0.0}
		p.indices = []uint32{0, 1, 2}
	case SquareShape:
		p.vertices = []float32{
			0.0, 0.0, 0.0,
			1.0, 0.0, 0.0,
			1.0, 1.0, 0.0,
			0.0, 1.0, 0.0}
		p.indices = []uint32{
			0, 1, 2,
			0, 2, 3}
	case RectangleShape:
		p.vertices = []float32{
			0.0, 0.0, 0.0,
			2.0, 0.0, 0.0,
			2.0, 1.0, 0.0,
			0.0, 1.0, 0.0}
		p.indices = []uint32{
			0, 1, 2,
			0, 2, 3}
	}
}

// Get the shape
func (p *Polygon2d) Shape() ShapeType {
	return p.shape_type
}

// Initialize the buffers
func (p *Polygon2d) InitBuffers() {
	// Create and Bind Vertex Arrays
	gl.GenVertexArrays(1, &p.vertexArray)
	gl.BindVertexArray(p.vertexArray)

	// Load shaders
	var err error
	p.shader, err = MakeShaderProgram("Resources/shape.vs", "Resources/shape.fs")
	if err != nil {
		panic(err)
	}

	// Get the uniform locations
	p.mvpUniform = gl.GetUniformLocation(p.shader, gl.Str("MVP\x00"))
	p.colorUniform = gl.GetUniformLocation(p.shader, gl.Str("ColorVector\x00"))

	// Initialize projection matrices
	p.projection = mgl32.Ortho2D(-10, 10, -10, 10)
	p.view = mgl32.LookAt(0.0, 0.0, 5.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0)

	// Set Some defaults
	p.SetTranslation(0.0, 0.0, 0.0)
	p.SetRotation(0.0)
	p.SetColor(0.0, 0.0, 0.0)
}

// Bind the Buffers
func (p *Polygon2d) BindBuffers() {
	// Create and bind vertex buffers
	gl.GenBuffers(1, &p.vertexBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, p.vertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, int(len(p.vertices)*4), gl.Ptr(p.vertices), gl.STATIC_DRAW)

	// Create and bind the element buffers
	gl.GenBuffers(1, &p.elementBuffer)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, p.elementBuffer)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, int(len(p.indices)*4), gl.Ptr(p.indices), gl.STATIC_DRAW)
}

// Render the polygon
func (p *Polygon2d) Draw() {
	p.UpdateMVPMatrix()
	p.BindBuffers()

	// Load Shaders
	gl.UseProgram(p.shader)

	// Pass uniforms to the Shader
	gl.UniformMatrix4fv(p.mvpUniform, 1, false, &p.mvp[0])
	gl.Uniform3f(p.colorUniform, p.color.X(), p.color.Y(), p.color.Z())

	// Load Arrays
	gl.EnableVertexAttribArray(0)

	// Bind the buffer again
	gl.BindBuffer(gl.ARRAY_BUFFER, p.vertexBuffer)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	// Bind the element buffer
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, p.elementBuffer)

	// Draw the arrays
	gl.DrawElements(gl.TRIANGLES, int32(len(p.indices)), gl.UNSIGNED_INT, nil)

	// Clean up
	gl.DisableVertexAttribArray(0)
}
