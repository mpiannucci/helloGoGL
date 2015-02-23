package drawable

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
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
	Translation() mgl32.Vec2
	SetTranslation(x, y, z float32)
	Rotation() float32
	SetRotation(angle float32)
	Scale() float32
	SetScale(mag float32)
	Color() mgl32.Vec3
	SetColor(r, g, b float32)
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
	// First create the new instance of attributes with default parameters
	a := getDefaultAttributes()

	// Now create a new instance of the circle drawable with the newly created
	// attributes
	c := new(Circle)
	c.Attributes = a
	c.SetRadius(radius)
	c.InitBuffers()
	return c
}

// Get a new shape object of your choice
func CreatePolygon(shape ShapeType) *Polygon2d {
	// First create the new instance of attributes with default parameters
	a := getDefaultAttributes()

	// Now create a new instance of the polygon with the newly created
	// attributes.
	p := new(Polygon2d)
	p.Attributes = a
	p.SetShape(shape)
	p.InitBuffers()
	return p
}

// Get a new instace of the defaulted drawable attributes
func getDefaultAttributes() *Attributes {
	a := new(Attributes)
	a.SetID("randomID")
	a.SetColor(0, 0, 1)
	a.SetRotation(0)
	a.SetTranslation(0, 0, 0)
	a.SetScale(1)
	a.UpdateMVPMatrix()
	return a
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
