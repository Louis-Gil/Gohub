package array

import (
	"testing"
	"unsafe"
)

func TestMyArray (t *testing.T) {
	arr := NewMyArray(10, unsafe.Sizeof(int(0)))

	val1 := 42
	arr.Set(0, unsafe.Pointer(&val1))

	val2 := 35
	arr.Set(1, unsafe.Pointer(&val2))
	
	if got := arr.Len(); got != 10 {
		t.Errorf("arr.Len() = %d; want 10", got)
	}

	if got := *(*int)(arr.Get(0)); got != 42 {
		t.Errorf("arr.Get(0) = %d; want 42", got)
	}

	if got := *(*int)(arr.Get(1)); got != 35 {
		t.Errorf("arr.Get(1) = %d; want 35", got)
	}
}

func TestOutOfBounds(t *testing.T) {
	arr := NewMyArray(2, unsafe.Sizeof(int(0)))

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	val := 42
	arr.Set(2, unsafe.Pointer(&val))
}