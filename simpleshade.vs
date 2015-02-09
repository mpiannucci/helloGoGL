#version 330 core

// Input vertex data, different for all executions of this shader.
layout(location = 0) in vec4 position;

void main() {
    gl_Position = position;
    gl_Position.w = 10.0;
}
