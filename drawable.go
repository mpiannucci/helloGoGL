package main

import (
	"github.com/go-gl/gl"
)

type drawable interface {
	GetID() string
	SetID(id string)
	InitBuffers()
	BindBuffers()
	Draw()
}

type triangle struct {
	id          string
	bufferData  []float32
	vertexArray gl.VertexArray
	buffer      gl.Buffer
	shader      gl.Program
}

func (t *triangle) GetID() string {
	return t.id
}

func (t *triangle) SetID(id string) {
	t.id = id
}

func (t *triangle) InitBuffers() {
	// Initialize to a basic triangle
	t.bufferData = []float32{
		-3.0, 0.0, 0.,
		3.0, 0.0, 0.,
		0., 5., 0.}

	// Create and Bind Vertex Arrays
	t.vertexArray = gl.GenVertexArray()
	t.vertexArray.Bind()

	// Load shaders
	t.shader = MakeShaderProgram("simpleshade.vs", "simpleshade.fs")
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
