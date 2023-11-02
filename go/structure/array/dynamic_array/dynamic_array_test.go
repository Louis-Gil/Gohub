package mydynamicarray

import (
	"testing"
)

func TestMyDynamicArray (t *testing.T) {
	arr := NewMyDynamicArray[int]()

	if got := arr.Len(); got != 0 {
		t.Errorf("arr.Len() = %d; want 0", got)
	}

	val1 := 42
	arr.Set(0, val1)

	val2 := 35
	arr.Set(9, val2)

	if got := arr.Len(); got != 10 {
		t.Errorf("arr.Len() = %d; want 10", got)
	}

	if got := arr.Get(0); got != 42 {
		t.Errorf("arr.Get(0) = %d; want 42", got)
	}

	if got := arr.Get(9); got != 35 {
		t.Errorf("arr.Get(9) = %d; want 35", got)
	}
}

func TestOutOfBounds(t *testing.T) {
	arr := NewMyDynamicArray[int]()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	val := 42
	arr.Set(-1, val)
}