package util

import (
	"github.com/nicholasblaskey/webgl/webgl"

	mgl "github.com/go-gl/mathgl/mgl32"
)

func SetBool(gl *webgl.Gl, program *webgl.Program, uniformName string, x bool) {
	ul := gl.GetUniformLocation(program, uniformName)

	if x {
		gl.Uniform1i(ul, 1)
	} else {
		gl.Uniform1i(ul, 0)
	}
}

func SetInt(gl *webgl.Gl, program *webgl.Program, uniformName string, x int) {
	ul := gl.GetUniformLocation(program, uniformName)
	gl.Uniform1i(ul, x)
}

func SetFloat(gl *webgl.Gl, program *webgl.Program, uniformName string, x float32) {
	ul := gl.GetUniformLocation(program, uniformName)
	gl.Uniform1f(ul, x)
}

func SetVec2(gl *webgl.Gl, program *webgl.Program, uniformName string, x mgl.Vec2) {
	ul := gl.GetUniformLocation(program, uniformName)
	gl.Uniform2f(ul, x[0], x[1])
}

func SetVec3(gl *webgl.Gl, program *webgl.Program, uniformName string, x mgl.Vec3) {
	ul := gl.GetUniformLocation(program, uniformName)
	gl.Uniform3f(ul, x[0], x[1], x[2])
}

func SetVec4(gl *webgl.Gl, program *webgl.Program, uniformName string, x mgl.Vec4) {
	ul := gl.GetUniformLocation(program, uniformName)
	gl.Uniform4f(ul, x[0], x[1], x[2], x[3])
}

func SetMat2(gl *webgl.Gl, program *webgl.Program, uniformName string, x mgl.Mat2) {
	ul := gl.GetUniformLocation(program, uniformName)
	gl.UniformMatrix2fv(ul, false, x[:])
}

func SetMat3(gl *webgl.Gl, program *webgl.Program, uniformName string, x mgl.Mat3) {
	ul := gl.GetUniformLocation(program, uniformName)
	gl.UniformMatrix3fv(ul, false, x[:])
}

func SetMat4(gl *webgl.Gl, program *webgl.Program, uniformName string, x mgl.Mat4) {
	ul := gl.GetUniformLocation(program, uniformName)
	gl.UniformMatrix4fv(ul, false, x[:])
}
