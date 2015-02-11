package main

// Abstract interface for OpenGL compatible drawable objects
type drawable interface {
	ID() string
	SetID(id string)
	SetTranslation(x, y, z float32)
	SetRotation(angle float32)
	SetColor(r, g, b float32)
	UpdateMVPMatrix()
	InitBuffers()
	BindBuffers()
	Draw()
}

// Get a new triangle object
func CreateTriangle() *polygon2d {
	return CreatePolygon(triangle)
}

// Get a new Square object
func CreateSquare() *polygon2d {
	return CreatePolygon(square)
}

// Get a new rectangle object
func CreateRectangle() *polygon2d {
	return CreatePolygon(rectangle)
}
