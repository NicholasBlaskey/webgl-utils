package util

import (
	"github.com/nicholasblaskey/webgl/webgl"

	"reflect"
	"unsafe"
)

type buffer struct {
	WebGlBuffer *webgl.Buffer
	BufferType  int
	DataType    int
	Size        int
	Usage       int
	VertexCount int
}

func NewBufferFloat(gl *webgl.Gl) *buffer {
	return NewBuffer(gl, webgl.ARRAY_BUFFER, webgl.FLOAT, 1, webgl.STATIC_DRAW)
}

func NewBufferVec2(gl *webgl.Gl) *buffer {
	return NewBuffer(gl, webgl.ARRAY_BUFFER, webgl.FLOAT, 2, webgl.STATIC_DRAW)
}

func NewBufferVec3(gl *webgl.Gl) *buffer {
	return NewBuffer(gl, webgl.ARRAY_BUFFER, webgl.FLOAT, 3, webgl.STATIC_DRAW)
}

func NewBufferVec4(gl *webgl.Gl) *buffer {
	return NewBuffer(gl, webgl.ARRAY_BUFFER, webgl.FLOAT, 4, webgl.STATIC_DRAW)
}

func NewBuffer(gl *webgl.Gl, bufferType, dataType, size, usage int) *buffer {
	return &buffer{gl.CreateBuffer(), bufferType, dataType, size, usage, 0}
}

func (b *buffer) BindData(gl *webgl.Gl, data interface{}) {
	b.VertexCount = getLen(data) / b.Size

	gl.BindBuffer(b.BufferType, b.WebGlBuffer)
	gl.BufferData(b.BufferType, data, b.Usage)
}

func (b *buffer) BindToAttrib(gl *webgl.Gl, program *webgl.Program, attrib string) {
	attribLoc := gl.GetAttribLocation(program, attrib)

	gl.BindBuffer(b.BufferType, b.WebGlBuffer)
	gl.VertexAttribPointer(attribLoc, b.Size, b.DataType, false, 0, 0)
	gl.EnableVertexAttribArray(attribLoc)
}

func getLen(s interface{}) int {
	switch s := s.(type) {
	case []int8:
		return (*reflect.SliceHeader)(unsafe.Pointer(&s)).Len
	case []int16:
		return (*reflect.SliceHeader)(unsafe.Pointer(&s)).Len
	case []int32:
		return (*reflect.SliceHeader)(unsafe.Pointer(&s)).Len
	case []int64:
		return (*reflect.SliceHeader)(unsafe.Pointer(&s)).Len
	case []uint8:
		return (*reflect.SliceHeader)(unsafe.Pointer(&s)).Len
	case []uint16:
		return (*reflect.SliceHeader)(unsafe.Pointer(&s)).Len
	case []uint32:
		return (*reflect.SliceHeader)(unsafe.Pointer(&s)).Len
	case []uint64:
		return (*reflect.SliceHeader)(unsafe.Pointer(&s)).Len
	case []float32:
		return (*reflect.SliceHeader)(unsafe.Pointer(&s)).Len
	case []float64:
		return (*reflect.SliceHeader)(unsafe.Pointer(&s)).Len
	}
	panic("Could not find number of vertices of passed in data")
}
