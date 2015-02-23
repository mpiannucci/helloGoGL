#version 330 core

// Ouput color
out vec3 color;

// Color uniform
uniform vec3 ColorVector;

void main() {
    // Output color
    color = ColorVector;
}