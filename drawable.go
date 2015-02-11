package main

// Abstract interface for OpenGL compatible drawable objects
type drawable interface {
	ID() string
	SetID(id string)
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
	return CreatePolygon(triangle)
}

// Get a new Square drawable
func CreateSquare() *polygon2d {
	return CreatePolygon(square)
}

// Get a new rectangle drawable
func CreateRectangle() *polygon2d {
	return CreatePolygon(rectangle)
}

// Get a new circle drawable
func CreateCircle(radius float32) *circle {
	c := new(circle)
	c.SetRadius(radius)
	c.InitBuffers()
	return c
}

// Get a new shape object of your choice
func CreatePolygon(shape PolygonShape) *polygon2d {
	p := new(polygon2d)
	p.SetShape(shape)
	p.InitBuffers()
	return p
}
