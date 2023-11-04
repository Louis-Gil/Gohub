package mydynamicarray

type MyDynamicArray[T any] struct {
	data []T
}

func NewMyDynamicArray[T any]() *MyDynamicArray[T] {
	return &MyDynamicArray[T]{
		data: make([]T, 0),
	}
}

func (a *MyDynamicArray[T]) Len() int {
	return len(a.data)
}

func (a *MyDynamicArray[T]) Set(index int, value T) {
	if index < 0 {
		panic("index out of range")
	}
	if index >= len(a.data) {
		a.data = append(a.data, make([]T, index-len(a.data)+1)...)
	}
	a.data[index] = value
}

func (a *MyDynamicArray[T]) Get(index int) T {
	if index < 0 || index >= len(a.data) {
		panic("index out of range")
	}
	return a.data[index]
}
