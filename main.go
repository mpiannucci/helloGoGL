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
	blueTriangle  triangle
	redTriangle   triangle
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
	blueTriangle.SetID("blueTriangle")
	blueTriangle.InitBuffers()
	blueTriangle.SetTranslation(-5.0, 0.0, 0.0)
	blueTriangle.SetRotation(0.0)
	blueTriangle.SetColor(0, 0.2, 1.0)

	redTriangle.SetID("redTriangle")
	redTriangle.InitBuffers()
	redTriangle.SetTranslation(5.0, 0.0, 0.0)
	redTriangle.SetRotation(45.0)
	redTriangle.SetColor(1.0, 0.2, 0.2)

	animateSwitch = 1.0
}

// Animate
func Animate() {
	// Increment the move counter
	move += 0.2 * animateSwitch

	// Get the new Y locations
	blueY := float32(move)
	redY := float32(-1.0 * move)

	if move > 10 || move < -10.0 {
		// Make sure nothing leaves the window so switch directions
		animateSwitch *= -1.0
	}

	// Set the translation for the shapes
	blueTriangle.SetTranslation(-5.0, blueY, 0.0)
	redTriangle.SetTranslation(5.0, redY, 0.0)
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

	window, err := glfw.CreateWindow(600, 400, "Hello Go GL", nil, nil)
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
		Animate()

		// Draw the drawablesss
		blueTriangle.Draw()
		redTriangle.Draw()

		// Swap Buffers
		window.SwapBuffers()
		glfw.PollEvents()
	}

}
