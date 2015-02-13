#version 330 core

// Input vertex data, different for all executions of this shader.
layout(location = 0) in vec4 position;

uniform mat4 MVP;

void main() {
    // Set the projection matrix
    gl_Position = MVP * position;
}
