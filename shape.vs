#version 330 core

#define M_PI 3.1415926535897932384626433832795

// Input vertex data, different for all executions of this shader.
layout(location = 0) in vec4 position;

uniform vec2 Offset;
uniform float RotAngle;

void main() {
    vec4 totalOffset = vec4(Offset.x, Offset.y, 0.0, 0.0);
    float rads = RotAngle * (M_PI / 180.0);
    mat4 rotMat = mat4( cos(rads), -sin(rads), 0.0, 0.0,
                        sin(rads), cos(rads), 0.0, 0.0,
                        0.0, 0.0, 1.0, 0.0,
                        0.0, 0.0, 0.0, 1.0 );

    gl_Position = rotMat * position + totalOffset;
    gl_Position.w = 10.0;
}
