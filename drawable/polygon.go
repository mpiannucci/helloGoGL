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
	indices    []gl.GLuint
	shape_type ShapeType

	// Buffers
	vertexArray   gl.VertexArray
	vertexBuffer  gl.Buffer
	elementBuffer gl.Buffer

	// Shaders
	shader       gl.Program
	mvpUniform   gl.UniformLocation
	colorUniform gl.UniformLocation
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
		p.indices = []gl.GLuint{0, 1, 2}
	case SquareShape:
		p.vertices = []float32{
			0.0, 0.0, 0.0,
			1.0, 0.0, 0.0,
			1.0, 1.0, 0.0,
			0.0, 1.0, 0.0}
		p.indices = []gl.GLuint{
			0, 1, 2,
			0, 2, 3}
	case RectangleShape:
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

// Get the shape
func (p *Polygon2d) Shape() ShapeType {
	return p.shape_type
}

// Initialize the buffers
func (p *Polygon2d) InitBuffers() {
	// Create and Bind Vertex Arrays
	p.vertexArray = gl.GenVertexArray()
	p.vertexArray.Bind()

	// Load shaders
	p.shader = MakeShaderProgram("Resources/shape.vs", "Resources/shape.fs")

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
func (p *Polygon2d) BindBuffers() {
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
func (p *Polygon2d) Draw() {
	p.UpdateMVPMatrix()
	p.BindBuffers()

	// Load Shaders
	p.shader.Use()
	defer p.shader.Unuse()

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
