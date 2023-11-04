package mystaticarray

import (
	"unsafe"
)

type MyStaticArray struct {
	data     unsafe.Pointer
	length   int
	elemSize uintptr
}

func NewMyStaticArray(length int, elemSize uintptr) *MyStaticArray {
	mem := make([]byte, length*int(elemSize))
	return &MyStaticArray{
		data:     unsafe.Pointer(&mem[0]),
		length:   length,
		elemSize: elemSize,
	}
}

func (a *MyStaticArray) Len() int {
	return a.length
}

func (a *MyStaticArray) Set(index int, value unsafe.Pointer) {
	if index < 0 || index >= a.length {
		panic("index out of range")
	}
	elem := unsafe.Pointer(uintptr(a.data) + uintptr(index)*a.elemSize)
	*(*uintptr)(elem) = uintptr(value)
}

func (a *MyStaticArray) Get(index int) unsafe.Pointer {
	if index < 0 || index >= a.length {
		panic("index out of range")
	}
	elem := unsafe.Pointer(uintptr(a.data) + uintptr(index)*a.elemSize)
	return *(*unsafe.Pointer)(elem)
}
