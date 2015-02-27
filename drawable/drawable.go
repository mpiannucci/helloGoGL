package drawable

import (
	"errors"
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"io/ioutil"
	"strings"
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
	Translation() mgl32.Vec3
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
func MakeShaderProgram(vertFname, fragFname string) (uint32, error) {
	vertSource, err := ioutil.ReadFile(vertFname)
	if err != nil {
		panic(err)
	}

	fragSource, err := ioutil.ReadFile(fragFname)
	if err != nil {
		panic(err)
	}
	return newProgram(string(vertSource)+"\x00", string(fragSource)+"\x00")
}

func newProgram(vertexShaderSource, fragmentShaderSource string) (uint32, error) {
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return 0, errors.New(fmt.Sprintf("failed to link program: %v", log))
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csource := gl.Str(source)
	gl.ShaderSource(shader, 1, &csource, nil)
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
