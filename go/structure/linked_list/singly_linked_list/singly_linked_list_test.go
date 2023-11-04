package mysinglylinkedlist

import (
	"reflect"
	"testing"
)

func TestMySinglyLinkedListAdd(t *testing.T) {
	arr := NewMySinglyLinkedList()

	arr.AddToTail(42)

	if got := arr.Len(); got != 1 {
		t.Errorf("arr.Len() = %d; want 1", got)
	}	

	if got := arr.Get(0); got != 42 {
		t.Errorf("arr.Get(0) = %d; want 42", got)
	}

	arr.AddToHead(35)

	if got := arr.GetAll(); !reflect.DeepEqual(got, []int{35, 42}) {
		t.Errorf("arr.GetAll() = %v; want [35, 42]", got)
	}

	arr.AddAtIndex(1, 17)

	if got := arr.GetAll(); !reflect.DeepEqual(got, []int{35, 17, 42}) {
		t.Errorf("arr.GetAll() = %v; want [35, 17, 42]", got)
	}
}

func TestMySinglyLinkedListRemove(t *testing.T) {
	arr := NewMySinglyLinkedList()

	arr.AddToTail(42)
	arr.AddToTail(35)
	arr.AddToTail(17)
	arr.AddToTail(28)

	arr.Remove(35)

	if got := arr.Len(); got != 3 {
		t.Errorf("arr.Len() = %d; want 3", got)
	}

	if got := arr.GetAll(); !reflect.DeepEqual(got, []int{42, 17, 28}) {
		t.Errorf("arr.GetAll() = %v; want [42, 17, 28]", got)
	}

	arr.RemoveFromHead()
	
	if got := arr.GetAll(); !reflect.DeepEqual(got, []int{17, 28}) {
		t.Errorf("arr.GetAll() = %v; want [17, 28]", got)
	}

	arr.RemoveFromTail()

	if got := arr.GetAll(); !reflect.DeepEqual(got, []int{17}) {
		t.Errorf("arr.GetAll() = %v; want [17]", got)
	}

	arr.RemoveAtIndex(0)

	if got := arr.GetAll(); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("arr.GetAll() = %v; want []", got)
	}
}

func TestMySinglyLinkedListGetOutOfBounds(t *testing.T) {
	arr := NewMySinglyLinkedList()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	arr.AddToTail(42)
	arr.AddToTail(35)

	arr.Get(2)
}

func TestMySinglyLinkedListAddAtIndexOutOfBounds(t *testing.T) {
	list := NewMySinglyLinkedList()

	defer func() {
			if r := recover(); r == nil {
					t.Errorf("AddAtIndex did not panic on out of bounds")
			}
	}()

	list.AddAtIndex(1, 42)
}

func TestMySinglyLinkedListRemoveAtIndexOutOfBounds(t *testing.T) {
	list := NewMySinglyLinkedList()
	list.AddToTail(42)

	defer func() {
			if r := recover(); r == nil {
					t.Errorf("RemoveAtIndex did not panic on out of bounds")
			}
	}()

	list.RemoveAtIndex(1)
}