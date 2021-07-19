package main

import (
	"fmt"
	"pratice-go/linked-list/list"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func TestSimple(t *testing.T) {
	linkedList := list.List{
		Head: &list.Node{
			Value: 0,
			Next:  nil,
		},
	}
	linkedList.PushFront(100)
	linkedList.PushFront(99)
	linkedList.PopBack()
	linkedList.PopFront()
	linkedList.PushBack(2)
	linkedList.PushBack(233)
	linkedList.Insert(1, 3)
	linkedList.Erase(1)
	linkedList.PrintAll()
	linkedList.Reverses()
	linkedList.PrintAll()
	linkedList.RemoveValue(2)
	linkedList.PrintAll()
	assertEqual(t, linkedList.Size(), 2, "The should be 2.")
	assertEqual(t, linkedList.Front(), 233, "The should be 233.")
	assertEqual(t, linkedList.Back(), 100, "The should be 100.")
}
