package main

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
	"io/ioutil"
)

// PloygonShape type to create polygon instances with
type ShapeType int

// Types of polygons
const (
	triangle_shape  ShapeType = iota
	square_shape    ShapeType = iota
	rectangle_shape ShapeType = iota
	circle_shape    ShapeType = iota
)

// Abstract interface for OpenGL compatible drawable objects
type drawable interface {
	ID() string
	SetID(id string)
	Shape() ShapeType
	SetTranslation(x, y, z float32)
	SetRotation(angle float32)
	SetScale(mag float32)
	SetColor(r, g, b float32)
	UpdateMVPMatrix()
	InitBuffers()
	BindBuffers()
	Draw()
}

// Get a new triangle drawable
func CreateTriangle() *polygon2d {
	return CreatePolygon(triangle_shape)
}

// Get a new Square drawable
func CreateSquare() *polygon2d {
	return CreatePolygon(square_shape)
}

// Get a new rectangle drawable
func CreateRectangle() *polygon2d {
	return CreatePolygon(rectangle_shape)
}

// Get a new circle drawable
func CreateCircle(radius float32) *circle {
	c := new(circle)
	c.SetRadius(radius)
	c.InitBuffers()
	return c
}

// Get a new shape object of your choice
func CreatePolygon(shape ShapeType) *polygon2d {
	p := new(polygon2d)
	p.SetShape(shape)
	p.InitBuffers()
	return p
}

// Utility function to grab shaders
func MakeShaderProgram(vertFname, fragFname string) gl.Program {
	vertSource, err := ioutil.ReadFile(vertFname)
	if err != nil {
		panic(err)
	}

	fragSource, err := ioutil.ReadFile(fragFname)
	if err != nil {
		panic(err)
	}
	return glh.NewProgram(glh.Shader{gl.VERTEX_SHADER, string(vertSource)}, glh.Shader{gl.FRAGMENT_SHADER, string(fragSource)})
}
