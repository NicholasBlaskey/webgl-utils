package util

import (
	"github.com/nicholasblaskey/webgl/webgl"

	"errors"
)

func CreateProgram(gl *webgl.Gl, vertexSource, fragmentSource string) (*webgl.Program, error) {
	vShader, err := createShader(gl, webgl.VERTEX_SHADER, vertexSource)
	if err != nil {
		return nil, err
	}
	fShader, err := createShader(gl, webgl.FRAGMENT_SHADER, fragmentSource)
	if err != nil {
		return nil, err
	}

	program := gl.CreateProgram()
	if program == nil {
		return nil, errors.New("Could not create program")
	}

	gl.AttachShader(program, vShader)
	gl.AttachShader(program, fShader)
	gl.LinkProgram(program)

	linked := gl.GetProgramParameter(program, webgl.LINK_STATUS)
	if linked == 0 {
		info := gl.GetProgramInfoLog(program)

		gl.DeleteProgram(program)
		gl.DeleteShader(vShader)
		gl.DeleteShader(fShader)
		return nil, errors.New(info)
	}

	gl.UseProgram(program)

	return program, nil
}

func createShader(gl *webgl.Gl, shaderType int, source string) (*webgl.Shader, error) {
	shader := gl.CreateShader(shaderType)
	if shader == nil {
		return nil, errors.New("Could not create shader")
	}

	gl.ShaderSource(shader, source)
	gl.CompileShader(shader)

	compiled := gl.GetShaderParameter(shader, webgl.COMPILE_STATUS)
	if compiled == 0 {
		info := gl.GetShaderInfoLog(shader)

		gl.DeleteShader(shader)
		return nil, errors.New(info)
	}

	return shader, nil
}
