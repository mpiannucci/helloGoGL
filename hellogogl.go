package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
)

var (
	shapes        []drawable
	animateSwitch float32
	move          float32
	scale         float32
)

// Initialize OpenGL
func Init() {
	// Set the scale
	scale = 3.0

	blueTriangle := CreateTriangle()
	blueTriangle.SetScale(scale)
	blueTriangle.SetColor(0, 0.2, 1.0)

	redRect := CreateSquare()
	redRect.SetScale(scale)
	redRect.SetColor(1.0, 0.2, 0.2)

	greenCircle := CreateCircle(0.5)
	greenCircle.SetScale(scale)
	greenCircle.SetColor(0.2, 1.0, 0.2)

	// Pack all of the drawables into one array
	shapes = []drawable{
		blueTriangle,
		redRect,
		greenCircle}

	animateSwitch = 1.0

	// Set the position of all of the drawables
	PositionDrawables()
}

// Position the shapes evenly in the X direction
func PositionDrawables() {
	// Equal distant between the shapes
	equiDistant := float32(20.0 / (len(shapes) + 1.0))

	for index, shape := range shapes {
		xPosition := -10.0 + (float32(index+1) * equiDistant)
		switch {
		case shape.Shape() != circle_shape:
			// If its a square, rectangle or triangle, the drawing starts in the left corner
			// So account for that and center it
			xPosition = xPosition - (float32(scale) * 0.5)
		}
		shape.SetTranslation(xPosition, 0.0, 0.0)
	}
}

// Animate
func Animate() {
	// Increment the move counter
	move += 0.2 * animateSwitch

	// Get the new Y locations
	blueY := float32(move)
	redY := float32(-1.0 * move)

	if move > 9.5 || move < -9.5 {
		// Make sure nothing leaves the window so switch directions
		animateSwitch *= -1.0
	}

	// Set the translation for the shapes
	shapes[0].SetTranslation(-5.0, blueY, 0.0)
	shapes[1].SetTranslation(5.0, redY, 0.0)
}

// Main Entry Point
func main() {
	runtime.LockOSThread()

	// Initialize the OpenGL Context
	if !glfw.Init() {
		fmt.Fprintf(os.Stderr, "Can't open GLFW")
		return
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Samples, 4)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenglProfile, glfw.OpenglCoreProfile)
	glfw.WindowHint(glfw.OpenglForwardCompatible, glfw.True) // needed for macs

	window, err := glfw.CreateWindow(400, 400, "Hello Go GL", nil, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	window.MakeContextCurrent()

	gl.Init()
	gl.GetError() // Ignore error
	window.SetInputMode(glfw.StickyKeys, 1)

	// Window background color
	gl.ClearColor(1.0, 1.0, 1.0, 0.0)

	// Init objects
	Init()

	// Equivalent to a do... while
	for ok := true; ok; ok = (window.GetKey(glfw.KeyEscape) != glfw.Press && !window.ShouldClose()) {
		gl.Clear(gl.COLOR_BUFFER_BIT)

		// Animate the triangles
		//Animate()

		// Draw the drawablesss
		for _, shape := range shapes {
			shape.Draw()
		}

		// Swap Buffers
		window.SwapBuffers()
		glfw.PollEvents()
	}

}
