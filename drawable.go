package main

import (
	"github.com/go-gl/gl"
)

type drawable interface {
	InitBuffers()
	Draw()
	GetID() string
	SetID(id string)
}

type triangle struct {
	id          string
	bufferData  []float32
	vertexArray gl.VertexArray
	buffer      gl.Buffer
	shader      gl.Program
}

func (t *triangle) InitBuffers() {
	t.bufferData = []float32{
		-2., -1., 0.,
		0.5, -1., 0.,
		0., 1., 0.}

	// Create and Bind Vertex Arrays
	t.vertexArray = gl.GenVertexArray()
	t.vertexArray.Bind()

	// Load shaders
	t.shader = MakeShaderProgram("simpleshade.vs", "simpleshade.fs")

	// Create and bind buffers
	t.buffer = gl.GenBuffer()
	t.buffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(t.bufferData)*4, &t.bufferData[0], gl.STATIC_DRAW)
}

func (t *triangle) Draw() {
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

func (t *triangle) GetID() string {
	return t.id
}

func (t *triangle) SetID(id string) {
	t.id = id
}