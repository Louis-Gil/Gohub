package myarray

import (
	"unsafe"
)

type MyArray struct {
	data unsafe.Pointer
	length int
	elemSize uintptr
}

func NewMyArray(length int, elemSize uintptr) *MyArray {
	return &MyArray{
		data: unsafe.Pointer(new([100000]byte)),
		length: length,
		elemSize: elemSize,
	}
}

func (a *MyArray) Len() int {
	return a.length
}

func (a *MyArray) Set(index int, value unsafe.Pointer) {
	if index < 0 || index >= a.length {
		panic("index out of range")
	}
	elem := unsafe.Pointer(uintptr(a.data) + uintptr(index) * a.elemSize)
	*(*uintptr)(elem) = uintptr(value)
}

func (a *MyArray) Get(index int) unsafe.Pointer {
	if index < 0 || index >= a.length {
		panic("index out of range")
	}
	elem := unsafe.Pointer(uintptr(a.data) + uintptr(index) * a.elemSize)
	return *(*unsafe.Pointer)(elem)
}