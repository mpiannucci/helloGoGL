package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"github.com/go-gl/glh"
)

var (
	shapes        []drawable
	animateSwitch float32
	move          float32
)

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

// Initialize OpenGL
func Init() {
	blueTriangle := CreateTriangle()
	blueTriangle.SetTranslation(-7.0, 0.0, 0.0)
	blueTriangle.SetRotation(0.0)
	blueTriangle.SetScale(3.0)
	blueTriangle.SetColor(0, 0.2, 1.0)

	redRect := CreateSquare()
	redRect.SetTranslation(5.0, 0.0, 0.0)
	redRect.SetRotation(0.0)
	redRect.SetScale(3.0)
	redRect.SetColor(1.0, 0.2, 0.2)

	greenCircle := CreateCircle(0.5)
	greenCircle.SetTranslation(0.0, 0.0, 0.0)
	greenCircle.SetRotation(0.0)
	greenCircle.SetScale(3.0)
	greenCircle.SetColor(0.2, 1.0, 0.2)

	shapes = []drawable{
		blueTriangle,
		redRect,
		greenCircle}

	animateSwitch = 1.0
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
