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
	vertexArray  uint32
	vertexBuffer uint32

	// Shaders
	shader       uint32
	mvpUniform   int32
	colorUniform int32
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
	gl.GenVertexArrays(1, &c.vertexArray)
	gl.BindVertexArray(c.vertexArray)

	// Get the shader
	var err error
	c.shader, err = MakeShaderProgram("Resources/shape.vs", "Resources/shape.fs")
	if err != nil {
		panic(err)
	}

	// Get the uniform locations
	c.mvpUniform = gl.GetUniformLocation(c.shader, gl.Str("MVP\x00"))
	c.colorUniform = gl.GetUniformLocation(c.shader, gl.Str("ColorVector\x00"))

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
	gl.GenBuffers(1, &c.vertexBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, c.vertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, int(len(c.vertices)*4), gl.Ptr(c.vertices), gl.STATIC_DRAW)
}

// Render the circle
func (c *Circle) Draw() {
	c.UpdateMVPMatrix()
	c.BindBuffers()

	// Load Shader
	gl.UseProgram(c.shader)

	// Pass uniforms so the shader
	gl.UniformMatrix4fv(c.mvpUniform, 1, false, &c.mvp[0])
	gl.Uniform3f(c.colorUniform, c.color.X(), c.color.Y(), c.color.Z())

	// Load Arrays
	gl.EnableVertexAttribArray(0)

	// Bind the buffer again
	gl.BindBuffer(gl.ARRAY_BUFFER, c.vertexBuffer)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	// Draw the arrays
	gl.DrawArrays(gl.TRIANGLE_FAN, 0, int32(len(c.vertices)))

	// Clean up
	gl.DisableVertexAttribArray(0)
}
