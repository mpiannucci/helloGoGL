package drawable

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
	"io/ioutil"
)

// PloygonShape type to create polygon instances with
type ShapeType int

// Types of polygons
const (
	TriangleShape  ShapeType = iota
	SquareShape    ShapeType = iota
	RectangleShape ShapeType = iota
	CircleShape    ShapeType = iota
)

// Abstract interface for OpenGL compatible drawable objects
type Drawable interface {
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
func CreateTriangle() *Polygon2d {
	return CreatePolygon(TriangleShape)
}

// Get a new Square drawable
func CreateSquare() *Polygon2d {
	return CreatePolygon(SquareShape)
}

// Get a new rectangle drawable
func CreateRectangle() *Polygon2d {
	return CreatePolygon(RectangleShape)
}

// Get a new circle drawable
func CreateCircle(radius float32) *Circle {
	c := new(Circle)
	c.SetRadius(radius)
	c.InitBuffers()
	return c
}

// Get a new shape object of your choice
func CreatePolygon(shape ShapeType) *Polygon2d {
	p := new(Polygon2d)
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
