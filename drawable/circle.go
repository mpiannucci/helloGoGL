package drawable

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

// Two Dimensional Polygon Drawable
type Circle struct {
	// Shared drawable attributes
	*Attributes

	// Shape specific properties
	radius     float32
	vertices   []float32
	shape_type ShapeType

	// Buffers
	vertexArray  gl.VertexArray
	vertexBuffer gl.Buffer

	// Shaders
	shader       gl.Program
	mvpUniform   gl.UniformLocation
	colorUniform gl.UniformLocation
}

// Set the shape of the circle
func (c *Circle) Shape() ShapeType {
	return c.shape_type
}

// Set the radius of the circle
func (c *Circle) SetRadius(radius float32) {
	c.radius = radius

	// Create the new vertices
	c.vertices = []float32{0.0, radius, 0.0}
	num_segments := 20
	for segment := 0; segment <= num_segments; segment++ {
		vertX := c.vertices[0] + radius*float32(math.Cos(float64(segment)*2.0*math.Pi/float64(num_segments)))
		vertY := c.vertices[1] + radius*float32(math.Sin(float64(segment)*2.0*math.Pi/float64(num_segments)))
		c.vertices = append(c.vertices, vertX, vertY, 0.0)
	}
}

// Initialize the buffers
func (c *Circle) InitBuffers() {
	// Identify as a circle
	c.shape_type = CircleShape

	// Create and Bind Vertex Arrays
	c.vertexArray = gl.GenVertexArray()
	c.vertexArray.Bind()

	// Load shaders
	c.shader = MakeShaderProgram("Resources/shape.vs", "Resources/shape.fs")

	// Get the uniform locations
	c.mvpUniform = c.shader.GetUniformLocation("MVP")
	c.colorUniform = c.shader.GetUniformLocation("ColorVector")

	// Initialize projection matrices
	c.projection = mgl32.Ortho2D(-10, 10, -10, 10)
	c.view = mgl32.LookAt(0.0, 0.0, 5.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0)

	// Set Some defaults
	c.SetTranslation(0.0, 0.0, 0.0)
	c.SetRotation(0.0)
	c.SetColor(0.0, 0.0, 0.0)
}

// Bind the Buffers
func (c *Circle) BindBuffers() {
	// Create and bind vertex buffers
	c.vertexBuffer = gl.GenBuffer()
	c.vertexBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(c.vertices)*4, c.vertices, gl.STATIC_DRAW)
}

// Render the circle
func (c *Circle) Draw() {
	c.UpdateMVPMatrix()
	c.BindBuffers()

	// Load Shaders
	c.shader.Use()
	defer c.shader.Unuse()

	// Pass uniforms so the shader
	c.mvpUniform.UniformMatrix4fv(false, c.mvp)
	c.colorUniform.Uniform3f(c.color.X(), c.color.Y(), c.color.Z())

	// Load Arrays
	attribLoc := gl.AttribLocation(0)
	attribLoc.EnableArray()
	defer attribLoc.DisableArray()

	// Bind the buffer again
	c.vertexBuffer.Bind(gl.ARRAY_BUFFER)
	defer c.vertexBuffer.Unbind(gl.ARRAY_BUFFER)
	attribLoc.AttribPointer(3, gl.FLOAT, false, 0, nil)

	// Draw the arrays
	gl.DrawArrays(gl.TRIANGLE_FAN, 0, len(c.vertices))

	// Clean up
	attribLoc.DisableArray()
}
