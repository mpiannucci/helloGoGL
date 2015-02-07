// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"math"
	"os"
	"runtime"

	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
)

var (
	wave    uint
	trans_x = 0.0
)

// Draw a 2-D WaveForm
func Wave(amplitude, duration, resolution float64) {
	gl.Begin(gl.LINE_STRIP)
	for i := float64(0.0); i <= duration; i += resolution {
		gl.Vertex3d(i, amplitude*math.Sin(i), 0.0)
	}
	gl.End()
}

// OpenGL drawing and timing
func Draw() {
	// Draw the clear color
	gl.Clear(gl.COLOR_BUFFER_BIT)

	// Draw a waveee
	gl.PushMatrix()
	gl.Translated(trans_x, 0.0, 0.0)
	gl.Rotated(0.0, 0.0, 0.0, 1.0)
	gl.CallList(wave)
	gl.PopMatrix()
}

// Animate!!
func Animate() {
	if trans_x < 19.8 {
		trans_x += 0.1
	} else {
		trans_x = 0.0
	}
}

// Initialize OpenGL
func Init() {

	// Make the waves
	wave = gl.GenLists(1)
	gl.NewList(wave, gl.COMPILE)
	Wave(4.0, 20.0, 0.1)
	gl.EndList()
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

	window, err := glfw.CreateWindow(1024, 768, "Running Waves", nil, nil)
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

	// Equivalent to a do... while
	for ok := true; ok; ok = (window.GetKey(glfw.KeyEscape) != glfw.Press && !window.ShouldClose()) {
		// Draw a wave
		Draw()

		// Animate it
		Animate()

		// Swap Buffers
		window.SwapBuffers()
		glfw.PollEvents()
	}

}
