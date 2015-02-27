package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/mpiannucci/helloGoGL/drawable"
)

var (
	shapes       []drawable.Drawable
	move         float32
	colorCounter int
	colorShift   int
	scale        float32
	window       *glfw.Window
	windowErr    error
)

// Initialize OpenGL
func Init() {
	// Set the scale
	scale = 3.0

	blueTriangle := drawable.CreateTriangle()
	blueTriangle.SetScale(scale)
	blueTriangle.SetColor(0.2, 0.2, 1.0)

	redRect := drawable.CreateSquare()
	redRect.SetScale(scale)
	redRect.SetColor(1.0, 0.2, 0.2)

	greenCircle := drawable.CreateCircle(0.5)
	greenCircle.SetScale(scale)
	greenCircle.SetColor(0.2, 1.0, 0.2)

	// Pack all of the drawables into one array
	shapes = []drawable.Drawable{
		blueTriangle,
		redRect,
		greenCircle}

	// Initialize the animation switch and the color change switch
	colorCounter = 0
	colorShift = 0

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
		case shape.Shape() != drawable.CircleShape:
			// If its a square, rectangle or triangle, the drawing starts in the left corner
			// So account for that and center it
			xPosition = xPosition - (float32(scale) * 0.5)
		}
		shape.SetTranslation(xPosition, 0.0, 0.0)
	}
}

// Change colors of the shapes
func SwitchColors() {
	// Only change the color every so often
	if colorCounter%30 == 0 {
		shapes[(colorShift+0)%len(shapes)].SetColor(0.2, 0.2, 1.0)
		shapes[(colorShift+1)%len(shapes)].SetColor(1.0, 0.2, 0.2)
		shapes[(colorShift+2)%len(shapes)].SetColor(0.2, 1.0, 0.2)
		colorShift++
	}
	colorCounter++
}

// Animate
func Animate() {

	// Get the key states
	rightPress := window.GetKey(glfw.KeyRight)
	leftPress := window.GetKey(glfw.KeyLeft)
	upPress := window.GetKey(glfw.KeyUp)
	downPress := window.GetKey(glfw.KeyDown)

	// If the horizontal keys were press change colors
	if rightPress == glfw.Press || leftPress == glfw.Press {
		SwitchColors()
	}

	// If the vertical keys were pressed then animate the shapes in the right direction,
	// but make sure the shapes stay in the main window
	if upPress == glfw.Press {
		if move < 7.0 {
			move += 0.2 * 1.0
		}
	} else if downPress == glfw.Press {
		if move > -10.0 {
			move += 0.2 * -1.0
		}
	}

	// Get the new Y locations
	yLocation := float32(move)

	for _, shape := range shapes {
		// Set the new translated location of each of the shapes
		translate := shape.Translation()
		shape.SetTranslation(translate.X(), yLocation, translate.Z())
	}
}

// Main Entry Point
func main() {
	runtime.LockOSThread()

	// Initialize the OpenGL Context
	glfwErr := glfw.Init()
	if glfwErr != nil {
		fmt.Fprintf(os.Stderr, "Can't open GLFW")
		return
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Samples, 4)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True) // needed for macs

	window, windowErr = glfw.CreateWindow(400, 400, "Hello Go GL", nil, nil)
	if windowErr != nil {
		fmt.Fprintf(os.Stderr, "%v\n", windowErr)
		return
	}

	window.MakeContextCurrent()

	gl.Init()
	gl.GetError() // Ignore error
	window.SetInputMode(glfw.StickyKeysMode, 1)

	// Window background color
	gl.ClearColor(0.9, 0.9, 0.9, 0.0)

	// Init objects
	Init()

	// Equivalent to a do... while
	for ok := true; ok; ok = (window.GetKey(glfw.KeyEscape) != glfw.Press && !window.ShouldClose()) {
		gl.Clear(gl.COLOR_BUFFER_BIT)

		Animate()

		// Draw the drawablesss
		for _, shape := range shapes {
			shape.Draw()
		}

		// Swap Buffers
		window.SwapBuffers()
		glfw.PollEvents()
	}

}
