package drawable

import "github.com/go-gl/mathgl/mgl32"

type Attributes struct {
	id            string
	xyTranslation mgl32.Vec3
	rotationAngle float32
	scale         float32
	color         mgl32.Vec3

	projection mgl32.Mat4
	view       mgl32.Mat4
	model      mgl32.Mat4
	mvp        mgl32.Mat4
}

// Get the id of the drawable
func (a *Attributes) ID() string {
	return a.id
}

// Set the id of the drawable
func (a *Attributes) SetID(id string) {
	a.id = id
}

// Get the translation of the drawable
func (a *Attributes) Translation() mgl32.Vec3 {
	return a.xyTranslation
}

// Set the translation of the polygon
func (a *Attributes) SetTranslation(x, y, z float32) {
	a.xyTranslation = mgl32.Vec3{x, y, z}
}

// Get the rotation of the drawable
func (a *Attributes) Rotation() float32 {
	return a.rotationAngle
}

// Set the rotation of the drawable
func (a *Attributes) SetRotation(angle float32) {
	a.rotationAngle = angle
}

// Get the scale of the drawable
func (a *Attributes) Scale() float32 {
	return a.scale
}

// Set the scale of the drawable
func (a *Attributes) SetScale(mag float32) {
	a.scale = mag
}

// Get the color of the drawable
func (a *Attributes) Color() mgl32.Vec3 {
	return a.color
}

// Set the color to draw the drawable
func (a *Attributes) SetColor(r, g, b float32) {
	a.color = mgl32.Vec3{r, g, b}
}

// Update the Model View Projection matrix for rendering in the shader
func (a *Attributes) UpdateMVPMatrix() {
	a.model = mgl32.Ident4().Mul4(mgl32.HomogRotate3DZ(a.rotationAngle)).Mul4(mgl32.Translate3D(a.xyTranslation.X(), a.xyTranslation.Y(), 0)).Mul4(mgl32.Scale3D(a.scale, a.scale, 0))
	a.mvp = a.projection.Mul4(a.model)
}
